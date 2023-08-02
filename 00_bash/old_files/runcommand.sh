#!!! выполняем команду ls если выведет больше 2 то выполнится команда  ls /etc/
if (( `ls /tmp/ | wc -l` > 2)); then true ; else false ; fi && ls /etc/
