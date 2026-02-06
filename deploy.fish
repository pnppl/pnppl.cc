argparse 'c/commit=?' 'm/mobile=?' -- $argv &&
set msg "deploy.fish $(date --rfc-3339='seconds')" &&
if set -q _flag_commit
	set msg $_flag_commit
end
if set -q _flag_mobile
	rsync -avh --progress --update "$_flag_mobile" site/txt/
end
rm -rf site/public/ &&
rm -rf site/+/search/ &&
fish build.fish &&
git stash -u &&
git pull &&
fish set_mtimes.fish &&
git stash pop -q &&
fish save_mtimes.fish &&
npx pagefind --site "site/" --output-path "site/+/search/pagefind/" --force-language "en" &&
git add . &&
git commit -m "$msg" &&
git push &&
for file in (path filter -t dir (fdfind . site/)); chmod 755 $file; end
for file in (path filter -t file (fdfind . site/)); chmod 655 $file; end
lftp -e "set ftp:skey-force; mirror -R site/ /; exit" -u pnppl,$FTP_PASSWORD w10.host &&
#lftp -e "set ftp:skey-force; set ftp:ssl-protect-data yes; set ftp:ssl-protect-list yes; mirror -R site/ /; exit" -u pnppl,$FTP_PASSWORD w10.host &&
echo "! DEPLOY OK !" ||
echo " !! ~~~~~~~ DEPLOY FAILED! ~~~~~~ !! "
