package xlog

import (
	"errors"
	"fmt"
	"html/template"
	"path"
	"slices"
	"strings"
	"time"
	"strconv"
	"context"
	"math/rand"

	"github.com/emad-elsaid/xlog/markdown/ast"
	gast "github.com/emad-elsaid/xlog/markdown/ast"
	emojiAst "github.com/emad-elsaid/xlog/markdown/emoji/ast"
)

var helpers = template.FuncMap{
	"ago":            ago,
	"properties":     Properties,
	"links":          Links,
	"widgets":        RenderWidget,
	"commands":       Commands,
	"quick_commands": QuickCommands,
	"isFontAwesome":  IsFontAwesome,
	"includeJS":      includeJS,
	"scripts":        scripts,
	"banner":         Banner,
	"emoji":          Emoji,
	"base":           path.Base,
	"dir":            dir,
	"raw":            raw,
	"trim":           trim,
    "accesskey":      accesskey,
    "resetaccess":    resetaccess,
//    "UTCoffset":      UTCoffset,
	"datetime":       datetime,
	"dateInName":     dateInName,
	"next":           next,
	"prev":           prev,
//	"newest":         newest,
	"active":         active,
	"activeStr":      activeStr,
//	"tagIsParent":    tagIsParent,
//	"tagChildren":    tagChildren,
//	"tagParent":      tagParent,
	"noop":           noop,
	"tagId":          tagId,
	"wday":           wday,
	"randomBadge":    randomBadge,
}

var ErrHelperRegistered = errors.New("Helper already registered")

// RegisterHelper registers a new helper function. all helpers are used when compiling
// templates. so registering helpers function must happen before the server
// starts as compiling templates happened right before starting the http server.
func RegisterHelper(name string, f any) error {
	if _, ok := helpers[name]; ok {
		return ErrHelperRegistered
	}

	helpers[name] = f

	return nil
}

// A function that takes time.duration and return a string representation of the
// duration in human readable way such as "3 seconds ago". "5 hours 30 minutes
// ago". The precision of this function is 2. which means it returns the largest
// unit of time possible and the next one after it. for example days + hours, or
// Hours + minutes or Minutes + seconds...etc
func ago(t time.Time) string {
	if Config.Readonly {
		return t.Format("Monday 2 January 2006")
	}

	d := time.Since(t)

	const day = time.Hour * 24
	const week = day * 7
	const month = day * 30
	const year = day * 365
	const maxPrecision = 2

	var o strings.Builder

	if d.Seconds() < 1 {
		o.WriteString("Less than a second ")
	}

	for precision := 0; d.Seconds() > 1 && precision < maxPrecision; precision++ {
		switch {
		case d >= year:
			years := d / year
			d -= years * year
			o.WriteString(fmt.Sprintf("%d years ", years))
		case d >= month:
			months := d / month
			d -= months * month
			o.WriteString(fmt.Sprintf("%d months ", months))
		case d >= week:
			weeks := d / week
			d -= weeks * week
			o.WriteString(fmt.Sprintf("%d weeks ", weeks))
		case d >= day:
			days := d / day
			d -= days * day
			o.WriteString(fmt.Sprintf("%d days ", days))
		case d >= time.Hour:
			hours := d / time.Hour
			d -= hours * time.Hour
			o.WriteString(fmt.Sprintf("%d hours ", hours))
		case d >= time.Minute:
			minutes := d / time.Minute
			d -= minutes * time.Minute
			o.WriteString(fmt.Sprintf("%d minutes ", minutes))
		case d >= time.Second:
			seconds := d / time.Second
			d -= seconds * time.Second
			o.WriteString(fmt.Sprintf("%d seconds ", seconds))
		}
	}

	o.WriteString("ago")

	return o.String()
}

func UTCoffset(t time.Time) string {
	// offset in seconds
	_, offset := t.Zone()
	hours := offset / 3600
	minutes := (offset % 3600) / 60
	return fmt.Sprintf("%+03d%02d", hours, minutes)
}
// return datetime formatted for html <time> datetime attr
func datetime(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d%05s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), UTCoffset(t))
}
// check if date is taken from filename (otherwise is from post text)
func dateInName(name string, date string) bool {
	return strings.Contains(name, date)
}

var js = []string{}

// RegisterJS adds a Javascript library URL/path to be included in the scripts used by the template
func RegisterJS(f string) {
	if slices.Contains(js, f) {
		return
	}

	js = append(js, f)
}

// RequireHTMX registes HTML library, this helps include one version of HTMX
//func RequireHTMX() {
//	RegisterJS("/public/htmx.min.js")
//}

func includeJS(f string) template.HTML {
	RegisterJS(f)

	return ""
}

func scripts() template.HTML {
	var b strings.Builder
	for _, f := range js {
		fmt.Fprintf(&b, `<script src="%s" defer></script>`, f)
	}

	return template.HTML(b.String())
}

func IsFontAwesome(i string) bool {
	return strings.HasPrefix(i, "fa")
}

func Banner(p Page) string {
	_, a := p.AST()
	if a == nil {
		return ""
	}

	paragraph := a.FirstChild()
	if paragraph == nil || paragraph.Kind() != gast.KindParagraph {
		return ""
	}

	img := paragraph.FirstChild()
	if img == nil || img.Kind() != gast.KindImage {
		return ""
	}

	image, ok := img.(*ast.Image)
	if !ok {
		return ""
	}

	dest := string(image.Destination)
	if len(dest) == 0 || dest == "#" {
		return ""
	}

	if !(path.IsAbs(dest) || strings.HasPrefix(dest, "http")) {
		d := path.Dir(p.FileName())
		dest = path.Join("/", d, dest)
	}

	return dest
}

