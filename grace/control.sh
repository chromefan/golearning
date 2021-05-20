#!/bin/bash

APP_NAME=grace

psid=0
checkpid() {
   pids=`ps -ef | grep $APP_NAME | grep -v grep | awk '{print $2}'`
   arr=${pids//\n/ }
   echo "目前已有的pid为: ${arr[*]}"
   num=${#arr[@]}
   echo "目前已有的pid数量为: $num"
   if (($num > 1)) ; then
      psid=${arr[1]}
   elif (($num == 1)); then
      psid=${arr[0]}
   else
      psid=0
   fi
   echo "目前pid数量为: $psid"
}

status() {
   checkpid
   if [[ $psid -ne 0 ]];  then
      echo "$APP_NAME is running! (pid=$psid)"
      exit 0
   else
      echo "$APP_NAME is not running"
   fi
  exit 1
}

start() {
   checkpid
   if [[ $psid -ne 0 ]]; then
      echo "================================"
      echo "warn: $APP_NAME already started! (pid=$psid)"
      echo "================================"
      exit 0
   else
      echo -n "Starting $APP_NAME ..."
      nohup ./$APP_NAME &> /dev/null 2>&1 &
      checkpid
      if [[ $psid -ne 0 ]]; then
         echo "(pid=$psid) [OK]"
         exit 0
      else
         echo "[Failed]"
         exit 1
      fi
   fi
   exit 1
}

stop() {
   checkpid
   if [[ $psid -ne 0 ]]; then
      echo -n "Stopping $APP_NAME ...(pid=$psid) "
      kill -USR2 $psid
      old_pid=$psid
      sleep 5
      checkpid
      if [[ $psid -eq $old_pid ]]; then
         echo "failed"
      else
         echo "ok"
         exit 0
      fi
   else
      echo "================================"
      echo "warn: $APP_NAME is not running"
      echo "================================"
   fi
}
case "$1" in
   'start')
      start
      ;;
   'stop')
     stop
     ;;
   'restart')
     stop
     start
     ;;
   'status')
     status
     ;;
  *)
echo "Usage: $0 {start|stop|restart|status}"
  exit 1
esac
  exit 0
