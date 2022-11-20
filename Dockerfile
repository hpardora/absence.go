# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o absencer cmd/main.go

FROM ubuntu:latest

#Install Cron
RUN apt-get update
RUN apt-get -y install cron ca-certificates

RUN update-ca-certificates

RUN touch /var/log/cron.log

RUN echo "0 7 * * * bash /bin/absencer scheduler > /dev/stdout" > /etc/cron.d/absencer
RUN chmod 0644 /etc/cron.d/absencer

COPY --from=build /app/absencer /bin/absencer

CMD cron && tail -f /var/log/cron.log