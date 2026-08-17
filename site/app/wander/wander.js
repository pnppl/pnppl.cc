const wander = {
	// Prefer consoles with: links to posts, not site root; non-techy links; few neighbors; up-to-date Wander version
	consoles: [
		'https://antonio.is/wander/',
		'https://exurd.neocities.org/wander/',
		'https://heckmeck.de/wander/',
	],
	// More info and alternate browsing mode at https://pnppl.cc/app/meander
	pages: [
		'https://pnppl.cc/app/wcb/',

		// SSL
		// Nonfiction
		'https://100r.co/site/tools_ecosystem.html',
		'https://27bslash6.com/overdue.html',
		'https://aresluna.org/the-hardest-working-font-in-manhattan/',
		'https://aria.dog/barks/every-platform-is-hitler/',
		'https://blog.polly.computer/untuck_NOW_queen/',
		'https://c4ss.org/content/61050',
		'https://collapseos.org/why.html',
		'https://daringfireball.net/2026/05/what_is_a_dickover',
		'https://dmitry.gr/?r=05.Projects&proj=27.%20rePalm',
		'https://drewdevault.com/2025/09/24/2025-09-24-Cloudflare-and-fascists.html',
		'https://ellie.clifford.lol/blog/0023-the-sixth-of-may/',
		'https://freethoughtblogs.com/nataliereed/2012/04/17/the-null-hypothecis/',
		'https://genderanalysis.net/2017/06/depersonalization-in-gender-dysphoria-widespread-and-widely-unrecognized/',
		'https://gillianbrockell.com/tracking-all-of-trumps-third-country-removals-that-we-know-of/',
		'https://haters.noblogs.org/zines/',
		'https://j3s.sh/thought/blogs-rot-wikis-wait.html',
		'https://kevinboone.me/web-adjacent.html',
		'https://maia.crimew.gay/posts/the-emails/',
		'https://moth.monster/blog/artificial-iconoclasm/',
		'https://ploum.net/2026-01-05-unteaching_github.html',
		'https://prairielanddefendants.com/resources/zines-flyers/',
		'https://sexabolition.blog/fuck-biological-sex-we-have/',
		'https://site.sebasmonia.com/posts/2026-04-09-the-redirection-of-traffic-from-web-to-ai-doesn-t-happen-in-a-vacuum.html',
		'https://tagonist.livejournal.com/199563.html',
		'https://theanarchistlibrary.org/library/peter-gelderloos-how-nonviolence-protects-the-state',
		'https://thomaswebb.net/2024/11/16/the-normal-people-understanders/',
		'https://web.archive.org/web/20230524201754/https://destroyedforcomfort.com/2018/09/04/4-weird-things-that-happen-to-you-when-a-loved-one-kills-themselves/',
		'https://www.filfre.net/2011/05/will-crowthers-adventure-part-1/',
		'https://www.fsf.org/blogs/licensing/2026-anthropic-settlement',
		'https://www.thedissident.news/every-trans-suicide-is-a-murder/',

		// Fiction
		'https://www.galactanet.com/oneoff/theegg_mod.html',

		// Comic/still image
		'https://anhvn.com/noir/',
		'https://hillhouse.neocities.org/cliques/library/',
		'https://kalebhorton.ghost.io/the-d-lbert-project/',
		'https://killsixbilliondemons.com/comic/kill-six-billion-demons-chapter-1/',
		'https://www.qwantz.com/',

		// Audio/video
		'https://archive.org/details/MindWebs_201410',
		'https://kevincraig.us/audio/gogulski/',
		'https://thefinalstrawradio.noblogs.org/post/2026/04/12/the-life-and-ideas-of-johann-most-with-tom-goyens/',
		'https://www.churchofeuthanasia.org/catalog/video.html',
		'https://www.fuckyoutube.lol/youtube_abyss',
		'https://www.juliaserano.com/music.html',
		'https://www.rifters.com/blindsight/vampires.htm',

		// Game/interactive
		'https://freegames.org/dys4ia/',
		'https://xrafstar.monster/games/twine/tails/',

		// Misc. resource
		'https://fishshell.com/',
		'https://frills.dev/blog/231207-edit-everything/',
		'https://gridbeam.xyz/guide',
		'https://littlebitspace.com/resources/',
		'https://marginalia-search.com/',
		'https://rubjo.github.io/victor-mono/',

		// Site root
		'https://adelfaure.net/',
		'https://fourthievesvinegar.org/',
		'https://fuckup.solutions/index3.html',
		'https://june11.noblogs.org/',
		'https://kvibber.com/',
		'https://solar.lowtechmagazine.com',
		'https://wiki.archiveteam.org/',
		'https://www.gothic-charm-school.com/',
		'https://www.halfbakery.com/',
		'https://www.vhemt.org/',
		'https://www.wendycarlos.com/',

		// News
		'https://unicornriot.ninja/',
		'https://www.techdirt.com/',

		// HTTP
		// Nonfiction
		'https://friendo.monster/posts/emojis-are-shit.html',
		'https://houseofselfindulgence.blogspot.com/2008/08/dr-caligari-stephen-sayadian-1989.html',
		'https://humaniterations.net/2020/09/06/bad-people/',
		'https://juliaserano.blogspot.com/2019/02/origins-of-social-contagion-and-rapid.html',
		'https://maurycyz.com/projects/bad_jpeg/',
		'https://mycophobia.org/dcs/index.html',
		'https://notes.highlysuspect.agency/are-we-doing-anything.html',
		'https://trekkiefeminist.com/star-trek-bechdel-wallace-test-results-graphed/',

		// Fiction
		'https://www.terrybisson.com/theyre-made-out-of-meat-2/',

		// Comic
		'https://baccyflap.com/mus/ti/',

		// Audio/video
		'https://echoesofbluemars.org/',
		'https://handthing.party/',
		'https://smokepowered.com/',
		'https://www.quoteunquoterecords.com/qur022.htm',

		// Misc. resource
		'https://temblast.com/android.htm',
		'https://vivivi.leprd.space/webmastery/pre-css-attributes/',
		'https://web1.0hosting.net/',
		'https://wiby.me',

		// Site root
		'https://petermolnar.net/',
		'https://velveteen.one/',
		'https://www.floodgap.com/',
		'https://www.prole.info/',
		'https://xn--gckvb8fzb.com/',
	],
	ignore: [
		// Rationalists, racists, fascists, transphobes
		'https://*.astralcodexten.com/',
		'https://dhh.dk/',
		'https://gwern.net/',
		'https://*.lesswrong.com/',
		'https://slatestarcodex.com/',
		'https://x.com/',
		'https://xahlee.info/',
		'https://xahlee.org/',
		'https://*.yudkowsky.net/',

		// Garbage silos
		'https://medium.com/',
		'https://*.substack.com/',

		// Blocked by HTTP headers, etc.
		// the best site in the universe
		'https://annas-archive.gl/',
		'https://codeberg.org/',
		'https://danielmiessler.com/',
		// i love https://dbushell.com/2026/01/09/death-to-scroll-fade/
		'https://dbushell.com/',
		'https://*.geek.nz/',
		'https://hyperdoc.khinsen.net/',
		'https://*.johndcook.com/',
		'https://lowendmac.com',
		'https://maggieappleton.com',
		// fantastic games
		'https://neal.fun/',
		'https://*.otherstrangeness.com',
		'https://thesweetbits.com/',

		// YC, HN, AI, capitalists
		'https://foundersatwork.posthaven.com/',
		'https://paulgraham.com/',
		'https://*.samaltman.com/',
		'https://simonwillison.net/',
		'https://stratechery.com/',
		'https://*.ycombinator.com',

		// Non-English language
		'https://*.64er-magazin.de/',
		'https://sannalund.se/',

		// WebGL
		'https://eightyeightthirty.one/',

		// Misc.
		// filter pointless "intentionally left blank" pages
		'https://*/blank/',
		// i'm sure she's great but that banner drives me up the fucking wall
		'https://sachachua.com',
		// paywalled; zealot
		'https://*.wheresyoured.at/',

		// Consoles
		// capitalist/HN shit
		'https://*.davidtran.me/wander/',
		'https://*.heyhomepage.com/wander/',
		// sorry Joshes but my name is not Josh
		'https://joshing.you/wander/',
	],
	styles: [
		// win9x style; display 'Open' on mobile
		'wander.css',
	],
	scripts: [
		// move 'About' into 'Console' menu, add button advertising Wander Console Builder
		'modify-menu.js',
	]
}
