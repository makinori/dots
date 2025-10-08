#!/bin/bash

underline=$(tput smul)
normal=$(tput sgr0)

# use timezone in case traveling 
year=$(date +%y)
yearWeek=$(date +%W) # starting with monday 

echo -n "inject ${underline}"
if [ $((yearWeek%2)) -eq 1 ]; then
	echo -n "left";
else
	echo -n "right";
fi
echo " leg${normal}"

echo "on a ${underline}wednesday${normal}"

totalYearWeeks=$(date -d "$year-12-31" +%W)
echo "week $yearWeek/$totalYearWeeks"
