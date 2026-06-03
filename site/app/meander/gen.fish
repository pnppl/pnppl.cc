function htmlify
#	set temp "$LC_ALL" &&
#	export LC_ALL=C &&
	echo -n '' > $argv.html &&
	for line in (cat $argv.md)
		if test (string sub -l 2 "$line") = '# '
			echo "<h1>$line</h1>" >> $argv.html
		else if test (string sub -l 3 "$line") = '## '
			echo "<h2>$line</h2>" >> $argv.html
		else if test (string sub -l 3 "$line") = '###'
			echo "<h3>$line</h3>" >> $argv.html
		else if string length -q "$line"
			echo "<a href=\"$line\">$line</a><br>" >> $argv.html
		else
			echo '' >> $argv.html
		end
	end
#	export "LC_ALL=$temp"
end

cat links.md | grep '^h' > .links.txt &&
sed -i 's|http://|https://|g' .links.txt &&
fish ../nagi/generate.fish --in=".links.txt" --out=".p"
htmlify blocked
htmlify links
