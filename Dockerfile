FROM golang:1.17-alpine

ENV GOPROXY=https://goproxy.io,direct 
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/main.go

EXPOSE 8082

CMD ["./main"]
LABEL maintainer="MohammadAmin Rahimi <marcoding78@gmail.com>"
