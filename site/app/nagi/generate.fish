set in ".smallweb-home.txt"
set out ".p"

set alphabet (string split '' '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz')
function base62
	set num "$argv[1]"
	set result ''
	if test $num -eq 0
		echo -n 0
	end
	while test $num -gt 0
		set remainder (math $num % 62)
		set index (math $remainder + 1)
		set result "$alphabet[$index]$result"
		set num (calc $num // 62)
	#		echo "remainder = $remainder; index = $index; result = $result; num = $num"
	end
	echo -n "$result"
end

function print_html
	set PREV (base62 $argv[1]) # INT
	set CURR $argv[2] # STRING
	set NEXT (base62 $argv[3]) # INT
	set INDEX (base62 $argv[4]) # INT
	echo -n "<!--#include file=\"!0.htm\"-->$PREV.htm<!--#include file=\"!1.htm\"-->$CURR\">open</a> | <a href=\"$NEXT.htm\">next</a></h1><iframe src=\"$CURR<!--#include file=\"!2.htm\"-->$CURR\">$CURR<!--#include file=\"!3.htm\"-->" > "$out/$INDEX.htm"
end

# randomize
sort -R "$in" > ".smallweb-shuffled.txt"
set in ".smallweb-shuffled.txt"

# handle first element
set last (math (count (cat $in)) - 1)
print_html $last (head -1 $in) 1 0
echo -n "1/$last"

# handle middle elements
set i 1
for line in (tail -n +2 $in)
	print_html (math $i - 1) "$line" (math $i + 1) $i
	set i (math $i + 1)
	echo -n " $i/$last"
end

# handle last element
print_html (math $last - 1) (tail -1 $in) 0 $last
echo -n " $last/$last"

for ssi in !*.htm
	cp "$ssi" ".p/$ssi"
end
