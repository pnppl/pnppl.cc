argparse 'c/commit=?' 'm/mobile=?' -- $argv &&
set msg "deploy.fish $(date --rfc-3339='seconds')" &&
if set -q _flag_commit
	set msg $_flag_commit
end
if set -q _flag_mobile
	rsync -avh --progress --update "$_flag_mobile" site/txt/
end

rm -rf site/.pagefind/ &&
fish build.fish &&
git add *.* &&
git add xlog/ &&
git add site/txt/ &&
git add site/img/ &&
git add site/humans.txt &&
git add site/favicon.ico &&
git stash &&
git pull &&
fish set_mtimes.fish &&
git stash pop -q &&
fish save_mtimes.fish &&
pagefind --site "site/" --output-subdir ".pagefind/" --force-language "en" &&
zip -r site/txt/!txt.zip site/txt/ -x \*.zip
zip -r site/img/1bitday/!1bitday.zip site/img/1bitday/ -x \*.zip
git commit -m "$msg" &&
git push &&

chmod -R 775 * &&
#for file in (path filter -t dir (fdfind . site/)); chmod 755 $file; end
#for file in (path filter -t file (fdfind . site/)); chmod 655 $file; end
lftp -e "set ftp:skey-force; mirror -R --delete site/ /; exit" -u pnppl,$FTP_PASSWORD w10.host &&
chmod -R -x * &&
echo "! DEPLOY OK !" ||
echo " !! ~~~~~~~ DEPLOY FAILED! ~~~~~~ !! "
