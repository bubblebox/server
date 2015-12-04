package main

import "net/http"

// EmberAssetFS wraps AssetFS to render index.html
// when a path cannot be found. This is what Ember.js apps like.
type EmberAssetFS struct {
	assetFS http.FileSystem
}

func (fs *EmberAssetFS) Open(name string) (http.File, error) {
	file, err := fs.assetFS.Open(name)

	if err != nil {
		file, _ = fs.assetFS.Open("index.html")
	}

	return file, nil
}

func emberAssetFS() *EmberAssetFS {
	return &EmberAssetFS{assetFS: assetFS()}
}
