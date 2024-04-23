FROM golang:1.21-alpine as builder

WORKDIR /usr/src

COPY go.mod go.sum ./

RUN go mod download
RUN go mod verify

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /usr/src/app ./cmd

FROM alpine

COPY --from=builder /usr/src/app /usr/src/app
COPY --from=builder /usr/src/poc-innovation-iot-43fd0cc8a313.json  /usr/src/poc-innovation-iot-43fd0cc8a313.json

ENTRYPOINT ["/usr/src/app"]
