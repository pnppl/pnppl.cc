set scratch_dir './.html_tmp'
set out 'site/txt/!html.zip'
rm -rf $scratch_dir &&
rm -rf $out &&
mkdir $scratch_dir &&
for file in site/**/*.html
	cp --parents -r $file $scratch_dir &&
	sd 'href="/' 'href="../' $scratch_dir/$file
end &&
cp --parents -r site/public/style-min.css $scratch_dir &&
cd $scratch_dir &&
rm site/error404.html &&
sd 'href="\.\./' 'href="' site/*.html &&
sd 'href="\.\./' 'href="../../' site/*/*/*.html &&
sd 'href="\.\./' 'href="../../../' site/*/*/*/*.html &&

zip -r ../$out site &&
cd ..

if test $status -ne 0
	echo "!~ HTML ZIP ERROR~!"
	set return 1
else
	echo "! html zip success !"
	set return 0
end

rm -rf $scratch_dir
return $return
