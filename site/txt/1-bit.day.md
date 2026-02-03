# 1-bit.day #
#pixel-art

I make 1-bit (black and white) 16x16 art here:
https://1-bit.day/pnppl

You can get it as actual 16x16 bitmaps (gif) here: [[/img/1bitday]]

Here's the Python script I made to do the conversion.
```python
#!/usr/bin/env python3
# turns 1024x1024 pixel art from 1-bit.day into actual size
# if filename contains 8x8 or 32x32 it will switch from default 16x16; override this by passing arg
# args: filename format=gif size=16/auto
# decided it's more useful if it just dumps to stdout

from sys import argv, stdout
from PIL import Image, ImageOps

output = "gif"
size = 16
appendix = ""
if len(argv) < 2:
	print("args: filename format=gif size=16/auto") # appendix")
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
