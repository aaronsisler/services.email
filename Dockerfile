FROM public.ecr.aws/lambda/provided:al2 AS builder

RUN yum install -y tar gzip curl

# Install specific Go version manually
RUN curl -LO https://go.dev/dl/go1.24.4.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz && \
    rm go1.24.4.linux-amd64.tar.gz && \
    ln -s /usr/local/go/bin/go /usr/bin/go && \
    ln -s /usr/local/go/bin/gofmt /usr/bin/gofmt

# Setup workdir
WORKDIR /app

# Get the dependencies for the project
COPY go.mod go.sum ./
RUN go mod download

# Copy Go source
COPY cmd ./cmd
COPY handlers ./handlers

# Optional: debug structure
RUN ls -R cmd && ls -R handlers

# Build binary for Amazon Linux 2
RUN GOOS=linux GOARCH=amd64 go build -o bootstrap ./cmd/email/post && \
    chmod +x bootstrap

# --- Stage 2: Output the built file to a shared volume
FROM alpine:3.19 AS output
WORKDIR /out
COPY --from=builder /app/bootstrap ./bootstrap