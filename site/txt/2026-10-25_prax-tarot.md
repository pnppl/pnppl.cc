# Émile Reynaud's Praxinoscope Tarot
#project #visual-art #tarot

## [[Émile Reynaud's Praxinoscope Tarot|http://pnppl.cc/img/tarot/emile-reynaud-praxinoscope/!erp-tarot.zip]] (3.5M)

!p180x250 ../img/tarot/emile-reynaud-praxinoscope/

## About

I came across [[this collection of early animations by Émile Reynaud|https://magot.neocities.org/grid/]] and quickly noticed two things:

1. They're very pretty and whimsical and resemble fortune-telling cards.
2. There are 23 of them... and you only need 22 for the tarot trumps ("Major Arcana").

So, yeah. I assigned one animation to each card and added some text and stuff. I also added a card backing taken from the same source.

Anyway, go check out [[Margot McEwen's awesome website|https://magot.neocities.org/]].

## Draw (js)
<div id="random" title="draw a card"><i id="del">This part of the page lets you draw random cards. It requires JavaScript, which your browser does not support or has disabled.</i></div>
<script><!--
cards = ['00_fool.gif',
	'01_magician.gif',
	'02_high-priestess.gif',
	'03_empress.gif',
	'04_emperor.gif',
	'05_hierophant.gif',
	'06_lovers.gif',
	'07_chariot.gif',
	'08_strength.gif',
	'09_hermit.gif',
	'10_wheel-of-fortune.gif',
	'11_justice.gif',
	'12_hanged-man.gif',
	'13_death.gif',
	'14_temperance.gif',
	'15_devil.gif',
	'16_tower.gif',
	'17_star.gif',
	'18_moon.gif',
	'19_sun.gif',
	'20_judgement.gif',
	'21_world.gif'];
function prep() {
	i = cards.length;
	deck = Array(i);
	while(i--) deck[i] = cards[i];
	deck.sort(() => Math.random() - 0.5);
	drawn = document.getElementsByClassName('drawn');
	while (drawn.length > 0) {
		drawn[0].remove();
	}
}
function draw() {
	if (deck.length == 0) {
		prep();
		container.title = 'draw a card';
		document.getElementById('last').remove();
	}
	i = Math.floor(Math.random()*deck.length);
	card_src = deck[i];
	deck.splice(i, 1);
	card = document.createElement('img');
	card.src = '../img/tarot/emile-reynaud-praxinoscope/' + card_src;
	card.width = 180;
	card.height = 250;
	card.className = 'tarot drawn';
	container.innerHTML += '\n\n';
	container.appendChild(card);
	if (deck.length == 0) {
		container.innerHTML += ' <i id="last">all cards have been drawn</i>';
		container.title = 'shuffle cards';
	}
}
prep();
document.getElementById('del').remove();
container = document.getElementById('random');
back = document.createElement('img');
back.src = '../img/tarot/emile-reynaud-praxinoscope/' + '00_back.gif';
back.width = 180;
back.height = 250;
back.className = 'tarot back';
container.appendChild(back);
container.addEventListener('click', draw, false);
//--></script>
