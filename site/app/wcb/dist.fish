# merge files, remove site-specific stuff, and replace 'demo' for easy reuse

# indent
cat wcb.css | sed 's|^|\t\t|g' | sed 's|^\t\t$||g' > css.tmp &&
cat wcb.js | sed 's|^|\t\t|g' | sed 's|^\t\t$||g' > js.tmp &&

# clean and inline
cat index.html |
sed '/<link rel="stylesheet" href="wcb.css">/r css.tmp' |
sed 's|\t\t<link rel="stylesheet" href="wcb.css">||' |
sed 's|/[*] !start [*]/|<style><!--|' |
sed 's|/[*] !end [*]/|--></style>|' |
sed 's| src="wcb.js" async||' |
sed '/<script>/r js.tmp' |
sed 's/ | pnppl.cc//' |
sed 's|\t\t<!--#include file="meta.ssi"-->||' |
sed 's|demo|index|' > wcb.html

# cleanup
rm -f {css,js}.tmp
