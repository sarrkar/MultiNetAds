FROM golang:1.22.5
COPY . /app
WORKDIR /app

ENV GOPROXY https://goproxy.io,direct
RUN go build -o publisher-site .
CMD ["./publisher-site"]