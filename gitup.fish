git remote set-url --add --push origin ssh://git@git.gay/pnppl/pnppl.cc.git &&
git remote set-url --add --push origin ssh://git@github.com/pnppl/pnppl.cc.git &&
git remote set-url --add --push origin ssh://git@codeberg.org/pnppl/pnppl.cc.git &&
#metastore -a -m -f meta.store &&
fish set_mtimes.fish &&
echo ".*" >> .git/info/exclude &&
echo >> .git/info/exclude &&
echo "w10hosting_default" >> .git/info/exclude
