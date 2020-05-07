#!/bin/bash

printf "# Code Exercises\n\n" > README.md
printf "Code exercises solution\n\n" >> README.md

last=()
list=($(ls -t $(find . -name README.md -not -path '\.\/README.md') | xargs -r0))
len=${#list[@]}
for (( numLine=0; numLine<${len}; numLine++ ));
do
    line=$(dirname ${list[$numLine]})
    names=(${line//\// })
    namesLength=${#names[@]}

    space=""
    for (( i=1; i<${namesLength}-1; i++ ));
    do
        if [ "${names[$i]}" != "${last[$i]}" ]
        then
            echo "$space- **${names[$i]//-/ }**" >> README.md
        fi
        space="$space  "
    done

    dir=${names[$namesLength-1]}
    name="$(tr '[:lower:]' '[:upper:]' <<< ${dir:0:1})${dir:1}"
    name=${name//_/: }
    date=$(git log --format=%cd --date=format:%d\ %b\ %Y $line | tail -n 1)
    dateRelative=$(git --no-pager log --format=%cd --date=relative $line | tail -n 1)
    echo "$space- [${name//-/ }](${line/\.\//}) - ***$date*** *$dateRelative*" >> README.md
    last=(${names[@]})
done
