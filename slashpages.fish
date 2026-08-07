set out site/txt/slashpages.md

set slashpages (fdfind --no-ignore --type dir --exact-depth 1 --color never --path-separator '' --base-directory site . | string split0)

echo -e '# Slashpages\n#meta\n' > "$out" &&

echo '## Special Pages' >> "$out" &&
for file in (echo "$slashpages" | grep -vE '\+|img|public|txt|vid|pdf' | grep -vE '^20.+' | grep -vE '^$')
	echo "- [[$file]]" >> "$out"
end &&

echo -e '\n## Site Resources' >> "$out" &&
for file in + img public txt vid pdf
	echo "- [[$file]]" >> "$out"
end

#echo -e '\n## Dated Posts' >> "$out"
#for file in (string split ' ' "$slashpages" | grep -E '^20.+')
#	echo "- [[$file]]" >> "$out"
#end

if test "$status" -ne 0
	echo "! slashpages fucked up"
	return 1
end
return 0
