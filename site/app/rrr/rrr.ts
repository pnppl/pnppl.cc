const debounceEndSeconds: number = 3;
const debounceTouchMs: number = 200;
const dt = new Date();

let progress: number | null = null;
let breaths: number = 0;
let breathTimer = setInterval(null, null);
let debounceEnd: boolean = false;
let debounceTouch: boolean = false;


function getIcon() {
	let bpm: number = breaths * 2;
	if (bpm > 30 || bpm < 15) {
		return "💔";
	}
	return "❤️";
}
function getBpmString() {
	let bpm: number = breaths * 2;
	return bpm.toString().padStart(2, "0");
}

function getResultString() {
	let month: string = dt.getMonth().toString().padStart(2, "0");
	let day: string = dt.getDate().toString().padStart(2, "0");
	let hour: string = dt.getHours().toString().padStart(2, "0");
	let min: string = dt.getMinutes().toString().padStart(2, "0");
	return `${getIcon()} | ${getBpmString()} breaths per minute | ${dt.getFullYear()}-${month}-${day} | ${hour}:${min}`;
}

function end() {
	clearInterval(breathTimer);
	progress = null;
	document.getElementById('result-string').innerHTML = `${getIcon()} ${breaths * 2} breaths per minute`;
	document.getElementById('results').style.display = 'initial';
	setTimeout(() => {
		debounceEnd = false;
	}, debounceEndSeconds * 1000)
}

function tick() {
	progress -= 1;
	document.getElementById('seconds').innerHTML = progress.toString(10);
	let progressInverse: number = Math.abs(progress - 30)
	document.getElementById('prog').setAttribute('value', progressInverse.toString());
	if (progress === 0) {
		end();
	}
}

function start() {
	progress = 30;
	breaths = 0;
	clearInterval(breathTimer);
	debounceEnd = true;
	dt.setTime(Date.now());
	document.getElementById('seconds').innerHTML = '30';
	document.getElementById('prog').setAttribute('value', '0');
	document.getElementById('guide').style.display = 'none';
	document.getElementById('results').style.display = 'none';
	document.getElementById('check').style.display = 'none';
	breathTimer = setInterval(tick, 1000);
}

function breathe() {
	if (progress != null) {
		breaths += 1;
		document.getElementById('breaths').innerHTML = breaths.toString();
	}
}

function press() {
	if (progress === null && !debounceEnd) {
		start();
	}
	breathe();
}

function setup() {
	document.getElementById('breath').addEventListener('pointerup', () => {
		if (!debounceTouch) {
			press();
			debounceTouch = true;
			setTimeout(() => {
				debounceTouch = false;
			}, debounceTouchMs)
		}
	});
	document.getElementById('breath').addEventListener('pointerdown', () => {
		if (progress === null && !debounceEnd) {
			start();
		}
	});
	document.getElementById('copy').addEventListener('click', () => {
		navigator.clipboard.writeText(getResultString());
		document.getElementById('check').style.display = 'initial';
	});
}


window.onload = function() {
	setup();
};
