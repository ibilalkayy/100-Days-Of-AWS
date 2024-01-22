FROM golang:1.21.4-alpine3.18
RUN mkdir /site
COPY . /site
WORKDIR /site
RUN go build -o main .
CMD [ "/site/main" ]