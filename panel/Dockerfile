FROM golang:1.22.5-alpine3.20 AS builder

WORKDIR /app

COPY go.* ./
ENV GOPROXY=https://goproxy.io,direct
RUN go mod download

COPY . ./

RUN go build -v -o server ./panel/main.go


FROM alpine:3.20

COPY --from=builder /app/server /app/server
COPY --from=builder /app/panel/api/templates /app/panel/api/templates
COPY --from=builder /app/panel/api/static /app/panel/api/static

ENV APP_ENV=docker

CMD [ "/app/server" ]
