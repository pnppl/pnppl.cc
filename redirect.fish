set in site/txt/aliases.rec

# get list of pages to redirect
for page in (recsel -t Alias -e 'alias != ""' -P page $in | string join0 | string split0 -n)
	# capitalizing it so the builtin doesn't cause problems
	for Alias in (recsel -t Alias -e "page = '$page'" -P 'alias' $in)
		# space to hyphen
		set Alias (string replace -a ' ' '-' $Alias)
		# we don't want to generate a redirect if we already have a page there
		if ! test -e "site/$Alias/index.html"
			mkdir -p "site/$Alias"
			echo \
"<!DOCTYPE html>
<html>
<head>
	<meta http-equiv=\"Refresh\" content=\"0; url='/$page'\">
	<title>Redirecting to /$page... | pnppl.cc</title>
</head>
<body>
	<p><a href=\"/$page\">Redirecting to /$page...</a></p>
</body>
</html>" > "site/$Alias/index.shtml"
			# .shtml so we can tell it apart from the others (already using .htm for something else)
		end
	end
end
