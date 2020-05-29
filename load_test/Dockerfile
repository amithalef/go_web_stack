############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
RUN apk update && apk add curl && apk add --no-cache git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR $GOPATH/src/github.com/amithnair91/go_web_stack/load_test
# Copy Source code
COPY . .
# Fetch dependencies.
RUN dep ensure

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /go/app app/main.go

RUN chmod +x /go/app
############################
# STEP 2 build a small image
############################
FROM scratch

ENV MONGO_HOST=""
ENV MONGO_PORT=""
ENV MONGO_DATABASE=""
ENV APP_PORT="8080"

EXPOSE 8080

# Copy our static executable.
COPY --from=builder /go/app /app

# Run the app binary.
CMD ["/app"]
