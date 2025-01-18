package main

import (
	"context"
	"sync"
	"time"
)

const sessionKeyInContext = "session"

type sessionRecord struct {
	mtx       sync.Mutex
	store     map[string]string
	timestamp time.Time
}

type sessionStore struct {
	mtx      sync.Mutex
	records  map[string]*sessionRecord
	lifetime time.Duration
}

func newSessionStore(ctx context.Context, lifetime time.Duration) *sessionStore {
	ss := &sessionStore{
		records:  make(map[string]*sessionRecord),
		lifetime: lifetime,
	}

	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				ss.cleanup()
			}
		}
	}()

	return ss
}

func getSession(ctx context.Context) *sessionRecord {
	s, _ := ctx.Value(sessionKeyInContext).(*sessionRecord)
	return s
}

func (ss *sessionStore) cleanup() {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	for k, s := range ss.records {
		if time.Since(s.timestamp) > ss.lifetime {
			delete(ss.records, k)
		}
	}
}

func (ss *sessionStore) create() (*sessionRecord, string) {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	id := randomString(16)
	s := &sessionRecord{
		store:     make(map[string]string),
		timestamp: time.Now(),
	}
	ss.records[id] = s

	return s, id
}

func (ss *sessionStore) delete(id string) {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	delete(ss.records, id)
}

func (ss *sessionStore) newContextWithSession(ctx context.Context, id string) (context.Context, *sessionRecord) {
	ss.mtx.Lock()
	defer ss.mtx.Unlock()

	s, ok := ss.records[id]
	if ok {
		s.timestamp = time.Now()
	}

	return context.WithValue(ctx, sessionKeyInContext, s), s
}

func (s *sessionRecord) get(key string) (string, bool) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	value, ok := s.store[key]
	return value, ok
}

func (s *sessionRecord) set(key, value string) {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	s.store[key] = value
}
