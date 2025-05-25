# build
FROM golang:1.24-bookworm AS be_builder

COPY . /build

WORKDIR /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o gopu main.go

# alpine
FROM alpine:3.21

RUN apk add ca-certificates

COPY --from=be_builder /build/gopu /gopu

WORKDIR /

CMD ["./gopu"]