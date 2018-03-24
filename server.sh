#!/bin/bash

CURDIR=$(cd $(dirname ${BASH_SOURCE[0]}); pwd )

logfile="/opt/wwwlogs/nginxlog/agif.access.log"
topic="raw-nginx-agif-new"
broker="10.18.75.121:9092,10.18.75.122:9092,10.18.75.123:9092,10.18.75.124:9092,10.18.75.125:9092,10.18.75.126:9092,10.18.75.127:9092,10.18.75.128:9092,10.18.75.129:9092,10.18.75.130:9092"
ip="10.13.89.92"
http_port="15202"

case $1 in
    start)
        cd 
        nohup ${CURDIR}/log-shiper  ${logfile} ${topic} ${broker}  ${ip} ${http_port} >> ${CURDIR}/collect-${topic}.log 2>&1  &
    ;;
    stop)
        ps auxwww| grep "${topic}"| grep -v grep | awk '{print $2}' | xargs kill -9
    ;;
    restart)
        ps auxwww| grep "${topic}"| grep -v grep | awk '{print $2}' | xargs kill -9
        nohup ${CURDIR}/log-shiper  ${logfile} ${topic} ${broker}  ${ip} ${http_port} >> ${CURDIR}/collect-${topic}.log 2>&1  &
    ;;
    *)
        echo "Usage: $0 [start|stop|restart]"
        exit -1
    ;;
esac
exit 0
    
