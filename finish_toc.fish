for file in site/*.html
	set output
	if grep "<h1 id=\"Footnotes\">" "$file" &>/dev/null
		set -a output "<li><a href=\"#Footnotes\">Footnotes<\/a><\/li>"
	end
	if grep "<h1 id=\"Backlinks\">" "$file" &>/dev/null
		set -a output "<li><a href=\"#Backlinks\">Backlinks<\/a><\/li>"
	end
	if grep "<h1 id=\"see-also\">" "$file" &>/dev/null
		set -a output "<li><a href=\"#see-also\">See Also<\/a><\/li>"
	end
	if test -n "$output"
		sed -Ezi "s/<\/ol>\s+<\/details>/$output\n<\/ol>\n<\/details>/" "$file"
	end
end

for file in site/*/index.html
	set output
	if grep "<h1 id=\"Footnotes\">" "$file" &>/dev/null
		set -a output "<li><a href=\"#Footnotes\">Footnotes<\/a><\/li>"
	end
	if grep "<h1 id=\"Backlinks\">" "$file" &>/dev/null
		set -a output "<li><a href=\"#Backlinks\">Backlinks<\/a><\/li>"
	end
	if grep "<h1 id=\"see-also\">" "$file" &>/dev/null
		set -a output "<li><a href=\"#see-also\">See Also<\/a><\/li>"
	end
	if test -n "$output"
		sed -Ezi "s/<\/ol>\s+<\/details>/$output\n<\/ol>\n<\/details>/" "$file"
	end
end
