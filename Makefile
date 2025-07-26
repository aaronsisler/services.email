.PHONY: test

# Set Go environment
GOOS=linux
GOARCH=amd64

# Building the handlers
build-handler-health-get:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/handler_health_get/bootstrap ./cmd/health/get

build-handler-email-post:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/handler_email_post/bootstrap ./cmd/email/post

# Ziping the handlers
zip-handler-health-get:
	zip handler_health_get.zip -j bin/handler_health_get/bootstrap

zip-handler-email-post:
	zip handler_email_post.zip -j bin/handler_email_post/bootstrap

# Build all handlers
build:
	make build-handler-health-get
	make build-handler-email-post

test:
	go install github.com/cucumber/godog/cmd/godog@latest
	cd test/steps && go test

# Zip all handlers
zip:
	make zip-handler-health-get
	make zip-handler-email-post