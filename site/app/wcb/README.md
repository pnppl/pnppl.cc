# Wander Console Builder

**WCB does not attempt to be secure. It's hacky. It will `eval()` any file and `<script src>` any URL you give it. You have been warned.**


## Tips

I hope WCB is mostly self-explanatory. Here are tips for more advanced, non-obvious aspects.


### Parameters

WCB will automatically load any URL or path passed as a parameter. It works identically to loading a URL with the box on the page. You can omit the protocol.

For example, you can load my console by visiting <https://pnppl.cc/app/wcb?../wander/wander.js> or <https://pnppl.cc/app/wcb?pnppl.cc/app/wander/wander.js> or <https://pnppl.cc/app/wcb?https://pnppl.cc/app/wander/wander.js>.

The nice thing about this is you can bookmark the URL for your console to easily update it.


### Client-side

WCB runs entirely client-side. My host won't see what you're doing unless you pass a parameter, and you can download it or copy it to your website and it should work fine.

If you put WCB in the same directory as your Wander console you can use the theme builder on it. I'd like to make this simpler, but for now:

1. Rename `index.html` to `wcb.html` so you don't overwrite Wander.
2. Open `wcb.html`, search for `demo.html`, and change it to `index.html`. (`demo.html` is just the `index.html` that comes with Wander.)
3. Copy `wcb.{css,html,js}` to your Wander directory. You can also copy `butterfly.ico` for the favicon.

It won't automatically load your console, just display a preview in the theme builder. For now you can attach the parameter `?wander.js` if you want to automatically load it in.
