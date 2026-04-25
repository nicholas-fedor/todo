ARG BASE_IMAGE=alpine:3.23.4@sha256:5b10f432ef3da1b8d4c7eb6c487f2f5a8f096bc91145e68878dd4a5019afde11
ARG GOOS=linux
ARG GOARCH

FROM --platform=$BUILDPLATFORM $BASE_IMAGE AS builder

RUN apk add --no-cache \
    ca-certificates \
    tzdata

FROM scratch

# Copy ca-certs and timezone from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy the binary built by GoReleaser (dockers_v2 puts it under $TARGETPLATFORM/)
ARG TARGETPLATFORM
COPY --chmod=0755 $TARGETPLATFORM/todo /todo

# Expose the application port
EXPOSE 3000

# Set the entrypoint
ENTRYPOINT ["/todo"]
