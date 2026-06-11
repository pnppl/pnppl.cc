const html = document.getElementsByTagName('html')[0];

const fileInput = document.getElementById("file-input");
const dragDrop = document.getElementById("drag-drop");
const fileContentDisplay = document.getElementById("file-content");
const messageDisplay = document.getElementById("message");

const consoles = document.getElementById("consoles");
const pages = document.getElementById("pages");
const ignore = document.getElementById("ignore");
const styles = document.getElementById("styles");
const scripts = document.getElementById("scripts");

const update = document.getElementById("update");
const save = document.getElementById("save");

const iframe = document.getElementById('iframe');

var wander = { consoles: [], pages: [], ignore: [], styles: [], scripts: [] };

// whole page drag and drop
document.addEventListener('dragover', (e) => {
		e.preventDefault();
		setBG('drag');
});
document.addEventListener('drop', (e) => {
	// without the test it intercepts ALL drops including text
	if (e.dataTransfer.files.length > 0) {
			document.getElementById('file-input').files = e.dataTransfer.files;
			e.target.files = e.dataTransfer.files;
			e.preventDefault();
			handleFileSelection(e);
			setBG();
	}
});
document.addEventListener('dragend', (e) => {
	setBG();
});
document.addEventListener('dragleave', (e) => {
	setBG();
});

fileInput.addEventListener("change", handleFileSelection);
update.addEventListener("click", (e) => {
	readFields();
	updateFields();
});
save.addEventListener("click", (e) => {
	exportJS();
});

// add listeners to all style inputs
window.onload = function() {
	const tools = document.querySelectorAll('.style-tool');
	tools.forEach(
		function(tool) {
			tool.addEventListener("change", (e) => {updateStyle(tool)});
		}
	);
};

// https://stackoverflow.com/a/44134328
function hslToHex(h, s, l) {
	l /= 100;
	const a = s * Math.min(l, 1 - l) / 100;
	const f = n => {
		const k = (n + h / 30) % 12;
		const color = l - a * Math.max(Math.min(k - 3, 9 - k, 1), -1);
		return Math.round(255 * color).toString(16).padStart(2, '0');	 // convert to Hex and prefix "0" if needed
	};
	return `#${f(0)}${f(8)}${f(4)}`;
}

// https://stackoverflow.com/a/62390801
function hexToHSL(hex) {
	const result = /^#?([a-f\d]{2})([a-f\d]{2})([a-f\d]{2})$/i.exec(hex);

	let r = parseInt(result[1], 16);
	let g = parseInt(result[2], 16);
	let b = parseInt(result[3], 16);

	r /= 255, g /= 255, b /= 255;
	let max = Math.max(r, g, b), min = Math.min(r, g, b);
	let h, s, l = (max + min) / 2;

	if (max == min){
		h = s = 0; // achromatic
	} else {
		var d = max - min;
		s = l > 0.5 ? d / (2 - max - min) : d / (max + min);
		switch(max) {
			case r: h = (g - b) / d + (g < b ? 6 : 0); break;
			case g: h = (b - r) / d + 2; break;
			case b: h = (r - g) / d + 4; break;
		}

		h /= 6;
	}

	h = Math.round(h*360);
	s = Math.round(s*100);
	l = Math.round(l*100);

	return { h, s, l };
}

