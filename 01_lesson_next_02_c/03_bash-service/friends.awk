#!/usr/bin/awk -f
BEGIN {
     print "Список друзей"
     year=0
}
{
     if ($2 == "1980") year=year+1
}
END {
   printf ("\tВсего записей:\t%d\n", NR)
   printf ("\tВ 1980 году родилось друзей:\t%d\n", year)
}

#Чтобы запустить сценарий, нужно сделать его исполнимым и передать ему файл друзей:
#$ chmod +x friends.awk
#$ ./friends.awk friends
