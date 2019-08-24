package cliutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func existFileMock(files ...string) ExistFile {
	return func(p string) bool {
		for _, f := range files {
			if f == p {
				return true
			}
		}
		return false
	}
}

func TestFindFile(t *testing.T) {
	data := []struct {
		files []string
		names []string
		title string
		wd    string
		exp   string
		isErr bool
	}{
		{
			title: "file is not found",
			names: []string{".cmdx.yml"},
			wd:    "/tmp/foo",
			isErr: true,
		},
		{
			title: "file is not found 2",
			files: []string{"/.cmdx.yaml"},
			names: []string{".cmdx.yml"},
			wd:    "/tmp/foo",
			isErr: true,
		},
		{
			title: "file is found",
			files: []string{"/.cmdx.yaml", "/tmp/.cmdx.yaml"},
			names: []string{".cmdx.yml", ".cmdx.yaml"},
			wd:    "/tmp/foo",
			exp:   "/tmp/.cmdx.yaml",
		},
		{
			title: "file is found 2",
			files: []string{"/.cmdx.yml", "/tmp/.cmdx.yaml"},
			names: []string{".cmdx.yml", ".cmdx.yaml"},
			wd:    "/tmp/foo",
			exp:   "/tmp/.cmdx.yaml",
		},
	}
	for _, d := range data {
		t.Run(d.title, func(t *testing.T) {
			s, err := FindFile(d.wd, existFileMock(d.files...), d.names...)
			if err != nil {
				if d.isErr {
					return
				}
				assert.Nil(t, err)
				return
			}
			assert.False(t, d.isErr)
			assert.Equal(t, d.exp, s)
		})
	}
}
