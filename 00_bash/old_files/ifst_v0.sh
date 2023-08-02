#! /bin/sh
# Пример iftst_v0
if test $# -ne 2; then
 echo "Команде должно быть передано ровно два параметра!"
 exit 1 
else
 echo "Параметр 1: $1. Параметр 2: $2" 
fi
