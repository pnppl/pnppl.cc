git remote set-url --add --push origin ssh://git@git.gay/pnppl/pnppl.cc.git &&
git remote set-url --add --push origin ssh://git@github.com/pnppl/pnppl.cc.git &&
git remote set-url --add --push origin ssh://git@codeberg.org/pnppl/pnppl.cc.git &&
#metastore -a -m -f meta.store &&
fish set_mtimes.fish &&
echo '
.*
w10hosting_default/
\!*.*
xlog/public/style-min.css
site/.pagefind/
site/+/
site/1-bit.day/
site/about/
site/ai/
site/compat/
site/fs-ir/
site/pnppl/
site/public/
site/sitemap/
site/202*/
site/app/index.html
site/sitemap.xml
site/*.html' >> .git/info/exclude
