set total 0
set number 0
set pages

# actual size of file
function siz
	stat -c '%s' $argv[1]
end

# size in kb
function kb
	math --scale=0 $argv[1] / 1024
end

# percentage of 2000kb
function dv
	math --scale=0 --scale-mode=round $argv[1] / 1024 / 2000 \* 100
end

# add up html and embedded assets
for file in site/*/index.html
	set pagesize (siz $file)
	set assets (cat $file |
		string match --regex --all 'src="([^"]+)"' |
		# `string match -raq` isn't working and i don't know why
		grep -v 'src="' |
		# we are not counting assets shared by all pages since they should only have to load once
		grep -v 'badges' |
		grep -v 'logo')
	for uri in $assets
		set uri (string replace --all '..' '' "$uri" |
			string replace --all '&#43;' '+' |
			string unescape --style="url")
		set pagesize (math $pagesize + (siz "site$uri"))
	end
	set total (math $total + $pagesize)
	set pages $pages $pagesize
	++ number
end

# page weights
set pages (string split ' ' $pages | sort -n)
set median $pages[(math $number / 2)]
set average $(math $total / $number)

#echo "pages: $pages"
#echo "median: $median"
#echo "median kb: $(math $median / 1024)"
#echo "average: $average"
#echo "average kb: $(math $average / 1024)"

# shared assets
set prefix 'site/public'
set css (siz "$prefix/style-min.css")
set badge (siz "$prefix/badges/victor-8831-wide.gif")
set favicon (siz "$prefix/favicon.ico")
set logo (siz "$prefix/logo.gif")
set shared (math $css + $badge + $favicon + $logo)

# "extra" assets that might or might not be loaded
set fonts 0
for file in $prefix/font/{VictorMono-Medium.woff2,VictorMono-MediumItalic.woff2,death-note.woff2}
	set fonts (math $fonts + (siz $file))
end

set badges
for file in $prefix/badges/*.gif
	set badges $badges (siz $file)
end
set badges (string split ' ' $badges | sort -n)

# extra = the max additional asset load size
set extra (math $fonts + $badges[-1])
set combined (math $shared + $extra)

echo "
<style><!--
	#nut-wrap {
		width: 33%;
		float: right;
	}
	#nut {
		width: 100%;
		border-style: solid;
		border-spacing: 0;
#		padding: 0.25em;
		background-color: white;
	}
	#nut * {
		font-family: Helvetica, sans-serif !important;
		color: black;
		border-color: black;
	}
	#heading {
		line-height: 1;
		text-align: center;
	}
	#heading,
	#nut td {
		border-width: 0 0 1px;
		border-style: solid;
		padding: 2px 0.25em 2px 0.5em;
	}
	#percent,
	#nut td + td {
		text-align: right;
		font-weight: bold;
	}
	#nut .small {
		font-size: 0.8em;
	}
	#nut .last {
		border-bottom-width: 1em;
	}
	#nut .borderless {
		border: 0;
	}
	#nut .med {
		font-size: 2em;
		font-weight: bold;
		vertical-align: bottom;
		padding-left: 0.25em;
	}
	#nut .big {
		letter-spacing: -0.1ch;
		font-size: 2.5em;
		padding: 2px;
	}
	#nut span,
	#nut b {
		display: inline-block;
		transform: scaleX(1.1) translateX(2%);
		margin-right: 1ch;
	}
	#nut td + td b {
		margin-right: 0.5ch;
	}
	#nut .bytes-row {
		border-bottom-width: 8px;
		padding-bottom: 0;
	}
	#bytes {
		line-height: 1;
	}
	#nut #amount {
		line-height: 0;
		padding-top: 1em;
	}
	.size {
		font-size: 1.25em;
	}
	#nut .foot {
		line-height: 1;
		padding
	}
	#per {
		display: inline-block;
		margin-bottom: -0.5em;
	}
--></style>
<div id=nut-wrap>
<table id=nut border=2 cellspacing=0 cellpadding=2 width=33%>
	<tr>
		<th colspan=2 class=big id=heading><span>Nutrition Facts</span></th>
	</tr>
	<tr>
		<td colspan=2 class=borderless id=per><span>$number servings per site</span></td>
	</tr>
	<tr>
		<td class=last><b class=size>Serving size</b></td>
		<td class=last><b class=size>1 page</b></td>
	</tr>
	<tr>
		<td colspan=2 class=\"small borderless\" id=amount><b>Amount per serving</b></td>
	</tr>
	<tr>
		<td class=\"med bytes-row\"><span>Bytes<span></td> <!-- kilocalories... kilobytes... -->
		<td class=\"big bytes-row\" id=bytes>$(kb $median)</td>
	</tr>
	<tr id=percent>
		<td colspan=2 class=small><span>% Daily Value*</span></td>
	</tr>
	<tr>
		<td><b>Average Page</b> <span>$(kb $average)k</span></td>
		<td><b>$(dv $average)%</b></td>
	</tr>
	<tr>
		<td><b>Total Assets</b> <span>$(kb $combined)k</span></td>
		<td><b>$(dv $combined)%</b></td>
	</tr>
	<tr>
		<td><span>&nbsp;&nbsp;&nbsp;&nbsp;Shared Assets $(kb $shared)k</span></td>
		<td><b>$(dv $shared)%</b></td>
	</tr>
	<tr>
		<td><span>&nbsp;&nbsp;&nbsp;&nbsp;<i>Trans</i> Assets $(kb $extra)k</span></td>
		<td><b>$(dv $extra)%</b></td>
	</tr>
	<tr>
		<td><b>AI Slop</b> <span>0k</span></td>
		<td></td>
	</tr>
	<tr>
		<td><b>Cookies</b> <span>0k</span></td>
		<td></td>
	</tr>
	<tr>
		<td class=last><b>JavaScript</b> <span>0k**<span></td>
		<td class=last></td>
	</tr>
	<tr>
		<td colspan=2 class=\"small foot\">* The % Daily Value (DV) tells you how much a component in a serving of website contributes to a daily diet. 2,000 bytes a day is used for general nutrition advice.</td>
	</tr>
	<tr>
		<td colspan=2 class=\"small borderless foot\">** The search page requires JavaScript and WebAssembly. Some \"special\" pages optionally use a little.</td>
	</tr>
</table>
</div>
"

