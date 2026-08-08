package storage

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/draw"
)

type ImageConfig struct {
	MaxWidth  int
	MaxHeight int
	Quality   int
}

var DefaultImageConfig = ImageConfig{
	MaxWidth:  1200,
	MaxHeight: 1200,
	Quality:   80,
}

func (s *FileStorage) CompressImage(data []byte) ([]byte, error) {
	return CompressImageWithConfig(data, DefaultImageConfig)
}

func CompressImageWithConfig(data []byte, config ImageConfig) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty image data")
	}

	contentType := DetectContentType(data)
	if !IsImage(contentType) {
		return data, nil
	}

	img, format, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("failed to decode image: %w", err)
	}

	if config.MaxWidth > 0 || config.MaxHeight > 0 {
		img = ResizeImage(img, config.MaxWidth, config.MaxHeight)
	}

	var buf bytes.Buffer

	switch format {
	case "jpeg", "jpg":
		quality := config.Quality
		if quality == 0 {
			quality = 80
		}
		err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
	case "png":
		err = png.Encode(&buf, img)
	case "gif":
		err = gif.Encode(&buf, img, nil)
	default:
		quality := config.Quality
		if quality == 0 {
			quality = 80
		}
		err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality})
	}

	if err != nil {
		return nil, fmt.Errorf("failed to encode image: %w", err)
	}

	return buf.Bytes(), nil
}

func ResizeImage(img image.Image, maxWidth, maxHeight int) image.Image {
	if maxWidth == 0 && maxHeight == 0 {
		return img
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	if (maxWidth == 0 || width <= maxWidth) && (maxHeight == 0 || height <= maxHeight) {
		return img
	}

	newWidth, newHeight := CalculateNewSize(width, height, maxWidth, maxHeight)

	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.ApproxBiLinear.Scale(dst, dst.Bounds(), img, bounds, draw.Over, nil)

	return dst
}

func CalculateNewSize(width, height, maxWidth, maxHeight int) (int, int) {
	if maxWidth == 0 && maxHeight == 0 {
		return width, height
	}

	if maxWidth == 0 {
		ratio := float64(maxHeight) / float64(height)
		return int(float64(width) * ratio), maxHeight
	}

	if maxHeight == 0 {
		ratio := float64(maxWidth) / float64(width)
		return maxWidth, int(float64(height) * ratio)
	}

	widthRatio := float64(maxWidth) / float64(width)
	heightRatio := float64(maxHeight) / float64(height)

	ratio := widthRatio
	if heightRatio < ratio {
		ratio = heightRatio
	}

	newWidth := int(float64(width) * ratio)
	newHeight := int(float64(height) * ratio)

	return newWidth, newHeight
}
