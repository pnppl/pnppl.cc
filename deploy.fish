argparse 'c/commit=?' 'm/mobile=?' -- $argv &&
set msg "deploy.fish $(date --rfc-3339='seconds')" &&
if set -q _flag_commit
	set msg $_flag_commit
end
if set -q _flag_mobile
	pkill crond &&
	rsync -avh --progress --update "$_flag_mobile" site/txt/
end

rm -rf site/.pagefind/ &&
rm -rf site/txt/!txt.zip &&
rm -rf site/img/1bitday/!1bitday.zip &&
rm -rf site/img/comics/!comics.zip &&
rm -rf site/img/photos/!photos.zip &&
fish build.fish &&
git stash -u &&
git pull &&
fish set_mtimes.fish &&
git stash pop -q &&
fish save_mtimes.fish &&
git add *.* &&
git add xlog/ &&
git add site/txt/ &&
git add site/img/ &&
git add site/app/ &&
git add site/humans.txt &&
git add site/favicon.ico &&
git commit -m "$msg" &&
git push &&

pagefind --site "site/" --output-subdir ".pagefind/" --root-selector "#main" --exclude-selectors "aside, .button, .buttons, .menu, #backlinks, #badge, #email, #footnotes, #see-also" --include-characters "#" --glob "*/*.{html}" --force-language "en" &&
zip -r site/txt/!txt.zip site/txt/ &&
zip -r site/img/1bitday/!1bitday.zip site/img/1bitday/ &&
zip -r site/img/comics/!comics.zip site/img/comics/ &&
zip -r site/img/photos/!photos.zip site/img/photos/ &&
chmod -R 755 site/ &&
for file in (fdfind -I -t f . site/); chmod 644 $file; end &&
lftp -e "set ftp:skey-force; mirror -R --parallel=20 --delete site/ /; exit" -u pnppl,$FTP_PASSWORD w10.host &&
echo "! DEPLOY OK !" ||
echo " !! ~~~~~~~ DEPLOY FAILED! ~~~~~~ !! "

if set -q _flag_mobile
	crond &
end
