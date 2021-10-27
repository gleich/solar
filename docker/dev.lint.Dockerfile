# hadolint ignore=DL3007
FROM golangci/golangci-lint:latest

# Meta data
LABEL maintainer="email@mattglei.ch"
LABEL description="🛰️ Clone all your starred GitHub repos"

# Copying over files
COPY . /usr/src/app
WORKDIR /usr/src/app

# Installing hadolint:
WORKDIR /usr/bin
RUN curl -sL -o hadolint "https://github.com/hadolint/hadolint/releases/download/v1.17.6/hadolint-$(uname -s)-$(uname -m)" \
    && chmod 700 hadolint

# Installing goreleaser
WORKDIR /
RUN git clone https://github.com/goreleaser/goreleaser
WORKDIR /goreleaser
RUN go get ./... \
    && go build -o goreleaser . \
    && mv goreleaser /usr/bin

# Installing make
RUN apt-get update && apt-get install make=4.3-4.1 -y --no-install-recommends \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# Installing go 1.17
RUN go install "golang.org/dl/go1.17@latest" \
    && go1.17 download \
    && mv "$(which go1.17)" "$(which go)"

WORKDIR /usr/src/app

CMD ["make", "local-lint"]
