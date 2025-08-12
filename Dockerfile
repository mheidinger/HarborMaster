FROM golang:1.21 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0
RUN go build -o HarborMaster ./cmd/HarborMaster

FROM alpine:3.18

WORKDIR /app
COPY --from=builder /build/HarborMaster .
COPY static ./static
COPY templates ./templates

ENV GIN_MODE release

EXPOSE 4181
ENTRYPOINT ["./HarborMaster"]
