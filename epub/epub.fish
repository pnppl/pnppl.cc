set scratch_dir './.epub_tmp'
set out 'site/txt/!pnppl.epub'
#set excluded 404 index app 1-bit.day fs-ir visual-art
rm -rf $scratch_dir
mkdir $scratch_dir &&
for file in site/txt/20*.md
#	if ! contains (path basename -E $file) $excluded
		set title (string trim (head -1 $file))
		set date (string sub -e 10 (path basename $file))
		set filename $scratch_dir/$(path basename $file)
		# has no title, create one
		if test (string sub -l 2 $title) != "# "
			set title (string shorten -m 70 $title)
			echo "# $title" >> $filename &&
			echo -e "<span class=\"date\">$date</span>\n" >> $filename &&
			cat $file >> $filename
		# has title, don't duplicate it
		else
			echo "$title" >> $filename &&
			echo -e "<span class=\"date\">$date</span>\n" >> $filename &&
			tail -n +2 $file >> $filename
		end
		# my shruggie!! (this is just \ -> \\)
		sed -i 's/\\\\/\\\\\\\\/g' $filename
		# fix footnotes without fucking everything up with pandoc's inscrutable --file-scope
#		sed -i -E "s/\[\^([[:digit:]]+)\]/\[\^$(path basename -E $file)-\1\]/g" $filename
		sed -i -E "s/\[\^/\[\^$(path basename -E $file)-/g" $filename
		# remove A/V embeds
		sed -i -E 's/!\[\[.+(mp4|mkv|webm|mp3|wav|flac|ogg|ogv|oga|m4a)\]\]//g' $filename
#	end
end
set markdown_features markdown-smart+ascii_identifiers+space_in_atx_header+backtick_code_blocks+fenced_code_attributes+pipe_tables+strikeout+footnotes+lists_without_preceding_blankline+hard_line_breaks+autolink_bare_uris+wikilinks_title_before_pipe
#pandoc -o $out -f $markdown_features --toc --reference-location=section --metadata title="pnppl.cc" --epub-metadata="epub/epub.xml" --epub-title-page=false --css="epub/epub.css" --lua-filter="epub/chapters.lua" --resource-path=site/txt $scratch_dir/*.md &&
pandoc -o $out -f $markdown_features --toc --reference-location=section --metadata title="pnppl.cc" --epub-metadata="epub/epub.xml" --epub-title-page=true --css="epub/epub.css" --resource-path=site/txt $scratch_dir/*.md &&
if test $status -ne 0
	echo "!!----- EPUB FUCKED UP -----!!"
	set return 1
else
	echo "! epub success !"
	set return 0
end
if test (du $out | cut -f1) -ge 1024
	echo "!!~~ WARNING: epub larger than 1MB ~~!!"
end
rm -rf $scratch_dir
return $return
