FROM golang:alpine

RUN mkdir /app
WORKDIR /app


COPY ./dobby /app
COPY ./push.db /app/resource/
COPY ./view /app/resource/view

ENV PORT=8702
EXPOSE $PORT

CMD ["./dobby"]
