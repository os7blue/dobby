FROM debian:11
CMD ['mkdir /data','mkdir /data/file']
COPY dobby '/data'
COPY ./view '/data/file'
COPY ./push.db '/data/file'


