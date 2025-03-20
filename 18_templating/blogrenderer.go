package blogrenderer

import (
	"embed"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")

	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

type PostRenderer struct {
	templ *template.Template
}

func (p *PostRenderer) RenderIndex(to io.Writer, posts []Post) error {
	var decoratedPosts []PostDecorator

	for i := range posts {
		decoratedPosts = append(decoratedPosts, newPostDecorator(posts[i]))
	}
	if err := p.templ.ExecuteTemplate(to, "index.gohtml", decoratedPosts); err != nil {
		return err
	}

	return nil
}

func (p *PostRenderer) Render(to io.Writer, post Post) error {
	if err := p.templ.ExecuteTemplate(to, "blog.gohtml", newPostDecorator(post)); err != nil {
		return err
	}

	return nil
}

type PostDecorator struct {
	Title       string
	Body        string
	Description string
	Tags        []string
	post        Post
}

func (p PostDecorator) SanitizedTitle() string {
	return strings.ToLower(strings.Replace(p.post.Title, " ", "-", -1))
}

func (p PostDecorator) HtmlBody() template.HTML {
	html := bluemonday.UGCPolicy().Sanitize(p.Body)

	return template.HTML(string(markdown.ToHTML([]byte(html), nil, nil)))
}

func newPostDecorator(p Post) PostDecorator {
	return PostDecorator{
		Title:       p.Title,
		Body:        p.Body,
		Description: p.Description,
		Tags:        p.Tags,
		post:        p,
	}
}

type Post struct {
	Title       string
	Body        string
	Description string
	Tags        []string
}
