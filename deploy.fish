argparse 'c/commit=?' 'm/mobile=?' 'n/nomod=?' -- $argv &&
set msg "deploy.fish $(date --rfc-3339='seconds')" &&
if set -q _flag_commit
	set msg $_flag_commit
end
if set -q _flag_mobile
	pkill crond &&
	rsync -avh --progress --update "$_flag_mobile" site/txt/
end

rm -rf site/.pagefind/ &&
#	site/txt/!txt.zip \
#	site/img/1bitday/!1bitday.zip \
#	site/img/comics/!comics.zip \
#	site/img/photos/!photos.zip &&
curl https://pnppl.cc/.gb.txt > site/.gb.txt &&
fish build.fish &&
git stash -u &&
git pull &&
fish set_mtimes.fish &&
git stash pop -q &&
if not set -q _flag_nomod
	fish save_mtimes.fish
end
git add *.* \
	epub/ \
	map/ \
	xlog/ \
	site/txt/ \
	site/img/ \
	site/app/ \
	site/vid/ \
	site/*.txt \
	site/favicon.ico &&
git commit -m "$msg" &&
git push &&

pagefind --site "site/" --output-subdir ".pagefind/" --root-selector "#main" --exclude-selectors "aside, .button, .buttons, .menu, .excerpt, #backlinks, #badge, #email, #footnotes, #see-also" --include-characters "#" --glob "*/*.{html}" --force-language "en" &&
zip -r -FS site/txt/!txt.zip site/txt/ -i \*.md &&
zip -r -FS site/img/1bitday/!1bitday.zip site/img/1bitday/ -i \*.gif
set imagetypes '*.gif' '*.jpg' '*.jpeg' '*.png'
zip -r -FS -0 site/img/comics/!comics.zip site/img/comics/ -i $imagetypes  &&
zip -r -FS -0 site/img/photos/!photos.zip site/img/photos/ -i $imagetypes &&
fish html.fish &&
for zip in site/**/*.zip
	zip $zip readme.txt
end &&
fish epub/epub.fish &&
# optimize
caesiumclt -R --lossless --same-folder-as-input site/+/thumb/ &&
ect -9 --strict -recurse site/+/thumb/ &&
for file in site/**/*.{zip,epub}
	ect -9 --strict -zip $file
end &&
fish map/sitemap.fish &&
chmod -R 755 site/ &&
for file in (fdfind -I -t f . site/); chmod 644 $file; end &&
lftp -e "set ftp:skey-force; mirror -R --parallel=20 --delete site/ /; exit" -u pnppl,$FTP_PASSWORD w10.host

if test $status -ne 0
	echo " !! ~~~~~~~ DEPLOY FAILED! ~~~~~~ !! "
	set return 1
else
	echo "! DEPLOY OK !"
	set return 0
end

if set -q _flag_mobile
	crond &
end

return $return
