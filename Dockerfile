# Build the manager binary
FROM docker.m.daocloud.io/golang:1.18.1 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY commands/ commands/
COPY controllers/ controllers/
COPY errors/ errors/
COPY operations/ operations/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o manager main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
# distroless cannot run `kubeadm upgrade apply` smoothly
# FROM gcr.m.daocloud.io/distroless/static:nonroot
FROM docker.m.daocloud.io/ubuntu
RUN apt-get update -q -y && apt-get install -q -y curl && apt clean all

WORKDIR /
COPY --from=builder /workspace/manager .
# USER nonroot:nonroot

ENTRYPOINT ["/manager"]
