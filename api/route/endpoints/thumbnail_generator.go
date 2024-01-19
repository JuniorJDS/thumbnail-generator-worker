package endpoints

import "net/http"

type ThumbnailGeneratorRoute struct{}

func NewThumbnailGeneratorRoute() *ThumbnailGeneratorRoute {
	return &ThumbnailGeneratorRoute{}
}

func (t *ThumbnailGeneratorRoute) Generate(w http.ResponseWriter, r *http.Request) {
}
