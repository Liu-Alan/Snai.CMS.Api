package common

import (
	"net/http"
	"os"
	"path"
)

// NoListFS 禁止目录列出
type NoListFS struct {
	FS http.FileSystem
}

func (n NoListFS) Open(name string) (http.File, error) {
	f, err := n.FS.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if stat.IsDir() {
		// 如果目录下没有 index.html，就禁止访问
		index := path.Join(name, "index.html")
		if _, err := n.FS.Open(index); err != nil {
			return nil, os.ErrPermission
		}
	}

	return f, nil
}
