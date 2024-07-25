FROM golang:1.22.5
COPY . /app
WORKDIR /app

ENV GOPROXY https://goproxy.io,direct
RUN go build -o event-server .
CMD ["./event-server"]