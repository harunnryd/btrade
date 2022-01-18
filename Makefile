SHELL                 = /bin/bash

APP_NAME              = btrade
VERSION               = $(shell git describe --always --tags)
GIT_COMMIT            = $(shell git rev-parse HEAD)
GIT_DIRTY             = $(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE            = $(shell date '+%Y-%m-%d-%H:%M:%S')

.PHONY: default
default: help

.PHONY: help
help:
	@echo 'Management commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make build                              Compile the project.'
	@echo '    make package                            Build, tag, and push Docker image.'
	@echo '    make run ARGS=                          Run with supplied arguments.'
	@echo '    make test                               Run tests on a compiled project.'
	@echo '    make clean                              Clean the directory tree.'
	@echo

.PHONY: build
build:
	@echo "Building ${APP_NAME} ${VERSION}"
	go build -ldflags "-w -X github.com/harunnryd/${APP_NAME}/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/harunnryd/${APP_NAME}/version.Version=${VERSION} -X github.com/harunnryd/${APP_NAME}/version.Environment=${ENV} -X github.com/harunnryd/${APP_NAME}/version.BuildDate=${BUILD_DATE}" -o bin/${APP_NAME}

.PHONY: package
package:
	@echo "Build, tag, and push Docker image ${APP_NAME} ${VERSION} ${GIT_COMMIT}"
	docker buildx build \
		--build-arg VERSION=${VERSION},GIT_COMMIT=${GIT_COMMIT}${GIT_DIRTY} \
		--cache-from type=local,src=/tmp/.buildx-cache \
		--cache-to type=local,dest=/tmp/.buildx-cache \
		--tag ${DOCKER_REPOSITORY}/${APP_NAME}:${GIT_COMMIT} \
		--tag ${DOCKER_REPOSITORY}/${APP_NAME}:${VERSION} \
		--tag ${DOCKER_REPOSITORY}/${APP_NAME}:latest \
		--push .

.PHONY: prune
prune:
	@echo "Removing Docker image ${DOCKER_REPOSITORY}/${APP_NAME}:${IMAGE_TAG}"
	gcloud container images delete \
		--force-delete-tags \
		--quiet \
		${DOCKER_REPOSITORY}/${APP_NAME}:${IMAGE_TAG}

.PHONY: run
run: build
	@echo "Running ${APP_NAME} ${VERSION}"
	bin/${APP_NAME} ${ARGS}

.PHONY: test
test:
	@echo "Testing ${APP_NAME} ${VERSION}"
	go test -race ./...

.PHONY: clean
clean:
	@echo "Removing ${APP_NAME} ${VERSION}"
	@test ! -e bin/${APP_NAME} || rm bin/${APP_NAME}

.PHONY: get-app-name
get-app-name:
	@echo ${APP_NAME}
