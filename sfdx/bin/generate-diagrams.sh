#/usr/bin/env bash

cd "$(dirname "$0")/.." || exit

diagrams=$(find . -name '*.mermaid' -type f)
generated=0

for d in $diagrams; do
    path=$(dirname "$d")
    name=$(basename "$d")
    image="${name%.*}.png"
#    echo $name
#    echo $image
    if [ ! -f "$path/$image" ] || [ "$d" -nt "$path/$image" ]; then
        generated=1
        echo "Generating $path/$image"
        mmdc -i "$d" -o "$path/$image" -s 4
    fi
done

if [ "$generated" -eq 0 ]; then
    echo "All diagrams are up to date."
fi