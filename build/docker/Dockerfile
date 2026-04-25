# Build arguments provided by Docker Buildx and GoReleaser dockers_v2
ARG BUILDPLATFORM
ARG TARGETPLATFORM

# Pinned base image used only to extract CA certificates and timezone data
ARG BASE_IMAGE=alpine:3.23.4@sha256:5b10f432ef3da1b8d4c7eb6c487f2f5a8f096bc91145e68878dd4a5019afde11

# Builder stage
FROM --platform=$BUILDPLATFORM ${BASE_IMAGE} AS builder
RUN apk add --no-cache ca-certificates tzdata

# Final minimal scratch image
FROM scratch

# Re-declare TARGETPLATFORM for this stage
ARG TARGETPLATFORM

# Copy supporting files from builder stage
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo

# Copy the Go binary (placed by GoReleaser under $TARGETPLATFORM/)
COPY --chmod=0755 ${TARGETPLATFORM}/todo /todo

# Run as non-root using the standard distroless UID/GID (best practice)
USER 65532

# Copy the binary built by GoReleaser (dockers_v2 puts
EXPOSE 3000

# Set the entrypoint
ENTRYPOINT ["/todo"]
