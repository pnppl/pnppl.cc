argparse 'c/commit=?' -- $argv &&
set msg "deploy.fish $(date --rfc-3339='seconds')" &&
if set -q _flag_commit
	set msg $_flag_commit
end
rm -rf site/public/ &&
fish build.fish &&
for file in (path filter -t dir (fdfind . site/)); chmod 755 $file; end
for file in (path filter -t file (fdfind . site/)); chmod 644 $file; end
git add . && git commit -m "$msg" &&
git pull &&
git push &&
lftp -e "set ftp:skey-force; mirror -R site/ /; exit" -u pnppl,$FTP_PASSWORD w10.host &&
#lftp -e "set ftp:skey-force; set ftp:ssl-protect-data yes; set ftp:ssl-protect-list yes; mirror -R site/ /; exit" -u pnppl,$FTP_PASSWORD w10.host &&
echo "! DEPLOY OK !"