func Emoji(p Page) string {
	_, tree := p.AST()
	if e, ok := FindInAST[*emojiAst.Emoji](tree); ok && e != nil {
		return string(e.Value.Unicode)
	}

	return ""
}

func dir(s string) string {
	v := path.Dir(s)

	if v == "." {
		return ""
	}

	return v
}

// raw a helper to output input string as safe HTML
func raw(i string) template.HTML {
	return template.HTML(i)
}

// remove last character (fix for trailing newline issue)
func trim(i string) string {
	return i[:len(i)-1]
}

// global var helpers so we can count headings
var counter = 1
func accesskey() string {
	var returni string = ""
	if counter < 11 {
		if counter == 10 {
			returni = "0"
		} else {
			returni = strconv.Itoa(counter)
		}
		counter = counter + 1
	}
	return returni
}
func resetaccess() string {
	counter = 1
	return ""
}

var pageNames []string
func buildSiblingIndex() {
	// already built
	if len(pageNames) > 0 {
		return
	}
	allPages := Pages(context.Background());
	for p, _ := range allPages {
		if !IsIgnoredPath(allPages[p].Name()) {
			pageNames = append(pageNames, allPages[p].Name())
		}
	}
	slices.Sort(pageNames)
}
func sibling(currPage Page, prev bool) string {
	buildSiblingIndex()
	idx := slices.Index(pageNames, currPage.Name())
	if prev && idx > 0 {
		return pageNames[idx - 1]
	}
	// next
	if !prev && idx < len(pageNames) - 1 && idx >= 0 {
		return pageNames[idx + 1]
	}
	return ""
}
func prev(currPage Page) string {
	return sibling(currPage, true)
}
func next(currPage Page) string {
	return sibling(currPage, false)
}
func Newest() string {
	for i := len(pageNames) - 1; i >= 0; i -- {
		if strings.HasPrefix(pageNames[i], "20") {
			return pageNames[i]
		}
	}
	return ""
}
func active(currPage Page, label map[string]string) bool {
	labelStr := label["labelStart"] + label["labelAccel"] + label["labelEnd"]
	if currPage.Name() == labelStr || currPage.Name() == "/+/" + labelStr {
		return true
	}
	return false
}
func activeStr(currPage Page, label string) bool {
	labelMap := map[string]string{
		"labelStart": label,
		"labelAccel": "",
		"labelEnd": "",
	}
	return active(currPage, labelMap)
}
/*
func tagIsParent(tag string) bool {
	return strings.ContainsAny(tag, "⹀︱")
}
func tagChildren(tag string) []string {
	if (tagIsParent(tag)) {
		tag := strings.Split(tag, "⹀")
		return strings.Split(tag[1], "︱")
	}
	return []string { tag, }
}
func tagParent(tag string) string {
	if (tagIsParent(tag)) {
		parent := strings.Split(tag,"⹀")
		return parent[0]
	}
	return tag
}
*/

// for whitespace removal
func noop() string { return "" }

// #tagname -> #tag-tagname. for reducing ID collisions
func tagId(tag string) string {
	if len(tag) < 2 {
		return ""
	}
	nohash := tag[1:len(tag)]
	return "#tag-" + nohash
}

// Wednesday -> Wed.
func wday(day string) string {
	return day[0:3]
}

func randomBadge() template.HTML {
	badges := [][]string{
		{"bgdc-ga-j", "", "BE GAY DO CRIME on rainbow/black flag"},
		{"bgdc-tr-j", "", "BE GAY DO CRIME on trans/black flag"},
		{"bgdc-bi-j", "", "BE GAY DO CRIME on bi/black flag"},
		{"bgdc-gq-j", "", "BE GAY DO CRIME on genderqueer/black flag"},
		{"invalid-html", "", "W3C validator result: HTML ??? X"},
		{"invalid-css-red", "", "W3C validator result: CSS level ‽ X"},
		{"my-anarchy-now", "", "ANARCHY NOW! 1312"},
		{"online", "", "Cords plugged together show you're ON-LINE"},
		{"she-her", "", "she/her"},
		{"copyleft", "/about#copyleft", "Copyleft: all wrongs reversed"},
		{"steal", "https://git.gay/pnppl/pnppl.cc", "STEAL THIS SITE"},
		{"civ", "https://hrt.pnppl.cc", "INDUSTRIAL CIVILIZATION IS MY ENDOCRINE SYSTEM"},
		{"fish", "https://fishshell.com/", "<3 fish shell"},
		{"gravity-wells", "https://blueshifted.net/", "no gods, no masters, no gravity wells"},
		{"lain", "https://fauux.neocities.org/lovelain", "LET'S ALL LOVE LAIN!"},
		{"nano-88x31", "https://nano-editor.org/", "made with GNU nano"},
		{"victor-8831-wide", "https://rubjo.github.io/victor-mono/", "best viewed in Victor Mono"},
	}
	badge := badges[rand.Intn(len(badges))]
	img := badge[0]
	href := badge[1]
	alt := badge[2]
	if len(href) == 0 {
		href = "/about"
	}
	return template.HTML(fmt.Sprintf(`<a id="badge" href="%s">
							<img id="badge-img" src="/public/badges/%s.gif" alt="%s" width="88" height="31">
					</a>`, href, img, alt))
}
