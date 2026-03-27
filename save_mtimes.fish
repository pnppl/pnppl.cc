# spaces in filenames are absolutely prohibited
rm mtimes.list &&
for file in (fdfind -t f . site/txt) (fdfind -t f . site/img)
	echo "$file $(path mtime $file)" >> mtimes.list
end
