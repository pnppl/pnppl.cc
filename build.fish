fish prep_title.fish &&
set pwd (pwd) &&
cd xlog/cmd/xlog &&
go run xlog.go -source $pwd/tmp -build $pwd/site -sitename pnppl.cc -rss.domain pnppl.cc -sitemap.domain pnppl.cc &&
cd $pwd &&
rm -rf tmp &&
cd site &&
fish $pwd/clean_autolinks.fish &&
cd $pwd &&
fish finish_toc.fish &&
echo "! BUILD OK !"