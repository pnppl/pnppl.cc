# remove feeds from urls
set out '.smallweb-home.txt'
rm -f $out
wget "https://github.com/kagisearch/smallweb/raw/refs/heads/main/smallweb.txt" -O .smallweb.txt
for line in (bat .smallweb.txt)
	set split (string split --no-empty -f 2 '/' $line)
#	if test $split = "http:/" -o $split = "https:/"
#		echo $line >> $out
#	else
		echo "$split" >> $out
#	end
end

#sed -i 's|https://|//|g' "$out"
#sed -i 's|http://|//|g' "$out"
