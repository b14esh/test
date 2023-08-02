#!/bin/bash

summa=0

myFunction()
{

     echo  "This is text from function"
     echo  "first parameter is: $1"
     echo  "Second parametr is: $2"
     summa=$(($1+$2))

}



myFunction 50 10
echo "Summa = $summa"
