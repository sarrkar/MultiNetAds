FROM golang:1.22.5-alpine3.20 AS builder

WORKDIR /app

COPY go.* ./
ENV GOPROXY=https://goproxy.io,direct
RUN go mod download

COPY . ./

RUN go build -v -o server ./eventserver/main.go


FROM alpine:3.20

COPY --from=builder /app/server /app/server

ENV APP_ENV=docker

CMD [ "/app/server" ]
