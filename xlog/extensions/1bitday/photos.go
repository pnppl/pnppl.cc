package photos

import (
	"embed"
	"html/template"
//	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/emad-elsaid/xlog/markdown/ast"

	"github.com/emad-elsaid/types"
	"github.com/emad-elsaid/xlog"
	"github.com/emad-elsaid/xlog/extensions/shortcode"
	_ "golang.org/x/image/webp"
)

//go:embed templates
var templates embed.FS

var supportedExt = types.Slice[string]{".jpg", ".jpeg", ".gif", ".png"}

func init() {
	xlog.RegisterExtension(Photos{})
}

type Photos struct{}

func (Photos) Name() string { return "photos" }
func (Photos) Init() {
	shortcode.RegisterShortCode("1bitday", shortcode.ShortCode{Render: photosShortcode("1bitday")})
	xlog.RegisterTemplate(templates, "templates")
}

type Photo struct {
	Original  string
}

func (p *Photo) Name() string {
	base := path.Base(p.Original)
	ext := path.Ext(base)
	return base[:len(base)-len(ext)]
}

func (*Photo) FileName() string         { return "" }
func (*Photo) Exists() bool             { return false }
func (*Photo) Content() xlog.Markdown   { return "" }
func (*Photo) Delete() bool             { return false }
func (*Photo) Write(xlog.Markdown) bool { return false }
func (*Photo) ModTime() time.Time       { return time.Time{} }
func (*Photo) AST() ([]byte, ast.Node)  { return nil, nil }
func (p *Photo) Render() template.HTML {
	return xlog.Partial("photo", xlog.Locals{"photo": p})
}

func NewPhoto(path string) (*Photo, error) {
//	stat, err := os.Stat(path)
//	if err != nil {
//		return nil, err
//	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return &Photo{
		Original:  path,
	}, nil
}

func photosShortcode(tpl string) func(xlog.Markdown) template.HTML {
	return func(input xlog.Markdown) template.HTML {
		p := strings.TrimSpace(string(input))

		photos := []*Photo{}

		err := filepath.WalkDir(p, func(file string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.Type().IsRegular() && supportedExt.Include(strings.ToLower(path.Ext(file))) {
				photo, err := NewPhoto(file)
				if err != nil {
					return err
				}

				photos = append(photos, photo)
			}

			return nil
		})

		if err != nil {
			return template.HTML(err.Error())
		}

		slices.SortFunc(photos, func(i, j *Photo) int {
//			return j.Time.Compare(i.Time)
			return strings.Compare(i.Name(), j.Name())
		})

		return xlog.Partial(tpl, xlog.Locals{
			"photos": photos,
		})
	}
}

func resizeHandler(r xlog.Request) xlog.Output {
	return func(w xlog.Response, r xlog.Request) { }
}

func photoHandler(r xlog.Request) xlog.Output {
	return func(w xlog.Response, r xlog.Request) { }
}
