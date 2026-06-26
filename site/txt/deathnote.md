# Digital Death Note
#project
<style><!--
/* --- death note cover look --- */
html {
	background-color: black;
}
#article {
	background-color: white;
}
@media(prefers-color-scheme:dark) {
	#article {
		background-color: transparent;
	}
}
.established {
	color: #CBCFD8;
}
/* it's ok the coolness allows it */
@font-face {
	font-family: dntitle;
	font-display: swap;
	src: local("no"), url("/public/font/death-note.woff2") format("woff2");
}
.title h1,
.column time {
	font-family: dntitle;
	font-style: normal;
	font-size: 2em;
	color: #ebecf0;
	text-transform: lowercase;
}
.title {
	font-size: 3em;
	margin: 0.5em 0;
	line-height: 0.8;
}
.subtitle,
.columns,
.tag,
h2,
aside {
	display: none;
}

/* --- notebook paper look --- */
#article {
	background-image:
		linear-gradient(90deg, transparent calc(2em - 1px), pink 2em, transparent 2em),
		repeating-linear-gradient(
			to bottom,
			transparent 0px,
			transparent calc(2em - 1px),
			#ccc calc(2em - 1px),
			#ccc 2em
		);
	padding: 2.15em 1em calc(1.85em - 1px) 3em;
	border: 1px solid gray;
	box-shadow: 0 0 5px gray;
}
#article * {
	margin: 0;
	padding: 0;
	font-family: "Victor Mono Medium", cursive, serif;
	font-style: italic;
	line-height: 2;
	font-size: 1em;
}

/* --- guestbook styling --- */
#gbformmessage,
.guest-item:has(a),
#guestbook hr {
	display: none;
}
#article #guest-messages {
	padding-bottom: 2em;
	display: flex;
	flex-direction: column-reverse;
}
#article div:has(>input),
#article div:has(>textarea) {
	margin-top: calc(-0.15em - 1px);
	padding-bottom: calc(0.15em + 1px);
}
textarea,
input[name="name"] {
	width: 100%;
}
#article textarea,
#article input {
	height: 2em;
	border-radius: 0;
	border-width: 1px;
	padding: 0 0.5em;
}
#article input[name=math] {
	margin-top: 1px;
}
#article input[type=submit] {
	font-family: dntitle;
	text-transform: uppercase;
	font-weight: bold;
	letter-spacing: 2px;
	display: block;
	margin: 2em auto 0;
}
.guest-timestamp {
	display: none;
}
.guest-name:after {
	content: ' — ';
	font-weight: normal;
}
.guest-name:has(~ .guest-message:empty):after {
	content: '';
}
form:before {
	content: "Add a name... if you dare...";
	display: block;
	font-weight: bold;
	text-align: center;
	font-family: dntitle;
	font-size: 2em;
	padding: 0.5em 0;
	text-transform: lowercase;
}
#gbformname:after {
	content: ":";
}
div:has(>textarea):before {
	content: "Cause of death:";
	display: block;
}
#gbformurl:after {
	content: " (posts with URL will not be shown):";
}
.message-error:after {
	content: " (message = cause of death)";
}
/* captcha */
#article div:has(>input[name=math]) {
	text-align: center;
	padding-top: 1.55em;
}
div:has(>input[name=math]):before {
	content: "Pass the shinigami's test: "
}
/* abuse filter - prepend ! to an entry's timestamp to hide */
#abuse {
	height: 2.5ch;
	margin-right: 1ch;
}
#abuse + label {
	font-style: normal;
}
.guest-item[id^='guest-!'] {
	display: none;
}
#abuse:checked ~ #guestbook .guest-item[id^='guest-!'] {
	display: inline;
}
--></style>
<!-- ~~~~~~ THE ACTUAL POST CONTENT GOES BELOW  ~~~~~~ -->
**Donald John Trump** — messy post-McDonalds ibogaine overdose
**Alexander Caedmon Karp** — shot by racist stalker
**Peter Andreas Thiel** — executed for homosexuality
**Benjamin Netanyahu** — Israeli missile targeting accident
**Elon Reeve Musk** — bludgeoned with sink by robot
**Joanne Rowling** — public bathroom sewage flood
<!-- ~~~~~~ THE ACTUAL POST CONTENT GOES ABOVE ~~~~~~ -->
<input id="abuse" name="abuse" type="checkbox"><label for="abuse">Show entries pnppl has marked as offensive</label>
<!--#include virtual="/cgi-bin/gb/?$args" -->
<script><!--// the dark side...
if (document.querySelector) {
	// clear out and hide URL field (pre-populated with protocol that prevents submission)
	document.querySelector('input[type=url]').value = '';
	document.querySelector('div:has(>input[type=url])').style.display = "none";
	document.querySelector('#gbformurl').style.display = "none";
	// fix timestamp formatting
	const timestamps = document.getElementsByClassName('guest-timestamp');
	for (let stamp of timestamps) {
		stamp.textContent = stamp.textContent.replaceAll('/', '-');
	}
	// label button
	document.querySelector('input[type=submit]').value = "Write";
}
//--></script>
