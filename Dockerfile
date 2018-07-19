FROM alpine:latest

WORKDIR /app

EXPOSE 5000

COPY ./bin/ /app/

ENTRYPOINT ["/app/w4w.a"]