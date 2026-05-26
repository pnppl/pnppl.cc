set in "../meander/links.md"
set out "wander.js"

echo "const wander = {
	consoles: [
		'https://antonio.is/wander/',
		'https://douglascuthbertson.com/wander/',
		'https://heckmeck.de/wander/',
		'https://exurd.neocities.org/wander/',
	],
	pages: [" > ../wander/wander.js &&
for url in (bat "$in" | grep '^h' | sed 's|http://|https://|g')
   echo "		'$url'," >> "$out"
end &&
echo "	],
	ignore: [
		'https://*.substack.com/',
		'https://gwern.net/',
		'https://slatestarcodex.com/',
		'https://*.astralcodexten.com/',
		'https://*.lesswrong.com/',
		'https://*.yudkowsky.net/',
		'https://x.com/',
	],
	referral: 'no',
	scripts: [
		'move-about.js',
	],
	styles: [
		'wander.css',
	]
}" >> "$out"
