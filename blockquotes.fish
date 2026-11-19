for file in site/*.html site/*/index.html
	sd -f m '<br>\n(<span class="mdash typog"></span><span class="is-hidden">---</span>.+)(</p>\n</blockquote>)' '$2\n<p class="quote-attr"><cite>$1</cite></p>' "$file"
end
