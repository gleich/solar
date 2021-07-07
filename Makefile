##########
# Building
##########

build-docker-prod:
	docker build -f docker/prod.Dockerfile -t mattgleich/solar:latest .
build-docker-dev:
	docker build -f docker/dev.Dockerfile -t mattgleich/solar:test .
build-docker-dev-lint:
	docker build -f docker/dev.lint.Dockerfile -t mattgleich/solar:lint .
build-go:
	go get -v -t -d ./...
	go build -v .
	rm solar

#########
# Linting
#########

lint-golangci:
	golangci-lint run
lint-gomod:
	go mod tidy
	git diff --exit-code go.mod
	git diff --exit-code go.sum
lint-goreleaser:
	goreleaser check
lint-hadolint:
	hadolint docker/prod.Dockerfile
	hadolint docker/dev.Dockerfile
	hadolint docker/dev.lint.Dockerfile
lint-in-docker: build-docker-dev-lint
	docker run mattgleich/solar:lint

#########
# Testing
#########

test-go:
	go get -v -t -d ./...
	go test -v ./...
test-in-docker: build-docker-dev
	docker run mattgleich/solar:test

##########
# Grouping
##########

# Testing
local-test: test-go
docker-test: test-in-docker
# Linting
local-lint: lint-golangci lint-goreleaser lint-hadolint lint-gomod
docker-lint: lint-in-docker
# Build
local-build: build-docker-prod build-docker-dev build-docker-dev-lint
