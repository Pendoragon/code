#!/bin/sh
TZ=Asia/Shanghai
echo_time() {
  date +"%Y-%m-%d-%H-%M-%S $*"
  echo `date +"%Y-%m-%d-%H-%M-%S $*"` >> /var/log/log
}

mkdir /var/log

while true; do
  echo_time "rattle"
  sleep 3
done
