FROM alpine:3.10

# TODO - to be tested and
RUN apk --no-cache add ca-certificates
COPY cmd/main/gobuild /app/gobuild
RUN ls -la /
EXPOSE 8080
ENTRYPOINT ["./app/gobuild"]
