package main

import (
	"context"
	"crypto/rand"
	_ "embed"
	"encoding/base64"
	"html/template"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

//go:embed template.html
var htmlTemplateRaw string
var htmlTemplate *template.Template

const (
	cookieKeyPair = "thisIsSecretKeyPair"
	sessionKey    = "sessionKeyString"
	// githubIssuer  = "https://token.actions.githubusercontent.com"
)

type templateData struct {
	Authenticated bool
	User          string
	GRPCSessionID string
}

func newHTTPHandler(ctx context.Context, sessionStore *sessionStore, clientID, clientSecret string) http.Handler {
	// Use session to store save state value.
	store := sessions.NewCookieStore([]byte(cookieKeyPair))
	// store.Options.HttpOnly = true

	/*
		provider, err := oidc.NewProvider(ctx, githubIssuer)
		if err != nil {
			log.Fatalf("failed to get provider: %v", err)
		}
	*/

	oauth2Github := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		// Endpoint:     provider.Endpoint(),
		Endpoint:    endpoints.GitHub,
		RedirectURL: "https://localhost:8443/callback_github",
		Scopes:      []string{"user"},
	}
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, sessionKey)
		if err != nil {
			session = sessions.NewSession(store, sessionKey)
		}

		session.Save(r, w)
		_, ok := session.Values["authenticate_by"]
		if session.IsNew || !ok {
			htmlTemplate.Execute(w, templateData{
				Authenticated: false,
			})
			return
		}

		htmlTemplate.Execute(w, templateData{
			Authenticated: true,
			User:          session.Values["github_user"].(string),
			GRPCSessionID: session.Values["grpc_session_id"].(string),
		})
	})

	// Serve static files
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/connect_github", func(w http.ResponseWriter, r *http.Request) {
		session := sessions.NewSession(store, sessionKey)

		// create new state string
		binaryData := make([]byte, 32)
		if _, err := rand.Read(binaryData); err != nil {
			log.Fatalf("failed to generate random string: %v", err)
		}
		state := base64.StdEncoding.EncodeToString(binaryData)
		session.Values["STATE"] = state
		session.Options.MaxAge = store.Options.MaxAge
		session.Save(r, w)

		url := oauth2Github.AuthCodeURL(state)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	})

	mux.HandleFunc("/callback_github", func(w http.ResponseWriter, r *http.Request) {
		httpSession, err := store.Get(r, sessionKey)
		if err != nil {
			http.Error(w, "Internal Server Error(failed to get session)", http.StatusInternalServerError)
			return
		}

		if httpSession.IsNew {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		// check state
		expectedState, ok := httpSession.Values["STATE"]
		if !ok {
			http.Error(w, "Internal Server Error(failed to get state)", http.StatusInternalServerError)
			return
		}
		if r.URL.Query().Get("state") != expectedState {
			http.Error(w, "Internal Server Error(state is invalid)", http.StatusInternalServerError)
			return
		}

		// exchange code
		code := r.URL.Query().Get("code")
		token, err := oauth2Github.Exchange(ctx, code)
		if err != nil {
			http.Error(w, "Internal Server Error(unable to exchange code for token)\n"+err.Error(),
				http.StatusInternalServerError)
			return
		}

		// verify id_token
		/*
			idToken := token.Extra("id_token").(string)
			fmt.Println("token:", *token)
			if _, err := provider.Verifier(&oidc.Config{ClientID: oauth2Github.ClientID}).Verify(ctx, idToken); err != nil {
				http.Error(w, "Internal Server Error(unable to verify id_token)\n"+err.Error(),
					http.StatusInternalServerError)
				return
			}
		*/

		// get github user
		client := github.NewClient(oauth2Github.Client(ctx, token))
		user, _, err := client.Users.Get(ctx, "")
		if err != nil {
			http.Error(w, "Internal Server Error(unable to get github user)\n"+err.Error(),
				http.StatusInternalServerError)
			return
		}

		// create session for grpc
		grpcSession, sessionID := sessionStore.create()
		grpcSession.set(sessionKeyUser, *user.Login)

		// recreate session
		httpSession = sessions.NewSession(store, sessionKey)
		httpSession.Options.MaxAge = store.Options.MaxAge
		httpSession.Values["authenticate_by"] = "github"
		httpSession.Values["github_user"] = *user.Login
		httpSession.Values["grpc_session_id"] = sessionID
		httpSession.Save(r, w)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})

	mux.HandleFunc("/disconnect", func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, sessionKey)
		if err != nil {
			http.Error(w, "Internal Server Error(failed to get session)", http.StatusInternalServerError)
			return
		}

		// destroy grpc session
		grpcSessionID, ok := session.Values["grpc_session_id"].(string)
		if ok {
			sessionStore.delete(grpcSessionID)
		}
		// destroy session
		session.Options.MaxAge = -1
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	})

	return mux
}

func init() {
	tmpl, err := template.New("").Parse(htmlTemplateRaw)
	if err != nil {
		panic(err)
	}
	htmlTemplate = tmpl
}
