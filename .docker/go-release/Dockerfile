ARG USER_NAME='app'
ARG USER_UID='1000'
ARG USER_GID='1000'
ARG USER_GROUP='app_users'
ARG REPO='github.com/kohirens/go-release'
ARG APP_NAME='go-release'
ARG ALPINE_VER='3.19'
ARG GO_VER='1.22'

FROM golang:${GO_VER}-alpine${ALPINE_VER} AS build

ARG USER_NAME
ARG USER_UID
ARG USER_GID
ARG USER_GROUP
ARG REPO

ENV GOPATH /home/${USER_NAME}
ENV CGO_ENABLED=0
ENV WORK_DIR=/home/${USER_NAME}/src/${REPO}
ENV SHELL=/bin/sh

# Update OS and other software
RUN apk --no-progress --purge --no-cache upgrade \
 && apk --no-progress --purge --no-cache add --upgrade \
    git \
    zip \
 && rm -vrf /var/cache/apk/* \
 && rm -rf /tmp/*

# Add a non-root group and user.
RUN addgroup --system --gid ${USER_GID} ${USER_GROUP} \
 && adduser --system \
    --disabled-password \
    --ingroup ${USER_GROUP} \
    --uid ${USER_UID} \
    ${USER_NAME}

USER ${USER_NAME}

# Make directories with current user permissions
RUN mkdir -p ~/bin ~/src/${REPO}

ENV PATH="${PATH}:/home/${USER_NAME}/bin"

WORKDIR "/home/${USER_NAME}/src/${REPO}"

# Add source files
COPY --chown=${USER_NAME}:${USER_GROUP} . .
COPY --from=kohirens/git-tool-belt:2.1.2 /usr/local/bin/git-tool-belt /usr/local/bin/

# Build
RUN git config --global --add safe.directory "/home/${USER_NAME}/src/${REPO}" \
 && go mod tidy \
 && go generate \
 && go build \
 && go install

FROM build AS dev

COPY .docker/go-release/start.sh /usr/local/bin/start.sh

ENTRYPOINT [ "start.sh" ]

HEALTHCHECK --interval=5s --timeout=3s --start-period=3s --retries=4 \
    CMD go-get-latest -help || exit 1

CMD [ ]

FROM build AS release

ARG USER_NAME

ENV PATH="${PATH}:/home/${USER_NAME}/bin"

ENTRYPOINT [ "go-release" ]
