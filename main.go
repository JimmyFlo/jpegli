package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/image/draw"
	_ "golang.org/x/image/webp"

	"github.com/gen2brain/jpegli"
)

const (
	defaultMaxWidth  = 2048
	defaultMaxHeight = 1920
	compressionQuality = 75
	ext              = ".jpg"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <image.jpg> [<image.png> <image.webp>…]\n", path.Base(os.Args[0]))
		os.Exit(1)
	}

	success := 0
	for _, imgPath := range os.Args[1:] {
		if err := standardize(imgPath); err != nil {
			fmt.Printf("‼ Unable to standardize '%s': %v\n", imgPath, err)
		} else {
			success++
		}
	}

	fmt.Printf("Complete. %d images standardized\n", success)
}

func standardize(imgPath string) error {
	newImgPath, originalImgPath := newPath(imgPath)

	if originalImgPath != "" {
		if err := os.Rename(imgPath, originalImgPath); err != nil {
			return fmt.Errorf("unable to move the original file to make way for the standardized one: %w", err)
		}
		imgPath = originalImgPath
	}

	f, err := os.Open(imgPath)
	if err != nil {
		return err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("unknown image format: %w", err)
	}

	newImg := resizeImage(img, defaultMaxWidth, defaultMaxHeight)

	w, err := os.OpenFile(newImgPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("unable to create file for standardized image: %w", err)
	}

	if err := jpegli.Encode(w, newImg, &jpegli.EncodingOptions{
		Quality:           compressionQuality,
		FancyDownsampling: true,
		ProgressiveLevel:  2,
	}); err != nil {
		return fmt.Errorf("unable to encode standardized JPEGli image: %w", err)
	}

	return nil
}

func resizeImage(src image.Image, maxWidth, maxHeight int) image.Image {
	newWidth, newHeight, needsResize := resizeAspectRatio(src.Bounds().Dx(), src.Bounds().Dy(), maxWidth, maxHeight)
	if !needsResize {
		return src
	}

	dst := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))
	draw.BiLinear.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	return dst
}

func resizeAspectRatio(width, height, maxWidth, maxHeight int) (int, int, bool) {
	if width <= maxWidth && height <= maxHeight {
		return width, height, false
	}

	if width > maxWidth {
		height = height * maxWidth / width
		width = maxWidth
	}

	if height > maxHeight {
		width = width * maxHeight / height
		height = maxHeight
	}

	return width, height, true
}

func newPath(old string) (string, string) {
	base := path.Base(old)
	core := strings.TrimSuffix(base, filepath.Ext(base))
	newPath := path.Join(path.Dir(old), core+ext)

	var original string
	if strings.ToLower(old) == strings.ToLower(newPath) {
		original = path.Join(path.Dir(old), core+".original"+ext)
	}

	return newPath, original
}
