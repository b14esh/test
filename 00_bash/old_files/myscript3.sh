#!/bin/bash
echo "superping $0"
echo "Нужно передать параметры!!"
echo "Первый параметр что пингуем ping   "
echo "Второй что traceroute"
echo "Третий  что host "
echo ">>>>>>>>>>>>>>>>Первый пошол $1"
ping -c 4 $1
echo ">>>>>>>>>>>>>>>>Воторй пошол $2"
traceroute  $2
echo ">>>>>>>>>>>>>>>>Третий пошол $3"
host $3
echo ">>>>>>>>>>>>>>>>Done"
