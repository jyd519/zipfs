package zipfs_test

import (
	"net/http"

	"github.com/jyd519/zipfs"
)

func Example() error {
	fs, err := zipfs.New("testdata/testdata.zip")
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", zipfs.FileServer(fs))
}
