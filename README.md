# logging-kills-app

## environment
* test setup created on archlinux

## tools needed
* prometheus, started with systemd
* grafana, started with systmed
* logstash, started manually `/usr/bin/logstash -f /etc/logstash/logstash.conf`

all tools installed with pacman

## start go apps

## generate 2000 events in logstash queue
* `for i in {1..2000}; do curl -X PUT http://localhost:8080 -d '{"bla":"blo"}'; echo " $i"; done`

## logstash queue location on disk
* `/usr/share/logstash/data/queue/main`

## logstash get node stats
* `curl -XGET 'localhost:9600/_node/stats/pipelines/?pretty'`
* prometheus converter: https://github.com/alxrem/prometheus-logstash-exporter, start: `/usr/local/bin/logstash-exporter`, query: http://localhost:9198/metrics
