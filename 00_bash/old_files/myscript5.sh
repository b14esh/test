#!/bin/bash
read -p "Enter something Vasya, Trump, Petya, * : " y

if [ "$y" == "Vasya" ]; then
                       echo "Privet $y"
elif [ "$y" == "Trump" ]; then
                          echo "Hello $y"
else echo "Zadarova $y"
fi

echo "Please enter something"
echo "1, 2-9, Petya, *"
read x



echo "Starting CASE selection......."
case $x in
        1) echo "This is one";;
        [2-9]) echo "Two-Nine";;
        "Petya") echo  "Privet $x";;
        *) echo "Parammetr Unknown, sorry!"
esac
