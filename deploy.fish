argparse 'c/commit' -- $argv &&
set msg "deploy.fish $(date --rfc-3339='seconds')" &&
if set -q _flag_commit
	set msg $_flag_commit
end
fish build.fish &&
git add . && git commit -m "$msg" && git push &&
echo "! DEPLOY OK !"
