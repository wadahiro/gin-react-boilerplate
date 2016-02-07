package main

import (
	"os"

	"net/http"
	"strings"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/nu7hatch/gouuid"
	"github.com/wadahiro/gin-react-boilerplate/server/controllers"
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

	r.LoadHTMLGlob("server/templates/*")

	// We can't use router.Static method to use '/' for static files.
	// see https://github.com/gin-gonic/gin/issues/75
	// r.StaticFS("/", assetFS())
	r.Use(static.Serve("/", BinaryFileSystem("assets")))

	r.Use(func(c *gin.Context) {
		id, _ := uuid.NewV4()
		c.Set("uuid", id)
	})

	// add routes
	r.GET("/api/home", controllers.Home)

	react := NewReact(
		"assets/js/bundle.js",
		true,
		r,
	)

	r.GET("/test", react.Handle)

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	r.Run(":" + port)
}
