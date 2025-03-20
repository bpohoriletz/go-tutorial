package blogrenderer_test

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
	blogrenderer "github.com/bpohoriletz/go-tutorial/18_templating"
)

func TestRenderer(t *testing.T) {
	approvals.UseFolder("tmp")
	var (
		postRenderer *blogrenderer.PostRenderer
		err          error
		aPost        = blogrenderer.Post{
			Title:       "Hello World!",
			Description: "This is a Description",
			Body:        "This is a post",
			Tags:        []string{"go", "tdd"},
		}
	)
	if postRenderer, err = blogrenderer.NewPostRenderer(); nil != err {
		t.Fatal(err)
	}

	t.Run("Post", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})

	t.Run("Index", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []blogrenderer.Post{{Title: "Hello World"}, {Title: "Hello World 2"}}
		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		postRenderer *blogrenderer.PostRenderer
		err          error
		aPost        = blogrenderer.Post{
			Title:       "Hello World!",
			Description: "This is a Description",
			Body:        "This is a post",
			Tags:        []string{"go", "tdd"},
		}
	)
	if postRenderer, err = blogrenderer.NewPostRenderer(); nil != err {
		b.Fatal(err)
	}

	for b.Loop() {
		postRenderer.Render(io.Discard, aPost)
	}
}
