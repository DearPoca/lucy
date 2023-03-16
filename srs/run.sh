#！/bin/bash

export CANDIDATE=$(ifconfig eth0 | grep 'inet ' | awk '{print $2}')

RTMP_PORT=1935
NGINX_HTTP_PORT=8080
NGINX_HTTPS_PORT=8088
HTTP_API_PORT=1985
RTC_SERVER_PORT=8000

if [ $# -ge 1 ]; then
  RTMP_PORT=$1
fi
if [ $# -ge 2 ]; then
  NGINX_HTTP_PORT=$2
fi
if [ $# -ge 3 ]; then
  NGINX_HTTPS_PORT=$3
fi
if [ $# -ge 4 ]; then
  HTTP_API_PORT=$4
fi
if [ $# -ge 5 ]; then
  RTC_SERVER_PORT=$5
fi

echo RTMP_PORT: ${RTMP_PORT}
echo NGINX_HTTP_PORT: ${NGINX_HTTP_PORT}
echo NGINX_HTTPS_PORT: ${NGINX_HTTPS_PORT}
echo HTTP_API_PORT: ${HTTP_API_PORT}
echo RTC_SERVER_PORT: ${RTC_SERVER_PORT}

docker run --rm --env CANDIDATE=$CANDIDATE \
  -p ${RTMP_PORT}:1935 -p ${NGINX_HTTP_PORT}:8080 -p ${NGINX_HTTPS_PORT}:8088 -p ${HTTP_API_PORT}:1985 -p ${RTC_SERVER_PORT}:8000/udp \
  registry.cn-hangzhou.aliyuncs.com/ossrs/srs:5 \
  objs/srs -c conf/https.rtmp2rtc.conf
