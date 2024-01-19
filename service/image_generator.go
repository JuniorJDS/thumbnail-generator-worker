package service

type ImageGenerator struct {
}

func NewImageGenerator() *ImageGenerator {
	return &ImageGenerator{}
}

func (i *ImageGenerator) GenerateThumbnail() error {
	return nil
}
