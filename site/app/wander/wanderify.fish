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
		'https://gwern.net/',
		'https://*.lesswrong.com/',
		'https://slatestarcodex.com/',
		'https://*.substack.com/',
		'https://*.yudkowsky.net/',
		'https://x.com/',

		// garbage silos
		'https://medium.com/',
		'https://*.substack.com/',

		// I'm sure she's great but that banner drives me up the fucking wall
		'https://sachachua.com',

		// Consoles
		// sorry Joshes but my name is not Josh
		'https://joshing.you/wander/',
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
