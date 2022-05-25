FROM golang:1.18 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static" -o HarborMaster ./cmd/HarborMaster

FROM alpine:3.11

WORKDIR /app
COPY --from=builder /build/HarborMaster .
COPY static ./static
COPY templates ./templates

ENV GIN_MODE release

EXPOSE 4181
ENTRYPOINT ["./HarborMaster"]