function updateStyle(tool) {
	let id = tool.id;
	let newValue = tool.value;
	let style = document.createElement('style');
	switch(id){
		case "bg-color":
			style.textContent = `body { background: ${newValue}; }`;
			break;
		case "text-color":
			style.textContent = `body, button, input { color: ${newValue}; }`;
			break;
		case "input-color":
			// automatically calculate hover/active/disabled
			const newHSL = hexToHSL(newValue);
			const h = newHSL.h;
			const s = newHSL.s;
			const l = newHSL.l;
			let hover, active, disabled;
			if (newHSL.l >= 50) {
				hover = hslToHex(h, s, l * 0.88);
				active = hslToHex(h, s, l * 0.8);
				disabled = hslToHex(h, s, l * 1.1);
			} else {
				// "dark mode"
				hover = hslToHex(h, s, l * 1.12);
				active = hslToHex(h, s, l * 1.2);
				disabled = hslToHex(h, s, l * 0.9);
			}
			style.textContent = `
				button, input, dialog, dialog details[open] { background: ${newValue}; }
				button:hover { background: ${hover}; }
				button:active { background: ${active}; }
				button:disabled, button:disabled:hover, button:disabled:active, dialog { background: ${disabled}; }
			`;
			break;
		case "border-color":
			style.textContent = `button, input, dialog, #wander-iframe { border-color: ${newValue}; }`;
			break;
		case "border-width":
			style.textContent = `button, input, dialog, #wander-iframe { border-width: ${newValue}px; }`;
			break;
		case "border-style":
			style.textContent = `button, input, dialog, #wander-iframe { border-style: ${newValue}; }`;
			break;
		case "font-family":
			style.textContent = `body, button, input { font-family: ${newValue}; }`;
			break;
//		this is black with an alpha value. the ffx color picker doesn't seem to support alpha... god it's so bad
//		change it to a range later. or just. leave it alone
//		case "backdrop":
//			style.textContent = `dialog::backdrop { background: rgba(from ${newValue} r g b / 0.6); }`;
//			break;
//		no idea how to preview this and it probably doesn't matter
//		case "noscript":
//			style.textContent = `noscript { color: ${newValue}; border-color: ${newValue}; }`;
	}
	iframe.contentDocument.head.appendChild(style);
}

// toggle background color
function setBG(color) {
	if (color === 'drag') {
		html.style.backgroundColor = 'lightskyblue';
	} else {
		html.style.backgroundColor = '';
	}
}

// file selection results
function showMessage(message, type) {
	messageDisplay.textContent = message;
	messageDisplay.style.color = type === "error" ? "red" : "green";
}

function handleFileSelection(event) {
	const file = event.target.files[0];
	fileContentDisplay.textContent = ""; // Clear previous file content
	messageDisplay.textContent = ""; // Clear previous messages

	// Validate file existence and type
	if (!file) {
		showMessage("No file selected. Please choose a file.", "error");
		return;
	}
	// should we really bother with this? what if the mime is fucked and it looks like plaintext?
	if (!file.type.match("application/x-javascript")) {
		showMessage("Unsupported file type. Please select your wander.js file.", "error");
		return;
	}

	// Read the file
	const reader = new FileReader();
	reader.onload = () => {
		fileContentDisplay.textContent = reader.result;
		try {
			eval(`${reader.result.replace("const ", "")}`);
			updateFields();
		} catch (err) {
			showMessage("Error parsing wander.js. Probably a missing comma or someting.", "error");
		}
//		showMessage("Loaded existing wander.js successfully.");
	};
	reader.onerror = () => {
		showMessage("Error reading the file. Please try again.", "error");
	};
	reader.readAsText(file);
}

function asCode() {
	return "const wander = " + JSON.stringify(wander, null, "\t");
}

function exportJS() {
	var blob = new Blob([ asCode() ], { type: "text/plain;charset=utf8" });
	var a = document.createElement('a');
	a.download = "wander.js";
	a.href = window.URL.createObjectURL(blob);
	a.click();
	if (a.remove) a.remove();
}

// replace page contents with contents of global wander var
function updateFields() {
	consoles.value = wander.consoles.join("\n");
	pages.value = wander.pages.join("\n");
	ignore.value = wander.ignore.join("\n");
	styles.value = wander.styles.join("\n");
	scripts.value = wander.scripts.join("\n");
	fileContentDisplay.textContent = asCode();
}

// update global wander var with <textarea> user input
function readFields() {
	// filter removes blank lines (empty string is "falsy" in this accursed language)
	wander.consoles = consoles.value.split("\n").filter(Boolean);
	wander.pages = pages.value.split("\n").filter(Boolean);
	wander.ignore = ignore.value.split("\n").filter(Boolean);
	wander.styles = styles.value.split("\n").filter(Boolean);
	wander.scripts = scripts.value.split("\n").filter(Boolean);
}


