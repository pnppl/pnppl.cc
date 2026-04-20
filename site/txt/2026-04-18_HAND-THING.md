# HAND THING
#project

I've just created **[[HANDTHING.PARTY|//handthing.party/]]**, a website with the sole purpose of doing that hand thing.

I initially registered handth.ing, which was a very cool domain, but then I discovered the .ing TLD (top-level domain) is on the dread [[HSTS preload list|https://serverfault.com/a/1067232]], meaning it would be impossible to use for unsecured HTTP connections. Or at least, I think it would be... I'm not entirely sure how it all works. It seems like it's enforced by the browser, so old ones ought to be immune, but maybe there's something involved on the TLD's end too. In any case, if I try to access a site with HSTS enabled over HTTP (or HTTPS with an expired cert), my browser gives an error saying "nothing can be done, fuck you"; since HTTP support is important for retro compat, I steer clear of these TLDs.

Luckily, I thought to check the preload list about 15 minutes after I placed the order, and my registrar was willing to refund it. In the end it worked out for the best since DOT PARTY is way more fun and half the price.

I haven't tested it in old environments yet, thanks to the untimely death of an AC adapter, but I'd like for it to work on anything. Yes, it's a bit of a personal fixation, but it's also in keeping with the web 1.0 nature of [[Shaye Saint John's website|//web.archive.org/web/20170901151201if_/http://shayesaintjohn.net/]]. Pre-HTML5 video embedding is kind of ridiculous, but ffmpeg lets you convert to swf, and it's working on Netscape 4.03 with the [[Macromedia Flash 6|https://archive.org/details/MacromediaFlashMX]] plugin installed; version 5 played the audio but not video. I think I need a separate player file for <6? Also, I realized I can make it almost as garish without any CSS. It even kind of works in Mosaic; everything is static and there's no video, but like... the gifs are there! And... words! `<body background="pic.gif">`

Not much else to say about it. If you need me, I'll be [[doing that hand thing|//handthing.party/]].
