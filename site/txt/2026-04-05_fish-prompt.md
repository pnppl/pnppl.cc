# Fish greeting & prompt
#code #fish

```fish
function fish_greeting
	clear -x
	tput cup 0 $(math -s 0 $COLUMNS/2-4) # center
	tput dim
    echo -n "="
    tput sgr0
    tput bold
    echo -n "^."
    tput setaf 5
    echo -n "ꞈ"
    tput sgr0
    tput bold
    echo -n ".^"
    tput sgr0
    tput dim
    echo -n "= "
    tput sgr0
    tput bold
    echo "∫"
    tput sgr0
end
```

```fish
function fish_prompt --description 'Write out the prompt'
	#Save the return status of the previous command
	set stat $status

	if not set -q __fish_prompt_normal
		set -g __fish_prompt_normal (set_color normal)
	end

	if not set -q __fish_color_blue
		set -g __fish_color_blue (set_color -o blue)
	end

	set __fish_color_status (set_color -o green)
	if test $stat -gt 0
		set __fish_color_status (set_color -o red)
	end

	switch "$USER"
		case root toor
			if not set -q __fish_prompt_cwd
				if set -q fish_color_cwd_root
					set -g __fish_prompt_cwd (set_color $fish_color_cwd_root)
				else
					set -g __fish_prompt_cwd (set_color $fish_color_cwd)
				end
			end

			printf '%s %s%s%s%s%s%s %s%s %s%s%s' \
				(date +"%H:%M") \
				(set_color -o red) $USER "$__fish_color_status" "$stat" (set_color -o red) (prompt_hostname) \
				(set_color normal) (prompt_pwd) \
				(set_color -o red) "⎬ ⸨#〉" (set_color normal)

		case '*'
			if not set -q __fish_prompt_cwd
				set -g __fish_prompt_cwd (set_color $fish_color_cwd)
			end

 			# "♓︎•〉" is excellent but the pisces confuses xfce4-terminal. Good eyes: ·︢ ⊙ ⊚ ⸟ Gills: ⸨ ⟪ ｟ Tails: × » › } ⦄ ⎬ ≻
			printf '%s%s%s %s%s%s%s%s%s%s %s%s %s%s%s' \
				(set_color -d brblue) (date +"%H:%M") (set_color normal) \
				(set_color brcyan) $USER "$__fish_color_status" "$stat" (set_color normal) (set_color cyan) (prompt_hostname) \
				"$__fish_prompt_cwd" (prompt_pwd) \
				"$__fish_color_blue" "⎬ ⸨•〉" (set_color normal)
	end
end
```

![[Screenshot of terminal emulator. There's text art of a cat at the top. The prompt ends with text art of a fish.|../img/screenshots/fish-prompt.gif]]

```
                                   =^.ꞈ.^= ∫
12:34 pnppl0chonk ~ ⎬ ⸨•〉
```
