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

// add listeners to all color inputs
window.onload = function() {
	const colors = document.querySelectorAll('input[type=color]');
	colors.forEach(
		function(color) {
			color.addEventListener("change", (e) => {updateColor(color)});
		}
	);
};

function updateColor(color) {
	let id = color.id;
	let newColor = color.value;
	let style = document.createElement('style');
	switch(id){
		case "body-bg":
			style.textContent = `body { background: ${newColor}; }`;
			break;
		case "main-fg":
			style.textContent = `body, button, input { color: ${newColor}; }`;
			break;
		case "input-bg":
			// automatically calculate hover/active
			// need to handle "dark mode" later (getting lighter instead of darker)
			style.textContent = `
				button, input, dialog, dialog details[open] { background: ${newColor}; }
				button:hover { background: hsl(from ${newColor} h s calc(l * 0.88)); }
				button:active { background: hsl(from ${newColor} h s calc(l * 0.8)); }
			`;
			break;
		case "borders":
			style.textContent = `button, input, dialog, #wander-iframe { border-color: ${newColor}; }`;
			break;
		case "grey-bg":
			style.textContent = `button:disabled, button:disabled:hover, button:disabled:active, dialog { background: ${newColor}; }`;
			break;
//		this is black with an alpha value. the ffx color picker doesn't seem to support alpha... god it's so bad. change it to a range later
//		case "backdrop":
//			style.textContent = `dialog::backdrop { background: rgba(from ${newColor} r g b / 0.6); }`;
//			break;
		case "noscript":
			style.textContent = `noscript { color: ${newColor}; border-color: ${newColor}; }`;
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


