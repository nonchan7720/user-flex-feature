FROM --platform=$BUILDPLATFORM node:18.19.0-bullseye-slim as node

FROM golang:1.22.1-bullseye

COPY --from=node /usr/local/bin/ /usr/local/bin/
COPY --from=node /usr/local/lib/node_modules /usr/local/lib/node_modules

ARG GIT_VERSION=2.39.1

RUN apt-get update \
  && apt-get install --no-install-recommends -y \
  make \
  sudo \
  protobuf-compiler \
  vim \
  procps \
  curl \
  openssh-client \
  gettext \
  libexpat1-dev \
  libz-dev \
  libssl-dev \
  libcurl4-openssl-dev \
  lsb-release \
  && cd /usr/local/src \
  && wget https://github.com/git/git/archive/refs/tags/v${GIT_VERSION}.tar.gz \
  && tar -xzf v${GIT_VERSION}.tar.gz \
  && cd git-${GIT_VERSION} \
  && make prefix=/usr/local all -j "$(nproc)" \
  && make prefix=/usr/local install \
  && go install github.com/cweill/gotests/gotests@v1.6.0 \
  && go install github.com/fatih/gomodifytags@v1.16.0 \
  && go install github.com/josharian/impl@v1.2.0 \
  && go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0 \
  && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2 \
  && go install golang.org/x/tools/gopls@v0.11.0 \
  && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.30.0 \
  && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0 \
  && go install github.com/stormcat24/protodep@v0.1.7 \
  && go install github.com/go-delve/delve/cmd/dlv@v1.22.1

# aws cliのインストール
ARG TARGETPLATFORM
RUN apt-get update \
  && apt-get install -y unzip
RUN if [ "$TARGETPLATFORM" = "linux/arm64" ]; then \
  curl "https://awscli.amazonaws.com/awscli-exe-linux-aarch64.zip" -o "/tmp/awscliv2.zip"; \
  elif [ "$TARGETPLATFORM" = "linux/amd64" ]; then \
  curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "/tmp/awscliv2.zip"; \
  fi; \
  cd /tmp/; \
  unzip awscliv2.zip; \
  ./aws/install; \
  rm -rf awscliv2.zip ./aws; \
  aws --version;

ARG USER_NAME=user \
  GROUP_NAME=user \
  USER_ID=1000 \
  GROUP_ID=1000

RUN groupadd -r -g $GROUP_ID ${GROUP_NAME} \
  && useradd -m -s /bin/bash -u $USER_ID -g $GROUP_ID -G sudo $USER_NAME \
  && echo $USER_NAME:$USER_NAME | chpasswd \
  && echo "$USER_NAME   ALL=(ALL) NOPASSWD:ALL" >> /etc/sudoers \
  && mkdir -p /src \
  && chown $USER_ID:$GROUP_ID -R /go /src \
  && rm -rf /var/lib/apt/lists/*

USER $USER_NAME

ENV TZ=Asia/Tokyo

WORKDIR /src
