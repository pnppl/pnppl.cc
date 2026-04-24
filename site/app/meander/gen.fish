cat links.md | grep '^h' > .links.txt &&
sed -i 's|http://|https://|g' .links.txt &&
fish ../nagi/generate.fish --in=".links.txt" --out=".p"

export LC_ALL=C &&
for line in (sort blocked.txt)
	echo "<a href=\"$line\">$line</a><br>" >> blocked.html
end
