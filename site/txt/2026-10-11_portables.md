# Portable Computers
#collection #hw #tech #list

This is just an inventory right now. Excludes modern phones, tablets, and laptops (~2010+) as well as game consoles, e-readers, mp3 players, synthesizers...

## Linux
- Sharp Zaurus SL5500 (2002)
	- In box with crap
	- The first Linux handheld IIRC
	- 128MB CF
	- Cradle
- PocketCHIP (2014)
	- In box
	- 3D printed keyboard which swivels on wires is not really usable due to lack of key labels

## #Palm OS (and GEOS)
### PDAs
- AST GridPad 2390 (1993)
	- Runs GEOS — handwriting software is by Palm
	- aka Tandy Zoomer/Z-PDA
	- In box
	- Bad digitizer, can't even get through initial setup
- Palm IIIx (1998)
	- `x` means it has the expansion slot
	- In box with crap - cradle, manuals, CDs - basically new
	- Kodak PalmPix camera
	- Palm Portable Keyboard (metal one)
	- There's something wrong with this, I think the digitizer
- Handspring Visor (1999)
	- Black soft case
	- I think this is the Deluxe version since it's transparent
	- Eyemodule camera
	- Magellan GPS Companion
		- Not sure if this works in present day
	- GameFace joystick adapter
		- Completely ridiculous
		- Maybe for a different version of the Visor? Can't get it to quite line up right and there's a notch that looks like it should be on the other side so you can push the power button
- Palm m125 (2001)
	- Flipcover
	- Removable batteries and an SD slot
	- NiMHs don't quite fit even though it looks like they do, need a shim
		- Also https://palmdb.net/app/battery-panel
- Palm Zire m150 (2002)
	- Silver hard case
	- Was mom's
- AlphaSmart Dana Wireless (2003)
	- Palm OS 4
	- Modified so that it runs on and recharges NiMH batteries
	- Two SD slots
- Sony Clié PEG-SJ22/U (2003) ×2
	- Palm OS 4
	- PEGA-GC10 Game Controller
		- Actual controller, plugs into proprietary connector and needs driver
		- In box with CD and manual
		- In its own little carrying case
		- Haven't tried it
	- Memory Stick Duo to SD adapter, SD to Micro SD adapter, and 2GB Micro SD
	- One of the cables with the weird connector
	- Two of the CDs
	- The one with the worn out Up button was mine in childhood
- Palm Zire 72s (2004)
	- Blue jelly case
	- Palm OS 4
	- Built-in camera
	- I think the `s` just means silver
- Palm TX (2005)
	- Palm OS 5 - last version
	- Got this brand new for like $50, still for sale on eBay
- Pixter Multimedia (2005) with Dmitry Grinberg rePalm board (2025)
	- Giant branded case that the GameCube-blue coating all flaked off of
	- One of the buttons is pretty worn out

### Accessories
- LandWare goVox voice recorder
	- aka i-Vox
	- Standalone voice recorder that runs on button cells and is designed to replace the flip cover
	- The part where it clips on is broken off. I think it's for Palm V?
	- Works (as well as it ever did)
	- Possibly the dumbest thing I've ever bought
- Margi Presenter-to-Go SD to VGA adapter
	- In box with all the crap including remote and carrying case
	- Don't remember if I've tried this
- Palm Portable Keyboard P10802U
	- Metal
	- I think this is the Palm Universal Connector which the m125 has
- Palm Wireless Keyboard 3245WW (Bluetooth)
	- Plastic
	- Have definitely used this on non-Palm devices
- Universal Wireless Keyboard by Palm 3169WWZ
	- Plastic
	- Unfolds super weird
	- CD with drivers for Palm OS, Windows Mobile
	- Connects via IR on an arm that swings out
	- Does not offer any indication that it is on, you just type and hope for the best

## PC
- Compaq LTE Elite 4/75CX - laptop (1994)
	- Same model grandma had
	- I think I fucked up the floppy drive with the cleaning disk
		- Uses one of those flat flex cables so I'm not sure if I can replace it with a Gotek
		- There is this but the store is closed: https://silvestron.com/openflops-w1d/
			- There's one on eBay for too much but it doesn't have the cable; doesn't seem to be unpluggable from drive, would have to desolder
	- HDD works and ought to be compatible with the SD adapter, but haven't gotten it to boot from SD yet
	- I reseated the keyboard bezel wrong somehow
	- Latches that hold the lid closed are broken
	- Battery was 12V NiMH so we could replace it with 10× AA
	- CMOS battery is NiCd and still holds a charge
	- https://revspace.nl/Compaq_LTE_Elite_4/50E
