#!/bin/sh
TZ=Asia/Shanghai
echo_time() {
  date +"%Y-%m-%d-%H-%M-%S $*"
}

while true; do
  echo_time "rattle"
  sleep 3
done
