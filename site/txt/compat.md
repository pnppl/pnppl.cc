# Compatibility with pnppl.cc
#meta

If you encounter a compatibility issue, first try disabling CSS. If you still can't browse the site with CSS disabled, or the issue is with a feed reader, please hit the feedback button and let me know.

I'd also appreciate reports of the site working okay in older browsers with CSS enabled. Please only report CSS problems on legacy systems if you have an easy fix for them.

If you visit my site on something really cool, please send me a picture!


## LibreWolf 147, Firefox ESR 140 (Linux) & IronFox 147, Fennec 147 (Android)
These are the baseline modern (CSS-enabled) browsers I use. They should be fully supported.


## NCSA Mosaic 2.7b6, ~1996 (AppImage)
https://github.com/AppImageCommunity/NCSA-Mosaic-AppImage

This is the HTML-only browser I regularly test with. It originally came out in 1996, but this version has some more recent patches.


## Firefox 4.0, 2011-03 (Linux)
https://ftp.mozilla.org/pub/firefox/releases/4.0/linux-x86_64/en-US/

Supporting this is not a priority, but it actually renders pretty well thanks to a few compatibility hacks in the CSS.

## Firefox 12.0, 2012-04 (Windows XP)
Better than Firefox 4.0. Not perfect, but perfectly adequate.

## Mypal 74.1.1, 2025-07 (Windows XP)
Based on Firefox 74 from 2020-03.

As far as I can tell everything renders perfectly, and the search even works. And no performance issues (on period-correct hardware).


## Offpunk
https://offpunk.net/

Many pages don't render at all without forcing 'view full', which also has some issues:
- Logo takes up an entire page
- Footnote links in the text body don't render
- See Also, etc., sections don't render; the footnotes section also doesn't render reliably

Your best bet is probably to type 'feed' and use that for navigation.

I tried both 2.7.1 from the repo and the latest git version as of 2026-02. They behave quite differently. Installed depends also make a big difference (see these with 'version').


## Links2, Lynx, ELinks, w3m
Seems to pretty much work perfectly.


## Netscape 4.03, 1997-06 (Wine)
https://winworldpc.com/product/netscape-navigator/40x

Absolute clusterfuck until you disable CSS globally in the preferences ("Advanced" section), then it works fine.

My curiosity got the better of me and I started poking at this, thinking it might be tractable now that Netscape 7.2 looks ok. Nope. It looks slightly better, but it's still a lost cause.

There are some very odd alignment and spacing issues, but the most offensive problem is the garish colors. It interprets "inherit" as neon green. I discovered you can defeat this with var(), eg: `color: var(--ns, inherit)`. But then I would still have to set it to plain `inherit` first for browsers that don't understand variables, so this is not satisfactory. There is no way I would ever get this to look even as good as the bare HTML. This is the definition of unsupportable.

## Netscape 7.2, 2004-08 (Wine)
https://winworldpc.com/product/netscape-navigator/7x

Pretty janky and I couldn't find a way to disable CSS, but a surprising amount of stuff renders right. The whole site kinda just looks stretched out. Firefox 4 rendering is nice enough now that I've switched to this as my main unnecessary-CSS-devolution target.


## Internet Explorer
https://winworldpc.com/product/internet-explorer/

### Wine
#### 1.0 / 4.40.308
Had to extract the .cab; couldn't install. **Completely** broken. It throws a bunch of errors and opens random links in my default browser, including Spaceship.com where I have my domain name. After freaking out for a while, it opens web1.0hosting.net, the homepage of my wonderful host, which does render mostly fine. Trying to navigate to my site via pnppl.w10.site does the same thing. Eventually it crashes.

#### 3.0 / 4.70.1155
Had to extract .cab. Wine always opens the built-in iexplore, even when I change the name of the executable.

#### 2.0; 5.51.4807.2300; 8.00.6001.17184
Refuses to install.

#### Wine's iexplore
Basically works fine except for some minor issues. Not bad for a browser that nobody is actually meant to use.

Clearly I should try these on Windows.

### Windows XP Home SP3 x86
#### 1.5 for Windows NT i386
Actually behaves pretty much exactly like 1.0 on Wine, only it doesn't spam the URLs it tries to open into another browser.

#### 3.0 / 4.70.115
Refuses to install. Extracting cab does not help.

#### 5.51.4807.2300
Refuses to install. Extracting all cabs and running iexplore.exe succeeds! ...but then it fails to load any sites and soon crashes.

#### 7.0.5730.11
Broken, but actually better than IE 8, which supports enough CSS to fuck it up worse. Everything fits on the screen. A bigger issue is there's no obvious way to disable stylesheets. You can supply your own, for all pages, in `Tools -> Internet Options -> Accessibility`, but a blank one doesn't seem to unset the styles.

On the plus side, the built-in feed reader works great.

#### 8.0.6001.18702 (2009)
Pretty... damn bad. Issues with positioning and button sizing/text alignment. Stuff floats off the screen and over text. Not sure if you'd want to use CSS. Works fine without.


## Dillo, Dillo+
Pretty good rendering, but TOC gets messed up at small window sizes, and it doesn't seem to want to display anything in bold.

Unlike both historical and modern Firefox, it renders the UTF-8 sequences in .md files correctly without being told. (Modern Firefox has `View -> Repair Text Encoding`; Firefox 4 has `View -> Character Encoding -> Unicode`. The Netscapes need to be coerced like Firefox and then fail to render correctly. Of the text-only browsers, only Links2 fails. Mosaic, of course, also fails.)


## Netsurf 3.11 (Flatpak)
Technically it works well; all the buttons look like buttons, the TOC looks right (but won't collapse), the txt button isn't floating over text... but it feels quite janky because stuff isn't aligned properly. Also, I noticed the 1-bit.day gifs render blurry, so it must not support the rendering mode option. (Don't think I checked this on other browsers, aside from Mosaic which just displays them at crisp actual size.) I don't see any way to disable stylesheets. It supports JS through a config option; with this enbled, search still does not work and generates errors.

This is tangential, but this version of the browser has a goddamn hamburger menu instead of a menu bar. I am irrationally annoyed by this. I think 3.09 might have a proper interface. Building from source seems like a pain in the ass though.


## Kristall v0.4 (Linux)
https://kristall.random-projects.net/

Seems to work fine. Text only. Nice browser.


## Chawan 0.3.3 (Linux)
https://chawan.net/

Not too shabby. Buttons are kinda large and imposing, list markers aren't quite aligned, some z-index issues.

Very intuitive for a TUI browser. Good mouse support.


## Litebrowser (Linux)
https://github.com/litehtml/litebrowser-linux

Mostly there but definitely janky. Navbar is offset to the right, annotations render but float way above text, button text off-center. Still, an impressive showing, and I'd be tempted to try to improve compat if it didn't use that horrid new UI thing where there's no titlebar (gtk4?).


## MicroWeb (86Box)
https://github.com/jhhoward/MicroWeb

Works remarkably well, but has a pretty severe problem: it doesn't understand how to navigate to `/+/` (eg, `/+/recent`) when encoded as an HTML entity. I've been wondering if this would become an issue. It seems to struggle with character entities in general. Maybe I'd have better luck with /!/, or maybe I should do /0/ or something. It really shouldn't matter, but it's nice to have all that stuff in one place and have it sort to the top. In any case, this browser seems like a great way to test the super low end; looking forward to trying it on old hardware.


## Modern Chromium browsers
Ought to work fine but I don't check very often.

## Modern MacOS, iOS, and Windows browsers
Ought to work fine but I have no idea.
