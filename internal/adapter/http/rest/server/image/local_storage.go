package image

import (
	"errors"
	"fmt"
	"github.com/corona10/goimagehash"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"path/filepath"
)

type Storage struct {
	basePath string
}

func NewStorage(basePath string) (*Storage, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	return &Storage{basePath: p}, nil
}

func (l *Storage) Save(path string, m image.Image) (*os.File, error) {
	fullPath := l.FullPath(path)
	dir := filepath.Dir(fullPath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	_, err = os.Stat(fullPath)
	if err == nil {
		err = os.Remove(fullPath)
		if err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}
	resizedImage, err := os.Create(fullPath)
   fmt.Println("error create  ",err)
	if err != nil {
		return nil, err
	}
	err = l.Encode(filepath.Ext(fullPath), resizedImage, m)
	fmt.Println("path ",filepath.Ext(fullPath))
	fmt.Println("error encode ",err)
	if err != nil {
		return nil, err
	}
	return resizedImage, nil
}

func (l *Storage) Get(path string) (*os.File, error) {
	fp := l.FullPath(path)
	f, err := os.Open(fp)
	if err != nil {
		return nil, err
	}
	return f, nil
}
func (l Storage) FileInfo(ext string,m image.Image) (*os.File, error) {
	resizedImage, err := os.CreateTemp("", "temp.*"+ext)
	if err != nil {
		return nil,err
	}
	defer os.Remove(resizedImage.Name()) // clean up
	err = l.Encode(ext, resizedImage, m)
	if err != nil {
		return nil, err
	}
	return resizedImage,nil

}

func (l *Storage) FullPath(path1 string) string {
	return filepath.Join(l.basePath, path1)
}
func (l *Storage) GetImageDimension(f *multipart.FileHeader) (int, int) {
	file, err := f.Open()
	//file, err := os.Open(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	images, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", file, err)
	}
	return images.Width, images.Height
}
func (l *Storage) AverageHash(f *multipart.FileHeader) (string, error) {
	file, err := f.Open()
	if err != nil {
		return "", err
	}
	img, format, err := image.Decode(file)
	fmt.Println("format ", format)
	if err != nil {
		return "", err
	}
	hash1, _ := goimagehash.AverageHash(img)
	return hash1.ToString(), nil

}
func (l *Storage) Resize(f *multipart.FileHeader, rwidth, rheiht uint) (image.Image, error) {
	file, err := f.Open()
	if err != nil {
		return nil, err
	}
	img, format, err := image.Decode(file)
	fmt.Println("format ", format)
	if err != nil {
		return nil, err
	}
	m := resize.Resize(rwidth, rheiht, img, resize.Lanczos3)
	return m, nil
}
func (l *Storage) Encode(formattype string, resizedImage *os.File, m image.Image) (err error) {
	switch formattype {
	case ".jpeg":
	case ".jpg":
		err = jpeg.Encode(resizedImage, m, nil)
	case ".gif":
		err = gif.Encode(resizedImage, m, nil)
	case ".png":
		err = png.Encode(resizedImage, m)
	default:
		err = errors.New("Unsupported format type!")
	}
	return
}
func Decode(formattype string, file *os.File) (img image.Image, err error) {
	switch formattype {
	case ".jpeg":
	case ".jpg":
		img, err = jpeg.Decode(file)
	case ".gif":
		img, err = gif.Decode(file)
	case ".png":
		img, err = png.Decode(file)
	default:
		err = errors.New("Unsupported format type!")
	}
	return
}
