package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
	"testing/fstest"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fsys fstest.MapFS) ([]Post, error) {
	dir, err := fsys.ReadDir(".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, entry := range dir {
		post, err := newPostFromFS(fsys, entry.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func newPostFromFS(fsys fs.FS, fileName string) (Post, error) {
	file, err := fsys.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	return newPost(file)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(r io.Reader) (Post, error) {
	scanner := bufio.NewScanner(r)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore a line

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	return strings.TrimSuffix(buf.String(), "\n")
}
