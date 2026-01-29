# About pnppl.cc
Welcome to my [[https://en.wikipedia.org/wiki/Commonplace_book|commonplace]]. It's also, perhaps, my digital garden? I won't be offended if you just call it my blog.

I've been a prodigious note-taker for years. Being able to carry my notes with me and search them digitally means I can actually refer back to them, and often do. It recently occurred to me that some of these notes (when tidied up a bit) might be of interest to other people. I call this site a commonplace because it's an attempt to translate that existing practice into a public form.

I also write stuff that I don't refer to but is hopefully valuable in other ways: fiction, essays, code, etc. And since I'm at it anyway, I might as well post art I make that doesn't fit in a text file.

I hope that sharing my commonplace publicly will encourage me to write more, more intentionally, for a wider audience. It would be especially nice if it made me actually document my tech projects so I don't have to re-figure them out whenever I come back to them after a long time. I said I was a prodigious note-taker, not a good one.


# Copyleft
This site is licensed [[https://creativecommons.org/licenses/by-sa/4.0/|CC BY-SA 4.0]], with the attribution requirement waived.

Intellectual property is a fake idea; unfortunately, we live in a world where it reigns supreme. Copyleft gives us a tool to turn copyright against itself. It means that you can do whatever you want with the content of this site, as long as you extend the same right to others.

I'm not wild about this specific license. It's irritating that it requires notice of attribution and derivatives. I don't care what you do with my work as long as you don't [[https://en.wikipedia.org/wiki/Enclosure|enclose]] it. Still, CC BY-SA 4.0 seems to be on strong legal footing and is compatible with (and can even be converted into) the GPLv3. If the license causes issues for you, let me know and I'll probably be happy to do a public domain declaration.


# Install Victor Mono
This site really looks much better in [[https://rubjo.github.io/victor-mono/|Victor Mono]]. It's my favorite font. Even though it's monospace, it's a pleasure to read.

I could force it on you, but I find that rude. I block remote fonts in my browser: they're usually hosted on Google's CDN, so they call home; they're a multi-megabyte blob you have to download just to read some text; and they don't respect your preference of default font. So, please install it. I promise you'll like it.

Unfortunately, if you're on an Android phone, I think you have to be rooted to install a font and it's a big pain in the ass.


# We have to go deeper: technical details, etc. #
My goal is to have to change as little as possible about my existing approach to writing: plaintext, formatted in whatever way feels right, written in my notes app or shell. I'd like for the software to adapt to me, not the other way around. My one concession is Markdown. I don't really like it, but it's cornered the market, and it's close enough that I'm willing to bend a little. Extensions make it much less painful — particularly Wikilink syntax support. The text files that pnppl.cc is built from are the One True Commonplace, with the website providing some extra features that are nice to have but not required. You can access the original text file for a post with the _src_ button in the top right corner.

You can also access them through [[https://git.gay/pnppl/pnppl.cc/|the]] [[https://codeberg.org/pnppl/pnppl.cc|git]] [[https://github.com/pnppl/pnppl.cc|repo]], which handily doubles as a version history and complete source of the site. I'd like to make it more convenient, but for now you can use `git` to grab all of the posts at once.
```sh
git clone -n --depth=1 --filter=tree:0 https://git.gay/pnppl/pnppl.cc.git &&
cd pnppl.cc &&
git sparse-checkout set --no-cone /site/txt &&
git checkout
```

pnppl.cc does not use any Javascript, although there might be a search page in the future that uses it. The software this site is based on actually has search as an option, but it fell prey to my elimination of Javascript.  It's a bit silly that the public version of my notebook, which only became indispensible to me due to the search feature, does not actually have a search right now. It's just important to me that my site reflect what I hope to see online. For the time being, if you want to search, you can grab the text files and search them locally. This is easily accomplished in the shell, eg `grep -i "search term" *.md`. I like [[https://github.com/BurntSushi/ripgrep|ripgrep]].

A major goal I have is for the site to be usable as plain HTML. The sites that have stood the test of time, that still display perfectly and are a pleasure to read, are all plain old HTML. (Usually hosted on URLs like `http://lab.cs.mit.edu/faculty/~johnsmith` and miraculously preserved in their entirety by the Wayback Machine.) As a fan of retrocomputing, I'd also like to be able to access my site on anything with a web browser, and trying to accomplish that level of compatibility is the stuff of web design nightmares... unless it's just plain old HTML. To that end, I've been testing the site in progress on NCSA Mosaic v2.7b6. Sure, it's the last version... but it's the last version of the *first* major browser, and [[https://github.com/AppImageCommunity/NCSA-Mosaic-AppImage|I can run it on my computer.]] I figure if it works in Mosaic, it ought to work in anything. HTML is wonderfully resilient. Pretty much every browser ever made happily displays malformed HTML without issue. Modern browsers still render supposedly obsolete tags just fine, and Mosaic will just ignore tags it doesn't understand. 

With legacy compatibility covered, we can use stylesheets to make the site pretty on a modern browser, since CSS doesn't make them slow to a crawl or execute arbitrary code. Now, I *hate* web design, and I find CSS particularly frustrating, so I've relied on the stylesheet that xlog came with as a starting point. It uses a big framework and is much larger than I'd like at 200K (uncompressed — was originally 4x the size, compressed), but it handles dark mode, responsiveness, etc., which all sounds exhausting.

I selected [[https://github.com/emad-elsaid/xlog|xlog]] as the best static site generator for my requirements: everything I cared about still worked with Javascript disabled and it supported must-have features: backlinks and mentions. (The way mentions work is that when the name of a post is used in another post, it's automatically turned into a link. Backlink support adds a section at the foot of a post that lists all the other posts that link to it.) If I'm going hypertext, I want all the benefits, but I also don't want to spend a lot of time fussing over the website when it's the words that really matter. And given the wide variety of subjects, it would be nice to have some ability to discover posts. 

For mentions, the underline is broken with dashes to de-emphasize it and so you can distinguish it from a manual link. External links are followed by a % sign, because it kind of looks like an X and also kind of like the little external link symbol a bunch of sites use. I wanted to keep it ASCII when possible, as well as webfont-free. HTML entities are fair game though, provided they work in Mosaic, which is where the section (&sect;) and not (&not; — returns you to the table of contents) come from that appear when you hover a heading. 

Still, I'm not quite willing to give up UTF-8 for my posts. UTF-8 is so good. I'm addicted to the emdash. Sometimes I even use an emoji. I figure a garbled character here and there in the unlikely event that someone actually uses Mosaic or another legacy browser is an acceptable sacrifice. I may change my mind if I end up spending a lot of time reading on a legacy system, or if anyone actually reads this and cares. It wouldn't be too hard to change. I also like the idea of making it available via Gopher for the true low-tech heads. We'll see.
