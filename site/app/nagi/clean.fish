# remove feeds from urls
set out '.smallweb-home.txt'
rm -f $out
wget "https://github.com/kagisearch/smallweb/raw/refs/heads/main/smallweb.txt" -O .smallweb.txt
for line in (bat .smallweb.txt)
	set split (string split --no-empty -f 1,2 '/' $line)
	echo "$split[1]//$split[2]" >> $out
end
