FROM alpine:latest AS syslog
RUN apk update && apk upgrade --available && apk add rsyslog
COPY conf/rsyslog.conf /etc/rsyslog.conf
ENTRYPOINT ["rsyslogd", "-n"]
