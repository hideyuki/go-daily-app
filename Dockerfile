FROM hideyuki/go-daily-app-setup:latest
MAINTAINER Hideyuki Takei <takehide22@gmail.com>

WORKDIR /usr/daily/app/go-daily-app

RUN git pull origin master

EXPOSE 22 3000
CMD ["go", "run", "server.go"]

