package blogposts

import (
	"io/fs"
)

func NewPostsFromFS(tfs fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(tfs, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post

	for _, f := range dir {
		post, err := getPost(tfs, f.Name())
		if nil != err {
			return []Post{}, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(tfs fs.FS, name string) (Post, error) {
	postFile, err := tfs.Open(name)
	if nil != err {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}
