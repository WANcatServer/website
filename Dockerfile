FROM ubuntu:latest
RUN apt-get update && apt-get install nginx php-fpm -y
EXPOSE 80


