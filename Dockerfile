FROM alpine:latest

RUN mkdir /app
RUN mkdir /app/resource
WORKDIR /app


COPY ./dobby /app
COPY ./push.db /app/resource/push.db
COPY ./view /app/resource/view


ENV PORT=8702
EXPOSE $PORT

CMD ["./dobby"]