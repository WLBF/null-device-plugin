FROM golang:1.17 as build

WORKDIR /build
COPY . .

RUN CGO_ENABLED=0 go build ./cmd/null-device-plugin

FROM alpine:3.15

COPY --from=build /build/null-device-plugin /usr/bin/null-device-plugin

ENTRYPOINT ["null-device-plugin"]
