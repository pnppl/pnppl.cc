set in "../meander/links.md"
set out "wander.js"

echo "const wander = {
	// Other Wander consoles that visitors can reach from my console.
	consoles: [
	],

	// My favourite websites and pages I recommend to the Wander community.
	pages: [" > ../wander/wander.js &&
for url in (bat "$in" | grep '^h')
   echo "		'$url'," >> "$out"
end &&
echo "	],

	// Websites and consoles to ignore.  My console will never fetch
	// consoles or web pages whose URLs match the following patterns.
	ignore: [
		'https://*.substack.com/',
	],
	referral: 'no',
	styles: [
		'wander.css',
	]
}" >> "$out"
