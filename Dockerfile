FROM golang:1.17-alpine as builder

WORKDIR /app
RUN --mount=type=bind,target=/app GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o /dp4cli.exe dp4cli.go


FROM debian:buster-slim as wine

RUN dpkg --add-architecture i386
RUN apt update \
 && apt install -y --no-install-recommends wine wine32 \
 && rm -fr /var/lib/apt/lists/*

ENV WINEDEBUG=-all
# noop to create .wine
RUN wine cmd /c echo


FROM wine

COPY DP4CAPI.dll /
COPY --from=builder /dp4cli.exe /
