#!/bin/bash

CURDIR=$(cd $(dirname ${BASH_SOURCE[0]}); pwd )

logfile="/opt/wwwlogs/nginxlog/agif.access.log"
topic="raw-nginx-agif-new"
broker="127.0.0.1:9092,127.0.0.2:9092"
ip="127.0.0.1"
http_port="5202"

case $1 in
    start)
        nohup ${CURDIR}/log-shiper -f ${logfile} -t ${topic} -b ${broker} -a ${ip} -p ${http_port} > ${CURDIR}/collect-${topic}.log 2>&1  &
    ;;
    stop)
        ps auxwww| grep "${topic}"| grep -v grep | awk '{print $2}' | xargs kill -9
    ;;
    restart)
        ps auxwww| grep "${topic}"| grep -v grep | awk '{print $2}' | xargs kill -9
        nohup ${CURDIR}/log-shiper -f ${logfile} -t ${topic} -b ${broker} -a ${ip} -p ${http_port} > ${CURDIR}/collect-${topic}.log 2>&1  &
    ;;
    *)
        echo "Usage: $0 [start|stop|restart]"
        exit -1
    ;;
esac
exit 0
    
