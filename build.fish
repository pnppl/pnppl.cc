# clean up auto redirects before generating slashpage
find site -type f -name '*.shtml' -delete &&
find site -type d -empty -delete &&
fish slashpages.fish &&
rm -rf site/public/ &&
rm -rf site/+/ &&
set pwd (pwd) &&
minify xlog/public/style.css |
sed -z 's/[\n\t]//g' > xlog/public/style-min.css &&
cd xlog/cmd/xlog &&
go run xlog.go -source $pwd/site/txt -build $pwd/site -sitename pnppl.cc -rss.domain pnppl.cc -sitemap.domain pnppl.cc -og.domain pnppl.cc &&
cd $pwd &&
fish finish_toc.fish &&
fish blockquotes.fish &&
fish redirect.fish

if test $status -ne 0
	echo " !! ~~~~~~~ BUILD FAILED! ~~~~~~ !! "
	return 1
end
echo "! BUILD OK !"
return 0
