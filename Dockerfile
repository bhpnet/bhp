# Simple usage with a mounted data directory:
# > docker build -t bhp .
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.bhpd:/root/.bhpd -v ~/.bhpcli:/root/.bhpcli bhp bhpd init
# > docker run -it -p 46657:46657 -p 46656:46656 -v ~/.bhpd:/root/.bhpd -v ~/.bhpcli:/root/.bhpcli bhp bhpd start
FROM golang:alpine AS build-env

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3

# Set working directory for the build
WORKDIR /go/src/github.com/bhpnet/bhp-dev

# Add source files
COPY . .

# Build BHP
RUN GOPROXY=http://goproxy.cn make install

# Install minimum necessary dependencies, build Cosmos SDK, remove packages
RUN apk add --no-cache $PACKAGES && \
    make tools && \
    make install

# Final image
FROM alpine:edge

# Install ca-certificates
RUN apk add --update ca-certificates
WORKDIR /root

# Copy over binaries from the build-env
COPY --from=build-env /go/bin/bhpd /usr/bin/bhpd
COPY --from=build-env /go/bin/bhpcli /usr/bin/bhpcli

# Run bhpd by default, omit entrypoint to ease using container with bhpcli
CMD ["bhpd"]
