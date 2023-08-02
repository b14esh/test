#!/bin/bash
a="`ls -al | wc -l`"
echo ${a}
 
if (( a > 2 )); then
    true
    echo "a > 2"
else
    false
    echo "a<2"
fi
