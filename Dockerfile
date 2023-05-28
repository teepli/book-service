FROM golang:alpine
RUN apk add build-base

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -ldflags '-s -w -extldflags "-static"' -o main .

EXPOSE 9000

CMD ["/build/main"]
