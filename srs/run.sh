#！/bin/bash

export CANDIDATE=$(ifconfig eth0|grep 'inet '|awk '{print $2}')
docker run --rm --env CANDIDATE=$CANDIDATE \
  -p $1:1935 -p $2:8080 -p $3:8088 -p $4:1985 -p $5:8000/udp \
  registry.cn-hangzhou.aliyuncs.com/ossrs/srs:5 \
  objs/srs -c conf/https.rtmp2rtc.conf
