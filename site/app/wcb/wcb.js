// instead of an init()
// yes the entire script is wrapped in this. yes it's deranged. yes it works on my machine
window.onload = function() {
	const styleName = "wander-wcb.css";

	const fileInput = document.getElementById("file-input");
	const urlInput = document.getElementById("url-input");
	const messageDisplay = document.getElementById("message");
	const fileContentDisplay = document.getElementById("file-content");

	const consoles = document.getElementById("consoles");
	const pages = document.getElementById("pages");
	const ignore = document.getElementById("ignore");
	const styles = document.getElementById("styles");
	const scripts = document.getElementById("scripts");

	const tools = document.querySelectorAll('.style-tool');
	const iframe = document.getElementById('iframe');
	const applyTheme = document.getElementById("apply-theme");
	const reset = document.getElementById("reset");

	let wander = { consoles: [], pages: [], ignore: [], styles: [], scripts: [] };
	let css = Array(7);

	// --- register all listeners ---//
	// whole page drag and drop file input
	document.addEventListener('dragover', (e) => {
			e.preventDefault();
			setBG('drag');
	});
	document.addEventListener('drop', (e) => {
		// without the test it intercepts ALL drops including text
		if (e.dataTransfer.files.length > 0) {
				fileInput.files = e.dataTransfer.files;
				e.target.files = e.dataTransfer.files;
				e.preventDefault();
				handleFileSelection(e);
				setBG();
		}
	});
	// the end
	document.addEventListener('dragend', setBG);
	document.addEventListener('dragleave', setBG);

	// actual file input element
	fileInput.addEventListener("change", handleFileSelection);

	// url input
	const urlForm = document.getElementById("url");
	urlForm.addEventListener("submit", loadRemote);

	// listener for iframe hack for url input
	window.addEventListener('message', handleRemote);

	// update button
	const update = document.getElementById("update");
	update.addEventListener("click", () => {
		readFields();
		updateFields();
	});

	// save JS button
	const save = document.getElementById("save");
	save.addEventListener("click", () => {
		exportBlob("js");
	});

	// reset button
	reset.addEventListener("click", resetFields);

	// add listeners to all style controls
	tools.forEach(
		function(tool) {
			tool.addEventListener("change", () => {updateStyle(tool)});
		}
	);

	// apply theme checkbox
	applyTheme.addEventListener("change", function() {
		applyThemeFn(this.checked);
	});

	// save theme button
	const saveCSS = document.getElementById("save-css");
	saveCSS.addEventListener("click", () => {
		exportBlob("css");
	});

	// check for url fragment to load immediately
	loadParam();

	// --- end init

	// --- helper functions --- //

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

	function asCode() {
		return "const wander = " + JSON.stringify(wander, null, "\t");
	}

	// toggle background color for drag and drop
	function setBG(color) {
		const html = document.getElementsByTagName('html')[0];
		if (color === 'drag') {
			html.style.backgroundColor = 'lightskyblue';
		} else {
			html.style.backgroundColor = '';
		}
	}

	// display file/url import results
	function showMessage(message, type) {
		messageDisplay.textContent = message;
		messageDisplay.style.color = type === "error" ? "firebrick" : "green";
	}


	// --- meaty functions --- //

	// file input
	// mostly copied from MDN
	function handleFileSelection(event) {
		const file = event.target.files[0];
		fileContentDisplay.textContent = ""; // Clear previous file content
		messageDisplay.textContent = ""; // Clear previous messages

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
				let wanderFile;
				eval(`${reader.result.replace("const wander", "wanderFile")}`);
				importConsole(wanderFile);
				updateFields();
				urlInput.value = "";
				showMessage("Loaded console from file.");
			} catch (err) {
				showMessage(`Error parsing wander.js. Probably a missing comma or someting. (${err})`, "error");
			}
		};
		reader.onerror = () => {
			showMessage("Error reading the file. Please try again.", "error");
		};
		reader.readAsText(file);
	}

	function cleanURL(url, mode) {
		if ( url.startsWith("//") ) {
			url = "https:" + url;
		}
		switch (mode) {
			case "remote":
				if ( url.startsWith(".") || url.startsWith("wander.js") ) {
					return url;
				}
				break;
			case "console":
				if ( url.endsWith("wander.js") ) {
					url = url.slice(0, -9);
				}
				break;
		}
		if (! url.startsWith("http") ) {
			url = "https://" + url;
		}
		return url;
	}

	// url input
	function loadRemote(e) {
		const remoteURL = cleanURL(e.srcElement[0].value, "remote");
		const loader = document.getElementById('loader-iframe')
		// stolen from susam
		loader.srcdoc = `
			<script src="${remoteURL}"></scr` + `ipt>
			<script>
				try {
					parent.postMessage({ wander: wander }, '*');
				} catch (err) {
					parent.postMessage({ err: err }, '*');
				}
			</scr` + 'ipt>'
	}
	// iframe hack for url input
	function handleRemote(e) {
		showMessage("");
		if (typeof e.data.err !== "undefined") {
			// the error will always be "wander is not defined" afaik, so no point showing it to the user
			showMessage("Couldn't fetch console from URL.", "error");
		} else {
			const ext = e.data.wander;
			importConsole(ext);
			updateFields();
			fileInput.value = "";
			showMessage("Loaded console from URL.");
		}
	}

	// load consle from url parameter, ie ?http://myurl.com
	function loadParam() {
		const param = window.location.search;
		if (param[0] === '?' && param.length > 1) {
			loadRemote({ srcElement: [{ value: param.slice(1) }] });
		}
		// for "dist" version -- on my version i use dem0.html
		else if (iframe.src.endsWith("index.html")) {
			loadRemote({ srcElement: [{ value: "wander.js" }] });
		}
	}

	// safe(r) imports from both url and file
	function importConsole(c) {
		if (typeof c === "undefined") {
			wander = { consoles: [], pages: [], ignore: [], styles: [], scripts: [] };
			return;
		}

		if (typeof c.consoles === "undefined") { wander.consoles = []; }
		else { wander.consoles = c.consoles; }

		if (typeof c.pages === "undefined") { wander.pages = []; }
		else { wander.pages = c.pages; }

		if (typeof c.ignore === "undefined") { wander.ignore = []; }
		else { wander.ignore = c.ignore; }

		if (typeof c.styles === "undefined") { wander.styles = []; }
		else { wander.styles = c.styles; }

		if (typeof c.scripts === "undefined") { wander.scripts = []; }
		else { wander.scripts = c.scripts; }
	}

	function exportBlob(filetype) {
		let blobtext, filename;
		if (filetype === "js") {
			blobtext = asCode();
			filename = "wander.js";
		} else {
			blobtext = css.join("\n").replaceAll("\t", "").replaceAll(/\n\n+/g, "\n").replace(/^\n/, "");
			filename = styleName;
		}
		let blob = new Blob([ blobtext ], { type: "text/plain;charset=utf8" });
		let a = document.createElement('a');
		a.download = filename;
		a.href = window.URL.createObjectURL(blob);
		a.click();
		if (a.remove) a.remove();
	}

	// update global wander var with <textarea> user input
	function readFields() {
		// filter removes blank lines (empty string is "falsy" in this accursed language)
		wander.consoles = consoles.value.split("\n").filter(Boolean);
		wander.consoles.forEach(
			function(c, i, arr) {
				arr[i] = cleanURL(c, "console");
			}
		);
		wander.pages = pages.value.split("\n").filter(Boolean);
		wander.pages.forEach(
			function(c, i, arr) {
				arr[i] = cleanURL(c, "");
			}
		);
		wander.ignore = ignore.value.split("\n").filter(Boolean);
		wander.ignore.forEach(
			function(c, i, arr) {
				arr[i] = cleanURL(c, "console");
			}
		);
		wander.styles = styles.value.split("\n").filter(Boolean);
		wander.scripts = scripts.value.split("\n").filter(Boolean);
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

	function resetFields() {
		if (reset.value === "Reset") {
			reset.value = "Clear everything?";
		} else if (reset.value === "Clear everything?") {
			reset.value = "Last chance. Are you sure?";
		} else {
			fileInput.value = "";
			urlInput.value = "";
			showMessage("");
			wander = { consoles: [], pages: [], ignore: [], styles: [], scripts: [] };
			updateFields();
			css = Array(7);
			document.getElementById("bg-color").value = "#696";
			document.getElementById("text-color").value = "#030";
			document.getElementById("input-color").value = "#bdb";
			document.getElementById("border-color").value = "#363";
			document.getElementById("border-width").value = "2";
			document.getElementById("border-style").value = "solid";
			document.getElementById("font-family").value = "courier, monospace";
			tools.forEach(
				function(tool) {
					updateStyle(tool);
				}
			);
			applyTheme.checked = false;
			reset.value = "Reset";
		}
	}

	// theme builder
	function updateStyle(tool) {
		const id = tool.id;
		const newValue = tool.value;
		let style = document.createElement('style');
		let styleStr;
		switch (id) {
			case "bg-color":
				styleStr = `body { background: ${newValue}; }`;
				css[0] = styleStr;
				style.textContent = styleStr;
				break;
			case "text-color":
				styleStr = `body, button, dialog, input { color: ${newValue}; }`;
				css[1] = styleStr;
				style.textContent = styleStr;
				break;
			case "input-color":
				// automatically calculate hover/active/disabled
				const newHSL = hexToHSL(newValue);
				const h = newHSL.h;
				const s = newHSL.s;
				const l = newHSL.l;
				let hover, active, disabled;
				if (newHSL.l >= 50) {
					// these attempt to mimic default style
					hover = hslToHex(h, s, l * 0.88);
					active = hslToHex(h, s, l * 0.8);
					disabled = hslToHex(h, s, l * 1.1);
				} else {
					// "dark mode"
	//				hover = hslToHex(h, s, l * 1.12);
	//				active = hslToHex(h, s, l * 1.2);
	//				disabled = hslToHex(h, s, l * 0.9);
					// adjusted to be more pronounced
					hover = hslToHex(h, s, l * 1.15);
					active = hslToHex(h, s, l * 1.25);
					disabled = hslToHex(h, s, l * 0.75);
				}
				styleStr = `
					button, input, dialog, dialog details[open] { background: ${newValue}; }
					button:hover { background: ${hover}; }
					button:active { background: ${active}; }
					button:disabled, button:disabled:hover, button:disabled:active, dialog { background: ${disabled}; }
				`;
				css[2] = styleStr;
				style.textContent = styleStr;
				break;
			case "border-color":
				styleStr = `button, input, dialog, #wander-iframe { border-color: ${newValue}; }`;
				css[3] = styleStr;
				style.textContent = styleStr;
				break;
			case "border-width":
				const bigBorder = parseInt(newValue) + 3;
				let smallBorder = newValue;
				if (newValue < 0) {
					smallBorder = 0;
				}
				styleStr = `
					button, input { border-width: ${smallBorder}px; }
					dialog, #wander-iframe { border-width: ${bigBorder}px; }
				`;
				css[4] = styleStr;
				style.textContent = styleStr;
				break;
			case "border-style":
				if (newValue === "outset") {
					styleStr = `
						button, dialog { border-style: outset; }
						input, #wander-iframe { border-style: inset; }
					`;
				} else if (newValue === "inset") {
					styleStr = `
						button, dialog { border-style: inset; }
						input, #wander-iframe { border-style: outset; }
					`;
				} else {
					styleStr = `button, input, dialog, #wander-iframe { border-style: ${newValue}; }`;
				}
				css[5] = styleStr;
				style.textContent = styleStr;
				break;
			case "font-family":
				styleStr = `body, button, input { font-family: ${newValue}; }`;
				css[6] = styleStr;
				style.textContent = styleStr;
				break;
		}
		iframe.contentDocument.head.appendChild(style);
	}

	// include/exclude custom CSS
	function applyThemeFn(checked) {
		let stylesheets = styles.value.split("\n");
		const index = stylesheets.indexOf(styleName);
		// missing; add
		if (checked && index === -1) {
			stylesheets.push(styleName);
		}
		// present; remove
		else if (!checked && index !== -1) {
			stylesheets.splice(index, 1);
		}
		styles.value = stylesheets.join("\n");
		readFields();
		updateFields();
	}
};
