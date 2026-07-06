# Semantic HTML is a Myth
#tech #snark

`<center>` is to be avoided, because it mixes up style and content, which must be kept separate at all costs.

But `<p class="centered">` plus `centered { text-align: center; }` — now *that's* some clean code!

Likewise, you should always use `<em>` and `<strong>` instead of `<i>` and `<b>`. Someone might want to emphasize (or... strengthen?) parts of your content differently than with italics or boldface. In this scenario, they can't just style the `<i>` and the `<b>` exactly the same way one would style the "semantic" versions, for... reasons.

A good rule of thumb is that the more you have to type, the more semantic your HTML is. Hence why we can no longer use `<blink>` or `<marquee>` and instead must write god knows how many lines of CSS to accomplish the exact same result. Typing is a ritual we perform to sanctify our HTML. The more you type, the more blessed your website.
