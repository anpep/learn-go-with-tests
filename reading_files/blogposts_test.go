package reading_files

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello-world.md": {Data: []byte(`Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`)},
		"hello-world2.md": {Data: []byte(`Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`)},
	}

	posts, err := NewPostsFromFS(fs)
	if err != nil {
		t.Fatal("didn't want an error but got one")
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	assertPostEquals(t, posts[0], Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})
	assertPostEquals(t, posts[1], Post{
		Title:       "Post 2",
		Description: "Description 2",
		Tags:        []string{"rust", "borrow-checker"},
		Body: `B
L
M`,
	})
}

func assertPostEquals(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
