# Fish: get containing folder
#code #fish #snippet

I think `path` is one of the best things about Fish. Here's a simple function to get the canonical name of the folder something is in.

```fish
function folder
	if ! isatty stdin
		cat - | read -at argv
	end
	for file in $argv
		echo (path resolve $file | path dirname | string split '/')[-1]
	end
end
```

Use it like this:
```shell
folder file.ext
echo file | folder
folder */*.*
folder .
```