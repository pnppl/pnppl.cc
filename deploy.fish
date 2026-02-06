argparse 'c/commit=?' 'm/mobile=?' -- $argv &&
set msg "deploy.fish $(date --rfc-3339='seconds')" &&
if set -q _flag_commit
	set msg $_flag_commit
end
if set -q _flag_mobile
	rsync -avh --progress --update "$_flag_mobile" site/txt/
end

rm -rf site/public/ &&
rm -rf site/.pagefind/ &&
#rm -f site/txt/!txt.zip &&
#rm -f site/img/1bitday/!1bitday.zip &&

fish build.fish &&
git stash -u &&
git pull &&
fish set_mtimes.fish &&
git stash pop -q &&
fish save_mtimes.fish &&
pagefind --site "site/" --output-subdir ".pagefind/" --force-language "en" &&
zip -r site/txt/!txt.zip site/txt/ -x \*.zip
zip -r site/img/1bitday/!1bitday.zip site/img/1bitday/ -x \*.zip
git add . &&
git commit -m "$msg" &&
git push &&

for file in (path filter -t dir (fdfind . site/)); chmod 755 $file; end
for file in (path filter -t file (fdfind . site/)); chmod 655 $file; end
lftp -e "set ftp:skey-force; mirror -R --delete site/ /; exit" -u pnppl,$FTP_PASSWORD w10.host &&
echo "! DEPLOY OK !" ||
echo " !! ~~~~~~~ DEPLOY FAILED! ~~~~~~ !! "
