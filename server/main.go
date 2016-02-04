package main

import (
	"os"

	"net/http"
	"strings"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/wadahiro/golang-web-boilerplate/server/controllers"
)

type binaryFileSystem struct {
	fs http.FileSystem
}

func (b *binaryFileSystem) Open(name string) (http.File, error) {
	return b.fs.Open(name)
}

func (b *binaryFileSystem) Exists(prefix string, filepath string) bool {
	if p := strings.TrimPrefix(filepath, prefix); len(p) < len(filepath) {
		if _, err := b.fs.Open(p); err != nil {
			return false
		}
		return true
	}
	return false
}

func BinaryFileSystem(root string) *binaryFileSystem {
	fs := &assetfs.AssetFS{Asset, AssetDir, AssetInfo, root}
	return &binaryFileSystem{
		fs,
	}
}

func main() {
	r := gin.Default()

    // We can't use router.Static method to use '/' for static files.
    // see https://github.com/gin-gonic/gin/issues/75
	if gin.IsDebugging() {
		r.Use(static.Serve("/", static.LocalFile("assets", true)))
		// r.Static("/", "./assets")
	} else {
		r.Use(static.Serve("/", BinaryFileSystem("assets")))
		// r.StaticFS("/", assetFS())
	}

	// add routes
	r.GET("/api/home", controllers.Home)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	r.Run(":" + port)
}
