git remote set-url --add --push origin ssh://git@git.gay/pnppl/pnppl.cc.git &&
git remote set-url --add --push origin ssh://git@github.com/pnppl/pnppl.cc.git &&
git remote set-url --add --push origin ssh://git@codeberg.org/pnppl/pnppl.cc.git &&
#metastore -a -m -f meta.store &&
fish set_mtimes.fish &&
echo '
.*
w10hosting_default/
\!*.*
!site/app/*/!*.htm
!site/.gb.txt
xlog/public/style-min.css
site/+/
site/1-bit.day/
site/about/
site/ai/
site/compat/
site/deathnote/
site/fs-ir/
site/mirrors/
site/pnppl/
site/public/
site/sitemap/
site/202*/
site/app/index.html
site/sitemap.xml
site/*.html' >> .git/info/exclude
