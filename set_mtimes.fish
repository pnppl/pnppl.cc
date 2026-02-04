for line in (cat mtimes.list)
	set item (string split " " $line) &&
	touch -m -c --date="@$item[2]" "$item[1]"
end
