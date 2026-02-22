# hideous shell kludge

rm -rf tmp
mkdir tmp
for file in site/txt/*
	set dest tmp/(path basename $file)
	echo (head -n1 "$file" | sed "s/^# //g" | sed "s/ #\$//g" | sed -z "s/\n//g") > "$dest" &&

#	set date (echo -n "$dest" | grep '[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9].*' - | head -c 10) &&
#	if [ -n "$date" ]
#		echo -n "<time class=\"is-hidden\" datetime=\"$date\">$date</time><br class=\"is-hidden\">" >> "$dest"
#	end &&

	cat "$file" >> "$dest" &&
	touch -d (date -Rr "$file") "$dest" ||
	echo "!!! error in prep_title !!!"
end