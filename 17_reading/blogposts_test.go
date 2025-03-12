package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/bpohoriletz/go-tutorial/17_reading"
)

type StubFailingFS struct {
	/* data */
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("Reliable failure")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("Data", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello World!
The body starts after the '---'`
			scondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
Hello World!
The body starts? 
After the '---'`
		)
		fs := fstest.MapFS{
			"hello world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(scondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if nil != err {
			t.Fatal(err)
		}

		assertPost(t, posts[0], blogposts.Post{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}, Body: "Hello World!\nThe body starts after the '---'"})
	})

	t.Run("Failed to open error", func(t *testing.T) {
		ffs := StubFailingFS{}

		_, err := blogposts.NewPostsFromFS(ffs)

		if nil == err {
			t.Error("Where is my error?")
		}
	})
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %+v, expected %+v", got, want)
	}
}
