cat links.md | grep '^h' > .links.txt &&
sed -i 's|http://|https://|g' .links.txt &&
fish ../nagi/generate.fish --in=".links.txt" --out=".p"
