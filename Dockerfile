FROM localstack/localstack:latest

# Install tar, gzip, curl (Alpine equivalents)
# Install required packages using apt
RUN apt-get update && apt-get install -y \
    tar \
    gzip \
    curl \
    unzip \
    && apt-get clean

# Install specific Go version manually
RUN curl -LO https://go.dev/dl/go1.24.4.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go1.24.4.linux-amd64.tar.gz && \
    rm go1.24.4.linux-amd64.tar.gz && \
    ln -s /usr/local/go/bin/go /usr/bin/go && \
    ln -s /usr/local/go/bin/gofmt /usr/bin/gofmt

