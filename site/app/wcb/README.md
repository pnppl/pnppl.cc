# Wander Console Builder

## About
Wander Console Builder is a web application that makes it easier to create, customize, and update a [Wander](https://codeberg.org/susam/wander) console.

You can try it at <https://pnppl.cc/app/wcb/>.

**WCB does not attempt to be secure. It's hacky. It will `eval()` any file and `<script src>` any URL you give it. You have been warned.**


## Install

You can easily install WCB on your own site by saving [wcb.html](https://git.gay/pnppl/pnppl.cc/raw/branch/main/site/app/wcb/wcb.html) to your Wander directory. It will automatically load your console for editing by default. (You will still have to save the changes you make manually.)

You can see how this works on my console at <https://pnppl.cc/app/wander/wcb.html>.


## URL Parameters

WCB will automatically load any URL or path passed as a parameter using `?`. It works identically to loading a URL with the box on the page. You can omit the protocol.

For example, you can load my console by visiting <https://pnppl.cc/app/wcb?../wander/wander.js> (relative path) or <https://pnppl.cc/app/wcb?pnppl.cc/app/wander/wander.js> (URL without protocol) or <https://pnppl.cc/app/wcb?https://pnppl.cc/app/wander/wander.js> (full URL).

The nice thing about this is you can bookmark the URL for your console to easily update it without installing WCB on your site, or bookmark multiple consoles to edit with a single WCB install.


## License
Wander Console Builder is licensed [AGPLv3](LICENSE.md). It was made by pnppl.

Wander is licensed [MIT](LICENSE_for_demo.html.md). It was made by Susam Pal. It is included in the WCB repo as `demo.html`.
