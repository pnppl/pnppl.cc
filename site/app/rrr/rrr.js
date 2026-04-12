const debounceEndSeconds = 3;
const debounceTouchMs = 200;
const dt = new Date();
let progress = null;
let breaths = 0;
let breathTimer = setInterval(null, null);
let debounceEnd = false;
let debounceTouch = false;
function getIcon() {
    let bpm = breaths * 2;
    if (bpm > 30 || bpm < 15) {
        return "💔";
    }
    return "❤️";
}
function getBpmString() {
    let bpm = breaths * 2;
    return bpm.toString().padStart(2, "0");
}
// results with date and time, formatted with pipe separators for markdown table
function getResultString() {
    // the month is off by one for reasons that are entirely obscure to me
    let month = (dt.getMonth() + 1).toString().padStart(2, "0");
    let day = dt.getDate().toString().padStart(2, "0");
    let hour = dt.getHours().toString().padStart(2, "0");
    let min = dt.getMinutes().toString().padStart(2, "0");
    return `${getIcon()} | ${getBpmString()} breaths per minute | ${dt.getFullYear()}-${month}-${day} | ${hour}:${min}`;
}
function end() {
    clearInterval(breathTimer);
    progress = null;
    document.getElementById('hellip').style.display = 'none';
    document.getElementById('result-string').innerHTML = `${getIcon()} ${getBpmString()} breaths per minute`;
    document.getElementById('results').style.display = 'initial';
    setTimeout(() => {
        debounceEnd = false;
    }, debounceEndSeconds * 1000);
}
function tick() {
    progress -= 1;
    document.getElementById('seconds').innerHTML = progress.toString(10);
    let progressInverse = Math.abs(progress - 30).toString();
    document.getElementById('prog').setAttribute('value', progressInverse);
    if (progress === 0) {
        end();
    }
}
function start() {
    if (progress === null && !debounceEnd) {
        progress = 30;
        breaths = 0;
        clearInterval(breathTimer);
        debounceEnd = true;
        dt.setTime(Date.now());
        document.getElementById('seconds').innerHTML = '30';
        document.getElementById('hellip').style.display = 'initial';
        document.getElementById('prog').setAttribute('value', '0');
        document.getElementById('guide').style.display = 'none';
        document.getElementById('results').style.display = 'none';
        document.getElementById('check').style.display = 'none';
        breathTimer = setInterval(tick, 1000);
    }
}
function breathe() {
    if (!debounceTouch && progress != null) {
        breaths += 1;
        document.getElementById('breaths').innerHTML = breaths.toString();
        debounceTouch = true;
        setTimeout(() => {
            debounceTouch = false;
        }, debounceTouchMs);
    }
}
function setup() {
    document.getElementById('breath').addEventListener('pointerup', () => {
        breathe();
    });
    document.getElementById('breath').addEventListener('pointerdown', () => {
        start();
    });
    document.getElementById('copy').addEventListener('click', () => {
        navigator.clipboard.writeText(getResultString());
        document.getElementById('check').style.display = 'initial';
    });
}
// do the thing
window.onload = function () {
    setup();
};
