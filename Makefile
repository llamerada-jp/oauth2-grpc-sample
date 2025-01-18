SHELL := /bin/bash -o pipefail
PROTO_GO_FILES := proto/authenticate.pb.go proto/authenticate_grpc.pb.go proto/commands.pb.go proto/commands_grpc.pb.go
GO_FILES := $(shell find . -name '*.go' -not -path './node_modules/*')

build: bin/client bin/server static/script.js

bin/client: $(PROTO_GO_FILES) $(GO_FILES)
	mkdir -p bin
	go build -o bin/client ./client/

bin/server: $(PROTO_GO_FILES) $(GO_FILES) server/template.html
	mkdir -p bin
	go build -o bin/server ./server/

proto/authenticate.pb.go: proto/authenticate.proto
	protoc --go_out=module=github.com/llamerada-jp/oauth2-grpc-sample:. $<

proto/authenticate_grpc.pb.go: proto/authenticate.proto
	protoc --go-grpc_out=module=github.com/llamerada-jp/oauth2-grpc-sample:. $<

proto/commands.pb.go: proto/commands.proto
	protoc --go_out=module=github.com/llamerada-jp/oauth2-grpc-sample:. $<

proto/commands_grpc.pb.go: proto/commands.proto
	protoc --go-grpc_out=module=github.com/llamerada-jp/oauth2-grpc-sample:. $<

static/script.js: proto/commands.proto ts/main.ts
	npm run build

ts/commands_pb.js: proto/commands.proto
	protoc \
	  --plugin=protoc-gen-grpc-web=./node_modules/.bin/protoc-gen-grpc-web \
	  --js_out=import_style=commonjs:ts \
		--grpc-web_out=import_style=typescript,mode=grpcwebtext:ts \
		--proto_path=proto commands.proto

.PHONY: setup
setup:
	sudo apt install -y protobuf-compiler
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	npm install

.PHONY: generate-cert
generate-cert:
	openssl req -x509 -out localhost.crt -keyout localhost.key \
    -newkey rsa:2048 -nodes -sha256 \
    -subj '/CN=localhost' -extensions EXT -config <( \
     printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:localhost\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")
