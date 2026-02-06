package main

import (
	// Core
	"context"

	"github.com/emad-elsaid/xlog"
	// All official extensions
	// the flag to disable them seemingly does nothing...
	_ "github.com/emad-elsaid/xlog/extensions/all"
	_ "github.com/emad-elsaid/xlog/extensions/autolink_pages"
	_ "github.com/emad-elsaid/xlog/extensions/custom_widget"
//	_ "github.com/emad-elsaid/xlog/extensions/date"
	_ "github.com/emad-elsaid/xlog/extensions/embed"
	_ "github.com/emad-elsaid/xlog/extensions/hashtags"
	_ "github.com/emad-elsaid/xlog/extensions/heading"
//	_ "github.com/emad-elsaid/xlog/extensions/hotreload"
	_ "github.com/emad-elsaid/xlog/extensions/images"
	_ "github.com/emad-elsaid/xlog/extensions/opengraph"
	_ "github.com/emad-elsaid/xlog/extensions/pandoc"
	_ "github.com/emad-elsaid/xlog/extensions/photos"
	_ "github.com/emad-elsaid/xlog/extensions/recent"
	_ "github.com/emad-elsaid/xlog/extensions/rss"
//	_ "github.com/emad-elsaid/xlog/extensions/shortcode"
// does urlset call home?
// unclear, but probably irrelevant -- only used by search engines
	_ "github.com/emad-elsaid/xlog/extensions/sitemap"
	_ "github.com/emad-elsaid/xlog/extensions/toc"
	_ "github.com/emad-elsaid/xlog/extensions/frontmatter"
// hideously lazy hack -- upload_file is actually frontmatter_hash
	_ "github.com/emad-elsaid/xlog/extensions/upload_file"
// disqus = all/alpha
//	_ "github.com/emad-elsaid/xlog/extensions/disqus"
	_ "github.com/emad-elsaid/xlog/extensions/search"
)

func main() {
	xlog.Start(context.Background())
}
