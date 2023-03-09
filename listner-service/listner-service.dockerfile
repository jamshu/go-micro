#build a tiny docker image 

FROM alpine:latest

RUN mkdir /app 

COPY listnerApp /app

CMD [ "/app/listnerApp" ]