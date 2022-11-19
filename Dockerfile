# Simple usage with a mounted data directory:
# > docker build -t simapp .
#
# Server:
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.simapp:/root/.fury fury furyd init redshift
# TODO: need to set validator in genesis so start runs
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.fury:/root/.fury fury furyd start
#
# Client: (Note the simapp binary always looks at ~/.fury we can bind to different local storage)
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.furycli:/root/.fury fury furyd keys add user1
# > docker run -it -p 26657:26657 -p 26656:26656 -v ~/.furycli:/root/.fury fury furyd keys list
# TODO: demo connecting rest-server (or is this in server now?)
FROM golang:1.19-alpine AS build-env

# Install minimum necessary dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3
RUN apk add --no-cache $PACKAGES

# Set working directory for the build
WORKDIR /go/src/github.com/redactedfury/fury

# Add source files
COPY . .

# install simapp, remove packages
RUN make build


# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/src/github.com/redactedfury/fury/build/furyd /usr/bin/furyd

EXPOSE 26656 26657 1317 9090

# Run simd by default, omit entrypoint to ease using container with furycli
CMD ["furyd"]