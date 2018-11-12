package vgimage

import (
	"fmt"
	"image/png"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const ImagePath = "/Users/vikas.goel/Pictures"

type ImagePixel struct {
	x, y int
	r, g, b, a uint32
}

type Image struct {
	Name string
	Width int
	Height int
	Pixels []ImagePixel
}

func main() {
	start := time.Now()

	files := getImageFiles(ImagePath)

	var wg sync.WaitGroup

	ch := make(chan *Image)

	for _, f := range files {
		wg.Add(1)
		go func(imgfile string) {
			defer wg.Done()
			ch <- getImage(imgfile)
		}(f)
	}

	go func(ch chan *Image, wg *sync.WaitGroup) {
		wg.Wait()
		close(ch)
	}(ch, &wg)

	images := addImages(0, ch)
	printImages(images)

	fmt.Println("Time:", time.Since(start), "for", len(images), "images.")
}

func getImageFiles(path string) (imgfiles []string) {
	wfn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fmt.Println(path)
		imgfiles = append(imgfiles, path)
		return nil
	}

	filepath.Walk(path, wfn)
	return
}

func getImage(pngfile string) (image *Image) {
	fimg, err := os.Open(pngfile)
	if err != nil {
		return
	}
	defer fimg.Close()

	img, err := png.Decode(fimg)
	if err != nil {
		return
	}

	bounds := img.Bounds()

	image = new(Image)
	image.Name = pngfile
	image.Width, image.Height = bounds.Dx(),bounds.Dy()
	//fmt.Printf("Image HxW %v %v\n", bounds.Dx(),bounds.Dy())

	image.Pixels = make([]ImagePixel, bounds.Dx()*bounds.Dy())
	for i := 0; i < bounds.Dx()*bounds.Dy(); i++ {
		x := i % bounds.Dx()
		y := i / bounds.Dx()

		r, g, b, a := img.At(x, y).RGBA()
		//fmt.Printf("RGBA(%v,%v) %v %v %v %v\n", x, y, r, g, b, a)

		image.Pixels[i].x = x
		image.Pixels[i].y = y
		image.Pixels[i].r = r
		image.Pixels[i].b = b
		image.Pixels[i].g = g
		image.Pixels[i].a = a
	}

	return
}

func addImages(debug int, ch <-chan *Image) (images []*Image) {
	for img := range ch {
		if img != nil {
			images = append(images, img)
		}
	}

	if debug != 0 {
		for _, img := range images {
			fmt.Printf("Image:%v %vx%v\n", img.Name, img.Width, img.Height)
			for _, pix := range img.Pixels {
				fmt.Printf("(%v,%v) {%v, %v, %v, %v}\n",
					pix.x, pix.y, pix.r, pix.g, pix.b, pix.a)
			}
		}
	}

	return
}

func printImages(images []*Image) {
	for _, img := range images {
		fmt.Printf("Image:%v %vx%v\n", img.Name, img.Width, img.Height)
		for _, pix := range img.Pixels {
			fmt.Printf("(%v,%v) {%v, %v, %v, %v}\n",
				pix.x, pix.y, pix.r, pix.g, pix.b, pix.a)
		}
	}
}
