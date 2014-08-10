FROM hideyuki/go-daily-app-setup:latest
MAINTAINER Hideyuki Takei <takehide22@gmail.com>

WORKDIR /usr/daily/app/go-daily-app

# Update app codes
RUN git pull origin master

# Add app service on daemontools
RUN mkdir -p /etc/service/go-daily-app
RUN mkdir -p /etc/service/go-daily-app/log/main
ADD config/go-daily-app_run /etc/service/go-daily-app/run
RUN chmod +x /etc/service/go-daily-app/run
ADD config/go-daily-app_log_run /etc/service/go-daily-app/log/run
RUN chmod +x /etc/service/go-daily-app/log/run

EXPOSE 22 3000
CMD ["/usr/bin/svscan", "/etc/service/"]

