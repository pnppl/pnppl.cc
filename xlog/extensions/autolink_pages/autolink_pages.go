package autolink_pages

import (
	"context"
	"embed"
	"html/template"
	"path"
	"sort"
	"strings"
	"sync"

	_ "embed"

	. "github.com/emad-elsaid/xlog"
	"github.com/emad-elsaid/xlog/markdown/ast"
	east "github.com/emad-elsaid/xlog/markdown/extension/ast"
)

//go:embed templates
var templates embed.FS

type NormalizedPage struct {
	page           Page
	normalizedName string
	normalizedFileName string
}

type fileInfoByNameLength []*NormalizedPage

func (a fileInfoByNameLength) Len() int      { return len(a) }
func (a fileInfoByNameLength) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a fileInfoByNameLength) Less(i, j int) bool {
	return len(a[i].normalizedName) > len(a[j].normalizedName)
}

var autolinkPages []*NormalizedPage
var autolinkPage_lck sync.Mutex

func UpdatePagesList(Page) (err error) {
	autolinkPage_lck.Lock()
	defer autolinkPage_lck.Unlock()

	ps := MapPage(context.Background(), func(p Page) *NormalizedPage {
//------- chatbot helped, be suspicious --- //
		content := p.Content()
		lines := strings.Split(string(content), "\n")
		firstLine := strings.TrimSpace(lines[0])
//----------------------------------------- //
		return &NormalizedPage{
			page:           p,
//			normalizedName: path.Base(strings.ToLower(p.Name())),
//			we might want to write a regex to strip out more than just h1
			normalizedName: strings.Replace(strings.Replace(strings.ToLower(firstLine), "# ", "", 1), " #", "", 1),
//			it would be nice to be able to make use of both firstline and filename, ie:
//			normalizedFileName: path.Base(strings.ToLower(p.Name())),
//			but i don't understand this program or language well enough to do that yet
//			we also really want it to disambiguate by distance, ie, siblings preferred, then children, nieces, etc
//			in the meantime links by filename work perfectly fine in normal markdown, which is how i would want to write them anyway
//			like, we want to insert the equivalent of <a href=filename>firstline</a>, so that links are descriptive but also unbreakable
//			but i'm not sure how to make that process less painful. my clients are subpotimal so far.
		}
	})
	sort.Sort(fileInfoByNameLength(ps))
	autolinkPages = ps
	return
}

func countTodos(p Page) (total int, done int) {
	_, tree := p.AST()
	tasks := FindAllInAST[*east.TaskCheckBox](tree)
	for _, v := range tasks {
		total++
		if v.IsChecked {
			done++
		}
	}

	return
}

func backlinksSection(p Page) template.HTML {
	if p.Name() == Config.Index {
		return ""
	}

	pages := MapPage(context.Background(), func(a Page) Page {
		_, tree := a.AST()
		if a.Name() == p.Name() || !containLinkTo(tree, p) {
			return nil
		}

		return a
	})

	return Partial("backlinks", Locals{"pages": pages})
}

func containLinkTo(n ast.Node, p Page) bool {
	if n.Kind() == KindPageLink {
		t, _ := n.(*PageLink)
		if t.page.FileName() == p.FileName() {
			return true
		}
	}
	if n.Kind() == ast.KindLink {
		t, _ := n.(*ast.Link)
		dst := string(t.Destination)

		// link is absolute: remove /
		if strings.HasPrefix(dst, "/") {
			path := strings.TrimPrefix(dst, "/")
			if string(path) == p.Name() {
				return true
			}
		} else { // link is relative: get relative part
			// TODO: what if another folder has the same filename?
			// * just ignore that fact
			// * dont support relative paths
			// there is no way to know who is the parent folder
			base := path.Base(p.Name())
			if dst == base {
				return true
			}
		}
	}

	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		if containLinkTo(c, p) {
			return true
		}

		if c == n.LastChild() {
			break
		}
	}

	return false
}
