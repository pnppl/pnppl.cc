set pages "../meander/links.md"
set ignore "ignore.md"
set out "wander.js"

function parsein -a md
	for line in (cat "$md" | tail -n +3)
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
	return 0 ||
	return 1
end

echo "const wander = {
	// Prefer consoles with: links to posts, not site root; non-techy links; few neighbors; up-to-date Wander version
	consoles: [
		'https://antonio.is/wander/',
		'https://exurd.neocities.org/wander/',
		'https://heckmeck.de/wander/',
	],
	// More info and alternate browsing mode at https://pnppl.cc/app/meander
	pages: [
		'https://pnppl.cc/app/wcb/',
" > "$out" &&

parsein "$pages" &&

# TODO: include domains from ../meander/blocked.md
echo "	],
	ignore: [" >> "$out" &&

parsein "$ignore" &&

echo "	],
	styles: [
		// win9x style; display 'Open' on mobile
		'wander.css',
	],
	scripts: [
		// move 'About' into 'Console' menu, add button advertising Wander Console Builder
		'modify-menu.js',
	]
}" >> "$out"
