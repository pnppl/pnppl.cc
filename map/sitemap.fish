# generate human-legible sitemap
# (the 'sitemap' extension generates sitemap.xml for search engines)
# yes this is quite ugly and the html fragments will need to be updated manually
# the difficulty is we need to run this at the very end of the build to ensure everything we will actually upload is listed

set in site
set out site/sitemap/index.html
set html map/index.html

cat $html.1 > $out &&
# pre AND tt seems to guarantee we get preserved whitespace and monospace text
echo -n '<pre class="sitemap"><tt class="sitemap">' >> $out &&
tree --charset=ascii --dirsfirst --noreport --hintro=/dev/null --houtro=/dev/null -ha -H -'' $in |
# move size to end
sd '(\[.+\]).+(<a.+)$' '$2 $1' |
# remove padding from size
sd '\[[^\d]*([\d]+[A-Z]?)\]' '[$1]' |
sd '<a href="/">site</a>' '<a href="/">pnppl.cc</a>' |
# why does it print so many breaks
sd '<br>' '' |
# remove indent
sd '\t' '' >> $out &&
echo '</tt></pre>' >> $out &&
cat $html.2 >> $out

if test $status -ne 0
	echo "!~~ SITEMAP FAIL ~~!"
	return 1
end
echo "! sitemap built !"
return 0


# hard spaces
#sd '   ' '&nbsp;&nbsp;&nbsp;' |
#sd ' ([\[<`|])' '&nbsp;$1' |
# hard tabs
#sd '[\t]' '&nbsp;&nbsp;&nbsp;&nbsp;' >> $out
