# initial stage
FROM golang:1.13-alpine AS builder
ADD . /go/src/Golang-Docker
WORKDIR /go/src/Golang-Docker
RUN apk add --update \
    curl \
    git
RUN curl https://glide.sh/get | sh
RUN glide install
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

# final stage
FROM alpine:3.10
RUN apk --no-cache add ca-certificates
COPY --from=builder main ./
COPY .env.prod ./
RUN mv .env.prod .env
RUN chmod +x ./main
ENTRYPOINT ["./main"]
EXPOSE 80 443