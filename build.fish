rm -rf site/public/ &&
rm -rf site/+/ &&
set pwd (pwd) &&
minify xlog/public/style.css -o xlog/public/style-min.css &&
sed -i -z 's/[\n\t]//g' xlog/public/style-min.css &&
cd xlog/cmd/xlog &&
go run xlog.go -source $pwd/site/txt -build $pwd/site -sitename pnppl.cc -rss.domain pnppl.cc -sitemap.domain pnppl.cc -og.domain pnppl.cc &&
cd $pwd &&
fish finish_toc.fish

if test $status -ne 0
	echo " !! ~~~~~~~ BUILD FAILED! ~~~~~~ !! "
	return 1
end
echo "! BUILD OK !"
return 0
