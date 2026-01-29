fish build.fish &&
git add . && git commit -m "deploy.fish $(date --rfc-3339='seconds')" && git push &&
echo "! DEPLOY OK !"