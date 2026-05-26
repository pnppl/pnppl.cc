// move the About button into the Console menu
const about = document.getElementById("about-button")
const menu = document.querySelector("#menu-section nav");
const base = document.getElementById("base-button");
menu.insertBefore(about, base)
