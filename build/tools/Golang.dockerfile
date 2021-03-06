ARG GO_VERSION
FROM "golang:${GO_VERSION}"

WORKDIR /app

COPY ./go.sum ./go.mod ./

# Get dependencies (cached)
RUN go mod download

# Copy the source code
COPY . .
