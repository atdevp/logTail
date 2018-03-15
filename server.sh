#!/bin/bash

name=`basename $0`
case $1 in
    start)
        nohup ./log-shiper  /opt/wwwlogs/nginxlog/agif.access.log agif_one 10.13.89.97:9092 8301  10.16.10.21 >> ./log-shiper.log &
    ;;
    stop)
        ps auxwww| grep "agif_one"| grep -v grep | awk '{print $2}' | xargs kill -9
    ;;
    restart)
        ps auxwww| grep "agif_one"| grep -v grep | xargs kill -9
        nohup ./log-shiper  /opt/wwwlogs/nginxlog/agif.access.log agif_one 10.13.89.97:9092 8301  10.16.10.21 >> ./log-shiper.log &
    ;;
    *)
        echo "Usage: $name [start|stop|restart]"
        exit -1
    ;;
esac
exit 0
    