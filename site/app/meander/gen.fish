cat links.md | grep '^h' > .links.txt &&
sed -i 's|http://|https://|g' .links.txt &&
fish ../nagi/generate.fish --in=".links.txt" --out=".p"

set temp "$LC_ALL" &&
export LC_ALL=C &&
echo -n '' > blocked.html &&
for line in (cat blocked.txt)
	if test (string sub -l 3 "$line") = '## '
		echo "<h1>$line</h1>" >> blocked.html
	else if test (string sub -l 3 "$line") = '###'
		echo "<h2>$line</h2>" >> blocked.html
	else if string length -q "$line"
		echo "<a href=\"$line\">$line</a><br>" >> blocked.html
	else
		echo '' >> blocked.html
	end
end
export "LC_ALL=$temp"
