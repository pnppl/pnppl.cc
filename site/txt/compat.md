# Compatibility with pnppl.cc
#meta

If you encounter a compatibility issue, first try disabling CSS. If you still can't browse the site with CSS disabled, or the issue is with a feed reader, please hit the feedback button and let me know.

I'd also appreciate reports of the site working okay in older browsers with CSS enabled. Please only report CSS problems on legacy systems if you have an easy fix for them.

If you visit my site on something really cool, please send me a picture!


# LibreWolf 147, Firefox ESR 140 (Linux) & IronFox 147, Fennec 147 (Android)
These are the baseline modern (CSS-enabled) browsers I use. They should be fully supported.

# NCSA Mosaic 2.7b6, ~1996 (AppImage)
https://github.com/AppImageCommunity/NCSA-Mosaic-AppImage

This is the HTML-only browser I regularly test with. It originally came out in 1996, but this version has some more recent patches.

# Firefox 4.0, 2011-03 (Linux)
https://ftp.mozilla.org/pub/firefox/releases/4.0/linux-x86_64/en-US/

Supporting this is not a priority, but it actually renders pretty well thanks to a few compatibility hacks in the CSS.

# Offpunk
https://offpunk.net/

Many pages don't render at all without forcing 'view full', which also has some issues:
- Logo takes up an entire page
- Footnote links in the text body don't render
- See Also, etc., sections don't render; the footnotes section also doesn't render reliably

Your best bet is probably to type 'feed' and use that for navigation.

I tried both 2.7.1 from the repo and the latest git version as of 2026-02. They behave quite differently. Installed depends also make a big difference (see these with 'version').

# Links2, Lynx, ELinks, w3m
Seems to pretty much work perfectly.

# Netscape 4.03, 1997-06 (Wine)
https://winworldpc.com/product/netscape-navigator/40x

Absolute clusterfuck until you disable CSS globally in the preferences ("Advanced" section), then it works fine.

# Netscape 7.2, 2004-08 (Wine)
https://winworldpc.com/product/netscape-navigator/7x

Pretty janky and I couldn't find a way to disable CSS, but a surprising amount of stuff renders right. The whole site kinda just looks stretched out. Some elements run off the right side of the screen.

# Internet Explorer (Wine) #
https://winworldpc.com/product/internet-explorer/

### 1.0 / 4.40.308
Had to extract the .cab; couldn't install. **Completely** broken. It throws a bunch of errors and opens random links in my default browser, including Spaceship.com where I have my domain name. After freaking out for a while, it opens web1.0hosting.net, the homepage of my wonderful host, which does render mostly fine. Trying to navigate to my site via pnppl.w10.site does the same thing. Eventually it crashes.

### 3.0 / 4.70.1155
Had to extract .cab. Wine always opens the built-in iexplore, even when I change the name of the executable.

### 2.0; 5.51.4807.2300; 8.00.6001.17184
Refuses to install.

### Wine's iexplore
Basically works fine except for some minor issues. Not bad for a browser that nobody is actually meant to use.

Clearly I should try these on Windows.

# Modern Chromium browsers
Ought to work fine but I don't check very often.

# Modern MacOS, iOS, and Windows browsers
Ought to work fine but I have no idea.
