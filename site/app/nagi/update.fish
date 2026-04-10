set agent 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.10 Safari/605.1.1'

# arg: headers
function checkframeable
	if test (echo "$argv[1]" | grep -iE "frame-(options|ancestors):? '?(deny|sameorigin|self|none)'?")
		return 1
	end
	return 0
end

# arg: url root
function checkhttp
	set site "$argv[1]"
	set headers (curl -sI -A "$agent" --max-time 15 "http://$site")
    set response (echo "$headers" | awk 'NR==1 {print substr($2, 1, 1)}')
    if test "$response" -eq 2
        set html (curl -s -A "$agent" --max-time 15 "$site")
        if test ! (echo "$html" | grep -iE 'window\.location\.replace|http-equiv ?= ?"?Refresh"?')
			if checkframeable "$headers"
				return 0
			end
        end
    end
	return 1
end

# arg: url root
function checkssl
	set headers (curl -sI -A "$agent" --max-time 15 "http://$argv[1]")
	if test "$status" -eq 0 && checkframeable "$headers"
    	return 0
	end
	return 1
end

git add smallweb.txt &&
wget "https://github.com/kagisearch/smallweb/raw/refs/heads/main/smallweb.txt" -O smallweb.txt &&
set new (git diff -U0 -- smallweb.txt | grep "^+[^+]" | sed -E 's|^[+]||g')

for site in $new
	set root (string split --no-empty -f 2 '/' "$site")
	echo -e "site: $site\nroot: $root result:"
	if checkhttp "$root"
		echo "http pass"
	end
	if checkssl "$root"
		echo "ssl pass"
	end
	echo
end
