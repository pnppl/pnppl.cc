# spaces in filenames are absolutely prohibited
rm mtimes.list &&
for file in (fdfind . site/)
	echo "$file $(path mtime $file)" >> mtimes.list
end
