FROM golang:latest
MAINTAINER pkrishnath@gmail.com
RUN mkdir /app
ADD . /app/
WORKDIR /app
EXPOSE 8080
RUN go build -o main .
CMD ["/app/main"]