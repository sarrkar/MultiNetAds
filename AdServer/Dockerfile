FROM golang:1.22.5
COPY . /app
WORKDIR /app

ENV GOPROXY https://goproxy.io,direct
RUN go build -o ad-server .
CMD ["./ad-server"]