# syntax=docker/dockerfile:experimental

FROM golang:1.16

ENV GOPRIVATE git.myrepo.com/*
ENV GO111MODULE on

RUN apt-get install openssh-client git

RUN mkdir -p -m 0600 ~/.ssh \
    && ssh-keyscan git.myrepo.com >> ~/.ssh/known_hosts \
    && git config --global url."ssh://git@git.myrepo.com/".insteadOf "https://git.myrepo.com/"

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar -xvz \
    && mv ./migrate.linux-amd64 /bin/migrate

RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.38.0 \
    mv ./bin/golangci-lint /bin/golangci-lint

WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=ssh go mod download

