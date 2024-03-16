count=$(find . -type f -o -type d | wc -l | sed 's/^[ \t]*//;s/[ \t]*$//')

echo $count
