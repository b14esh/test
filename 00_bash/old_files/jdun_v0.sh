#/bin/bash
x=0
y=10
while  true
 do 
  uptime
  x=$((x+1))
  y=$((y-1))
  echo ${x}
  echo ${y}
  sleep 1
done

