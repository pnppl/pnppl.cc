set in "../meander/links.md"
set out "wander.js"

echo "const wander = {
	consoles: [
		'https://antonio.is/wander/',
		'https://douglascuthbertson.com/wander/',
		'https://heckmeck.de/wander/',
		'https://exurd.neocities.org/wander/',
	],
	// More info and alternate browsing mode at <https://pnppl.cc/app/meander>
	pages: [" > ../wander/wander.js &&
for line in (cat "$in" | tail -n +3)
	if test (string match -r '^\#' "$line")
		set line (string replace -a '#' '' "$line" | string replace -r ' $' '')
		echo "		//$line" >> "$out"
	else if string length -q "$line"
		set line (string replace -r '^http:' 'https:' "$line")
		echo "		'$line'," >> "$out"
	else
		echo >> "$out"
	end
end &&

echo "	],
	ignore: [
		// rationalists, racists, fascists
		'https://*.astralcodexten.com/',
		'https://dhh.dk/',
		'https://gwern.net/',
		'https://*.lesswrong.com/',
		'https://slatestarcodex.com/',
		'https://*.substack.com/',
		'https://*.yudkowsky.net/',
		'https://x.com/',

		// garbage silos
		'https://medium.com/',
		'https://*.substack.com/',

		// blocked by http headers
		'https://annas-archive.gl/', // the best site in the universe
		'https://codeberg.org/',
		'https://danielmiessler.com/',
		'https://hyperdoc.khinsen.net/',
		'https://maggieappleton.com',
		'https://neal.fun/',
		'https://*.otherstrangeness.com'.
		'https://thesweetbits.com/',

		// YC, HN, AI, capitalists
		'https://foundersatwork.posthaven.com/',
		'https://paulgraham.com/',
		'https://*.samaltman.com/',
		'https://simonwillison.net/',
		'https://stratechery.com/',
		'https://*.ycombinator.com',

		// I'm sure she's great but that banner drives me up the fucking wall
		'https://sachachua.com',

		// paywalled; zealot
		'https://www.wheresyoured.at/',

		// Consoles
		// sorry Joshes but my name is not Josh
		'https://joshing.you/wander/',
		// capitalist slop
		'https://www.davidtran.me/wander/wander.js',
	],
	referral: 'no',
	scripts: [
		// move 'About' into 'Console' menu
		'move-about.js',
	],
	styles: [
		// win9x style; display 'Open' on mobile
		'wander.css',
	]
}" >> "$out"
