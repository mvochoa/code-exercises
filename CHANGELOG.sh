#!/bin/bash

echo "# Hackerrank" > README.md
echo "Soluciones a ejercicios de [https://www.hackerrank.com/](https://www.hackerrank.com/)" >> README.md
echo "" >> README.md

last=()
list=($(find . -name README.md -not -path '\.\/README.md' -print0 | sort -z | xargs -r0))
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
    echo "$space- [${name//-/ }]($CI_PROJECT_URL/tree/master${line/\./})" >> README.md
    last=(${names[@]})
done
