function frame
	set url $argv[1]
	set headers (curl -sIL $url)
	if test (echo $headers | grep -iE "frame-(options|ancestors):? '?(deny|sameorigin|self|none)'?")
		echo $url >> blocked.txt
	end
end
