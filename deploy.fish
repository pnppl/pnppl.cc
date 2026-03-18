argparse 'c/commit=?' 'm/mobile=?' -- $argv &&
set msg "deploy.fish $(date --rfc-3339='seconds')" &&
if set -q _flag_commit
	set msg $_flag_commit
end
if set -q _flag_mobile
	pkill crond &&
	rsync -avh --progress --update "$_flag_mobile" site/txt/
end

rm -rf site/.pagefind/ \
	site/txt/!txt.zip \
	site/img/1bitday/!1bitday.zip \
	site/img/comics/!comics.zip \
	site/img/photos/!photos.zip &&
fish build.fish &&
git stash -u &&
git pull &&
fish set_mtimes.fish &&
git stash pop -q &&
fish save_mtimes.fish &&
git add *.* \
	epub/ \
	xlog/ \
	site/txt/ \
	site/img/ \
	site/app/ \
	site/humans.txt \
	site/favicon.ico &&
git commit -m "$msg" &&
git push &&

pagefind --site "site/" --output-subdir ".pagefind/" --root-selector "#main" --exclude-selectors "aside, .button, .buttons, .menu, #backlinks, #badge, #email, #footnotes, #see-also" --include-characters "#" --glob "*/*.{html}" --force-language "en" &&
zip -r site/txt/!txt.zip site/txt/ -i \*.md &&
zip -r site/img/1bitday/!1bitday.zip site/img/1bitday/ -i \*.gif
set imagetypes '\*.gif' '\*.jpg' '\*.jpeg' '\*.png'
zip -r site/img/comics/!comics.zip site/img/comics/ -i $imagetypes  &&
zip -r site/img/photos/!photos.zip site/img/photos/ -i $imagetypes &&
fish epub/epub.fish
chmod -R 755 site/ &&
for file in (fdfind -I -t f . site/); chmod 644 $file; end &&
lftp -e "set ftp:skey-force; mirror -R --parallel=20 --delete site/ /; exit" -u pnppl,$FTP_PASSWORD w10.host &&
echo "! DEPLOY OK !" ||
echo " !! ~~~~~~~ DEPLOY FAILED! ~~~~~~ !! "

if set -q _flag_mobile
	crond &
end