- Datavue 25 - lunchbox luggable
	- Branded blue soft case
	- Modem
	- Left alkalines in the keyboard at one point so this might need some TLC
- NEC Multispeed EL - laptop (1997)
	- Works but currently in pieces
		- Seemed dead until I found the dipswitches
		- Replacing one of the drives with a Gotek emulator
		- Replacing the NiCd with NiMH (or maybe Li-ion)
	- SD card adapter project exists https://loopgain.net/multispeedhdd/

## Psion
- Psion Organiser II XP (1986)
	- Missing cover
	- Mysterious pak which I dumped and reversed engineered and can't make heads or tails of the purpose of, called PLT something
	- 32k rampak
	- 16k datapak
- Psion Organiser II LZ 32K (1989)
	- Missing cover
	- "CYGNUS INSTRUMENTS LTD CYGPAK SOFTWARE" pak - some kind of data logging software
	- Mysterious "SES DATA SYSTEMS" rugged case which plugs into the cartridge slot and provides a metal 7-pin connector - I believe this goes with the Cygpak and is for truckers
	- 32k rampak
- Arduino OrgII pak flasher on a breadboard - think I copied this from Menadue
- Started writing OrgII RTTTL player/creator but dropped off when I moved
	- Must disable interrupts or playback is choppy from the interruptions
	- Need to write utilities to make it easier to call assembly from OPL
	- None of the existing emulators supported sound so the dev loop is really annoying — might be able to replace pak flashing/on-device editing with a DIY link cable
- Oregon Scientific Osaris - EPOC (1999)
	- Psion 5 clone

## Windows CE/PocketPC/Mobile
- HP Jornada 680 (1999)
	- CE
	- Was working perfectly but then the screen got fucked up, I think loose/torn ribbon cable but haven't looked
	- Working wifi card which I can't remember how to remove without breaking the plastic
	- Not sure if the stylus is original, it's weird (clicks down like a pen) and falls out but it's the same color and shape
- Dell Axim X5 (2002)
	- Cradle
	- Metal case
- NEC MobilePro 900 (2003)
	- LinkSys WPC11 wifi PC Card
- HP iPAQ 210 (2007)
	- Soft black case
- LinkSys Wireless CompactFlash Card WCF12
	- I think I bought this for the Jornada or MobilePro? Slightly different model number than the one in the MobilePro and says CF not PC Card
	- In box with ridiculous custom carrying case
- Cisco Aironet 350 Wireless Adapter
	- ???

## Calculators & Organizers
### Programmable
- Casio Planeo Color Digital Diary NX-6000 (1996)
	- While this is technically programmable, I did not make much headway with it
	- Mostly just bought this for the gorgeous graphics after I saw it on Posy's YouTube
- Sharp Wizard IQ-7720M (1984)
	- OZ-707 Scientific Computer Card
	- Getting serial is a massive pain in the ass because I don't have the weird connector, just jam jumper leads in and be gentle
- TI-84 Plus Silver Edition (2004)
	- Screen has large blemish
	- TI Presentation Link
	- TI View Screen
	- Was gonna put Collapse OS on here but didn't happen for reasons I don't recall
- TI-92
	- Fucked up and got the version without the replaceable ROM, II or Plus or something
	- Banned in schools for... having a keyboard?

### Free Junk Organizers
- TI PS-2400 Data Bank
- Sharp Memo Master 100
- Sharp ZQ-2200

## Misc.
- AlphaSmart Neo2 (2007)
- SMART Response XE
	- Classroom remote responder that can be modded
- Zipit Wireless Messenger ZWM1 (2005?)
	- Not easily modified to run Linux like the Z2, I needed a serial chip which I bought and then it was too small to solder or something. Or maybe that was the SMART Response XE. Both stupid purchases

## Wish List
- Datavue Spark (laptop)
- IBM Simon (first smartphone, hella pricey)
- Motorola Envoy or Sony Magic Link (Magic Cap)
- Psions, especially Organiser I
- Sharp Expert Pad (Apple Newton)
- Weird Cliés
- Working GridPad
- RangeLan2 router (pre-wifi wireless LAN — I have a fuckton of the PC Cards for some reason; I think the Jornada ROM supports them? pops up on eBay for absurd prices)
- Someone to take all these fucking cradles
