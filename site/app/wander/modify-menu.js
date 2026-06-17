// move the About button into the Console menu
const about = document.getElementById("about-button")
const menu = document.querySelector("#menu-section nav");
const base = document.getElementById("base-button");
menu.insertBefore(about, base)

// add button advertising Wander Console Builder
//const lastItem = document.getElementById("crawl-button");
const wcb = document.createElement("button");
wcb.textContent = "Wander Console Builder";
wcb.addEventListener("click", () => {
	window.open("//pnppl.cc/app/wcb", "_blank");
});
wcb.style.fontFamily = "sans-serif";
wcb.style.background = "linear-gradient(to left, #ef5350, #f48fb1, #7e57c2, #2196f3, #26c6da, #43a047, #eeff41, #f9a825, #ff5722)";
wcb.style.color = "white";
wcb.style.fontWeight = "bold";
wcb.style.textShadow = "1px 1px 0px black";
menu.insertAdjacentElement("beforeend", wcb);
