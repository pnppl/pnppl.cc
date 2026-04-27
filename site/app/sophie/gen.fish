# background colors
set colors lightskyblue pink white plum teal black

# unitless h/w attr helper
function print_size
	set size (string split -f3 ', ' (file -b "$argv"))
	set width (string split -f1 ' x ' "$size")
	set height (string split -f2 ' x ' "$size")
	echo -n "\" width=\"$width\" height=\"$height"
end

# bg color helper
function print_colors_css
	set i 0
	for color in $colors
		echo "		label[for=bgcolor-$i].bgcolor { background-color: $color; }"
		set i (math $i + 1)
	end
end

# the whole old script shoved into a function and made slightly reusable (for bg colors)
function print_list_html
	set folder "$argv[1]"
	set total "$argv[2]"
	set img_tag "$argv[3]" # whether we print img tag (for printing colors)
	set pics (string split -n ' ' "$argv[4..]")

	# leading <input>
	echo -n '		<input type="radio" name="'
	echo -n "$folder"
	echo -n '" class="'
	echo -n "$folder"
	echo -n '" id="'
	echo -n "$folder-0"
	if test "$total" -eq 1
		echo '" checked>'
	else
		echo '">'
	end

	# <label><input> pairs
	set i 1
	for img in $pics[2..]
		echo -n '		<label for="'
		echo -n "$folder-$i"
		echo -n '" class="'
		echo -n "$folder"
		echo -n '" title="'
		echo -n "$folder"
		if test "$img_tag" -eq 1
			echo -n '"><img src="'
			echo -n "$img"
			echo -n '" alt="'
			echo -n "$folder"
			print_size "$img"
		end
		echo '"></label>'

		echo -n '		<input type="radio" name="'
		echo -n "$folder"
		echo -n '" class="'
		echo -n "$folder"
		echo -n '" id="'
		echo -n "$folder-$i"
		set i (math $i + 1)
		if test "$i" -eq "$total"
			echo '" checked>'
		else
			echo '">'
		end
	end

	# trailing <label>
	# (the last <label> is for the first <input>, so that they loop)
	echo -n '		<label for="'
	echo -n "$folder-0"
	echo -n '" class="'
	echo -n "$folder"
	echo -n '" title="'
	echo -n "$folder"
	if test "$img_tag" -eq 1
		echo -n '"><img src="'
		echo -n "$pics[1]"
		echo -n '" alt="'
		echo -n "$folder"
		print_size "$pics[1]"
	end
	echo '"></label>'
	echo
end

# main
echo '<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="Content-Type" content="text/html;charset=utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<meta name="referrer" content="no-referrer">
		<link rel="icon" href="favicon.ico">
		<meta property="og:site_name" content="pnppl.cc">
		<meta property="og:title" content="Build-A-Sophie">
		<meta property="og:description" content="Fan-made character creator for Puzzle Castle (1996)">
		<meta property="og:image" content="https://pnppl.cc/app/sophie/img/!2base/!2base-!lowered-light.gif">
		<meta property="og:url" content="https://pnppl.cc/app/sophie">
		<meta property="og:type" content="website">
		<meta name="description" content="Fan-made character creator for Puzzle Castle (1996). Listen to MIDIs and dress up Sophie.">
		<title>Build-A-Sophie | pnppl.cc</title>
		<style><!--
		html {
			cursor: url("shield.cur"), auto;
			height: 100%;
			width: 100%;
			image-rendering: pixelated;
			background-color: lightskyblue;
		}
		body {
			margin: 0;
		}
		input {
			position: absolute;
			left: -99999px;
			top: -99999px;
		}
		input:checked + label {
			display: inline;
		}
		a {
			position: absolute;
			right: 1px;
			top: 1px;
			opacity: 33%;
		}
		#character {
			position: relative;
			width: 300px;
			height: 300px;
			margin: auto;
		    border-image:  url("border.gif") 7 /  7px / 0 round;
    		border-width:  7px;
    		border-style:  solid;
		}
		label {
			display: none;
			cursor: pointer;
			cursor: url("sword.cur"), pointer;
			position: absolute;
			line-height: 0;
		}

		/* vvv --- pixel offsets --- vvv */
		label.body {
			left: 149px;
			top: 68px;
		}
		label.tights {
			left: 177px;
			top: 182px;
		}
		label.skirt {
			left: 165px;
			top: 152px;
		}
		label.hair {
			left: 171px;
			top: 83px;
		}
		label.hat {
			left: 154px;
			top: 48px;
		}
		label.shoes-left {
			left: 156px;
			top: 213px;
		}
		label.shoes-right {
			left: 185px;
			top: 211px;
		}
		label.necklace {
			left: 179px;
			top: 116px;
		}
		label.pet {
			left: 0;
			top: 53px;
		}
		/* ^^^ --- pixel offsets --- ^^^ */

		#bgcolor,
		label.bgcolor {
			width: 100%;
			height: 100%;
			position: absolute;
			top: 0;
			left: 0;
		}'

print_colors_css

echo '		audio {
			position: absolute;
			width: 100%;
			height: 20px;
			bottom: 0;
			display: none;
		}
		#music:checked ~ audio {
			display: block;
		}
		label#turntable {
			display: block;
			left: 1px;
			top: 1px;
			position: absolute;
			opacity: 33%;
		}
		#music:checked ~ label#turntable {
			opacity: 100%;
		}

		/* vvv --- queries --- vvv */
		@media screen and (max-width: 313px) {
			#character {
				border: 0;
			}
			audio,
			object,
			embed,
			label#turntable {
				display: none;
			}
		}
		@media screen and (min-width: 666px) and (min-height: 629px) {
			#character {
				transform: scale(2);
				transform-origin: 50% 0;
			}
		}
		--></style>
	</head>
	<body>
	<h2 style="display:none">This page requires CSS.</h2>
	<div id="bgcolor">'

print_list_html "bgcolor" (string split ' ' "$colors" | count) 0 "$colors"

echo '	</div>
	<div id="character">'

for folder in img/*/
	set pics "$folder"/*.gif
	set folder (path basename "$folder" | string replace -r '^![0-9]?' '')
	set total (string split -n ' ' "$pics" | count)
	print_list_html "$folder" "$total" 1 $pics
end

echo '		</div>
		<input type="checkbox" name="music" id="music">
		<label for="music" id="turntable"><img src="turntable.gif" width="32" height="17" alt="turntable; show/hide music" title="show/hide music"></label>
		<audio controls autoplay loop playsinline>
			<source src="music/concat.mp3" type="audio/mpeg">
			<source src="music/concat.ogg" type="audio/ogg">
			<source src="music/concat.aac" type="audio/aac">
			<source src="music/concat.wav" type="audio/wav">
			<object data="music/concat.swf">
				<embed src="music/concat.swf">
				<p><a href="music/concat.mp3">Download background music</a></p>
			</object>
		</audio>
		<a href="/2026-04-26_sophie">?</a>
	</body>
</html>'
