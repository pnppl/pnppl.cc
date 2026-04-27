# Generate a simple retro character creator/paper doll in Netscape 7.2-compatible HTML+CSS using GIMP and position: absolute

## Build layers in GIMP
- Install Batcher plugin: https://kamilburda.github.io/batcher/
- Create layer group for each doll part with options as layers
	- For example, layer group "hat" with layers "crown", "ballcap", "top hat"
	- Prefix the default option with ![0-9] to sort it to the top
	- Each layer in a group should have the same dimensions, big enough to encompass all layers in the group. (This way you can align each option differently, with transparent canvas padding, and align the entire group the same in the CSS.)
		- ie: make all layers in group visible and selected, rectangle select to encompass all of them, then Layer -> Resize Layers to Selection
	- Open sophie.xcf to see what I'm talking about
- Export the layers with File -> Export Layers
	- Check "Resize to layer size" (or else all parts will be the same size as the base)
	- Set Name to "[layer path]"
	- Export to img/

## Finish CSS (edit fish script)
- Align each group appropriately using absolute px offsets (left: x; top: y;)
	- This could probably be automated, but you only need to do it once and doesn't take long to do by hand if you don't have many parts
- Set pixel values for character width/height, media queries
- Change background colors
- Make a new border with Broider[^b]
- Etc. etc., lots of stuff to alter. The script just automates writing boilerplate so you can focus your time on GIMP, you'll probably want to make lots of changes to it.

## Generate HTML
- `fish gen.fish > doll.html`


# About Sophie
All assets are from Puzzle Castle, a children's video game based on the picture book of the same name, released as a CD-ROM for PC and Mac in 1996. I did not make them, just altered them. They were dumped with Media Station[^m] and some manual screenshooting. Plz no sue.

Code, such as it is, is AGPLv3.


[^b]: https://maxbittker.github.io/broider/
[^m]: https://github.com/npjg/MediaStation/