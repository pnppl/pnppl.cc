# 1-bit.day
#pixel-art #visual-art #code #fav

!1bitday ../img/1bitday

You can grab these all at once here: [[../img/1bitday/!1bitday.zip]]

~~I make them on [[1-bit.day|https://1-bit.day/pnppl]].~~ 1-bit.day got abandoned shortly after I took up using it; it's now been totally killed off. I still make B&W 16x16 pixel art and add it here.

I've started experimenting with color versions. I suppose these would be 3-bit.day since they use up to eight colors.

!1bitday ../img/1bitday/color

---

Here's the Python script I wrote to convert the files spit out by 1-bit.day:
```python
#!/usr/bin/env python3
# turns 1024x1024 pixel art from 1-bit.day into actual size
# if filename contains 8x8 or 32x32 it will switch from default 16x16; override this by passing arg
# args: filename format=gif size=16/auto
# dumps to stdout

from sys import argv, stdout
from PIL import Image, ImageOps

output = "gif"
size = 16
if len(argv) < 2:
	print("args: filename format=gif size=16/auto")
	exit(1)
image = Image.open(argv[1])
if "32x32" in argv[1]:
	size = 32
elif "8x8" in argv[1]:
	size = 8
if len(argv) > 2:
	output = argv[2]
if len(argv) > 3:
	size = int(argv[3])
image.resize((size, size), Image.Resampling.NEAREST, (256, 256, 768, 768)).save(stdout.buffer, format=output)
```

I used it like this:
```sh
1bitday.py pixel-art-16x16-1_1.png > out.gif
```
