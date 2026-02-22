# super simple hack to get pages to stop linking to themselves
# leaves behind </a> but browsers just ignore it
# as written, should run in base build directory

for file in *.html
	set url (echo -e "from urllib import parse\nprint(parse.quote(\"$(path basename -E $file)\"))" | python3)
	sed -i -E "s/<a class=\"autolink\" href=\"\/$url\">(<span class=\"tag is-rounded \">[0-9]+\/[0-9]+<\/span> )?//g" "$file"
#	sed -i "s/<a class=autolink href=\"\/$url\">//g" "$file"
end

for file in */index.html
	set url (echo -e "from urllib import parse\nprint(parse.quote(\"$(path dirname $file)\"))" | python3)
	sed -i -E "s/<a class=\"autolink\" href=\"\/$url\">(<span class=\"tag is-rounded \">[0-9]+\/[0-9]+<\/span> )?//g" "$file"
#	sed -i "s/<a class=autolink href=\"\/$url\">//g" "$file"
end
