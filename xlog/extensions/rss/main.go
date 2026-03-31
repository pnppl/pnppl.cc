package rss

import (
	"encoding/xml"
	"flag"
	"fmt"
	"html/template"
	"net/url"
	"slices"
	"strings"
	"time"

	. "github.com/emad-elsaid/xlog"
	"github.com/emad-elsaid/xlog/markdown/util"
)

var domain string
var description string
var limit int
const rfc822 = "Mon, 02 Jan 2006 15:04:05 -0700"

func init() {
	flag.StringVar(&domain, "rss.domain", "", "RSS domain name to be used for RSS feed. without HTTPS://")
	flag.StringVar(&description, "rss.description", "", "RSS feed description")
	flag.IntVar(&limit, "rss.limit", 30, "Limit the number of items in the RSS feed to this amount")

	RegisterExtension(RSS{})
}

type RSS struct{}

func (RSS) Name() string { return "rss" }
func (RSS) Init() {
	RegisterWidget(WidgetHead, 0, metaTag)
	RegisterBuildPage("/+/feed.rss", false)
//	RegisterLink(links)
	Get(`/+/feed.rss`, feed)
}

type rssLink struct{}

func (rssLink) Icon() string { return "fa-solid fa-rss" }
func (rssLink) Name() string { return "RSS" }
func (rssLink) Attrs() map[template.HTMLAttr]any {
	return map[template.HTMLAttr]any{
		"href": "/+/feed.rss",
	}
}
func (rssLink) Label() map[string]string {
		return map[string]string {
			"labelStart": "RSS ",
			"labelAccel": "F",
			"labelEnd": "eed",
	}
}

func links(p Page) []Command {
	return []Command{rssLink{}}
}

func metaTag(p Page) template.HTML {
	tag := `<link rel="alternate" href="/+/feed.rss" title="%s RSS" type="application/rss+xml">`
	return template.HTML(fmt.Sprintf(tag, template.JSEscapeString(Config.Sitename)))
}

type rss struct {
	Version string  `xml:"version,attr"`
	Xmlns string	`xml:"xmlns:atom,attr"`
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Language    string `xml:"language"`
	Description string `xml:"description"`
	Copyright	string `xml:"copyright"`
	WebMaster	string `xml:"webMaster"`
	AtomLink	AtomLink `xml:"atom:link"`
	Items       []Item `xml:"item"`
}

type AtomLink struct {
	Href		string `xml:"href,attr"`
	Rel			string `xml:"rel,attr"`
	Type		string `xml:"type,attr"`
}

type Item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
	Link        string `xml:"link"`
}

func feed(r Request) Output {
	f := rss{
		Version: "2.0",
		Xmlns:   "http://www.w3.org/2005/Atom",
		Channel: Channel{
			Title: Config.Sitename,
			Link: (&url.URL{
				Scheme: "http",
				Host:   domain,
				Path:   "/+/feed.rss",
			}).String(),
			Description: description,
			Copyright:	"Copyleft: All Wrongs Reversed (CC BY-SA 4.0)",
			WebMaster:	"webmaster@pnppl.cc (pnppl)",
			Language:    "en-US",
			AtomLink: AtomLink{
				Href:	(&url.URL{
							Scheme: "http",
							Host:   domain,
							Path:   "/+/feed.rss",
						}).String(),
				Rel:	"self",
				Type:	"application/rss+xml",
			},
			Items:       []Item{},
		},
	}

	pages := Pages(r.Context())
	slices.SortFunc(pages, func(a, b Page) int {
		if modtime := b.ModTime().Compare(a.ModTime()); modtime != 0 {
			return modtime
		}

		return strings.Compare(a.Name(), b.Name())
	})

	if len(pages) > limit {
		pages = pages[0:limit]
	}

	for _, p := range pages {
		properties := Properties(p)
		if ! IsIgnoredPath(p.Name()) {
			title := properties["title"].Value().(string)
			f.Channel.Items = append(f.Channel.Items, Item{
				Title:			title,
				Description:	string(p.Render()) + "\n\n<br>\n<hr>\n<br>\n" + `<a href="mailto:&#102;&#101;&#101;&#100;&#98;&#97;&#99;&#107;&#64;&#112;&#110;&#112;&#112;&#108;&#46;&#99;&#99;?subject=` + string(util.URLEscape([]byte(title), false)) + `">~> send me an email <~</a>`,
				PubDate:		timeFromName(p.Name(), p.ModTime()),
	//			LastBuildDate:	p.ModTime().Format(rfc822),
				GUID:			(&url.URL{
									Scheme: "http",
									Host:   domain,
									Path:   "/" + p.Name(),
								}).String(),
				Link:			(&url.URL{
									Scheme: "http",
									Host:   domain,
									Path:   "/" + p.Name(),
								}).String(),
			})
		}
	}

	buff, err := xml.MarshalIndent(f, "", "    ")
	if err != nil {
		return InternalServerError(err)
	}

	return PlainText(xml.Header + string(buff))
}

// This made more sense when I thought lastBuildDate was an item property, but it's a channel property
// still, might as well stick with this for now
func timeFromName(name string, mtime time.Time) string {
	location, _ := time.LoadLocation("America/New_York")
	parsedTime, err := time.ParseInLocation("2006-01-02_15-04-05", name, location)
	if err != nil {
		if len(name) >= 10 {
			parsedTime, err = time.Parse("2006-01-02", name[:10])
				if err != nil {
					parsedTime = mtime
				}
		} else {
			parsedTime = mtime
		}
	}
	return parsedTime.Format(rfc822)
}
