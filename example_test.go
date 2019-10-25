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
	h := zipfs.FileServer(fs)
	return http.ListenAndServe(":8080", h)
}
