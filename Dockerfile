FROM alpine:latest

RUN mkdir /app

COPY go-weather /app

CMD [ "/app/go-weather" ]