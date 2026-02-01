argparse 'c/commit=?' -- $argv &&
set msg "deploy.fish $(date --rfc-3339='seconds')" &&
if set -q _flag_commit
	set msg $_flag_commit
end
git pull &&
rm -rf site/public/ &&
fish build.fish &&
chmod -R 755 site/ &&
git add . && git commit -m "$msg" && git push &&
lftp -e "set ftp:skey-force; set ftp:ssl-protect-data yes; set ftp:ssl-protect-list yes; mirror -R site/ /" -u pnppl,$FTP_PASSWORD w10.host &&
#lftp -e "mirror -R site/ /" -u pnppl,$FTP_PASSWORD w10.host &&
echo "! DEPLOY OK !"
