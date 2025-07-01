# Set Go environment
GOOS=linux
GOARCH=amd64

# Building the handlers
build-handler-health-get:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/handler_health_get/bootstrap ./cmd/health/get

# Ziping the handlers
zip-handler-health-get:
	zip handler_health_get.zip -j bin/handler_health_get/bootstrap

# Build all handlers
build:
	make build-handler-health-get

# Zip all handlers
zip:
	make zip-handler-health-get