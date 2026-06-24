# archive site html in wayback machine
# API documentation: https://docs.google.com/document/d/1Nsv52MvSjbLb2PCpHlat0gkzw0EvtSgpKHu4mk0MnrA/

set MIN_AGE "30d"
set delay "2m"

set pwd (pwd)
cd ~/pnppl.cc/site/
for FILE in **/*.html
	echo -e "\n$FILE ..."
	echo "$FILE" | grep --extended-regexp --invert-match --quiet '(\+/(tag|thumb|search|date)/)|(app/(meander|nagi)/\.(p|http|ssl))|/_'
	if test "$status" -eq 0
		curl --request POST \
			--header "Accept: application/json" --header "Authorization: LOW $ARCHIVEORG_S3" \
			--data "url=https://pnppl.cc/$FILE&skip_first_archive=1&if_not_archived_within=$MIN_AGE&js_behavior_timeout=0&delay_wb_availability=1" \
			https://web.archive.org/save
		echo -e "\n ... sleeping $delay"
		sleep "$delay"
	else
		echo "... skipped"
	end
end
cd "$pwd"
