////////////////////////////////////////////////////////////////////////////////
// 享元模式
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

// 享元模式从对象中剥离出不发生改变且多个实例需要的重复数据，独立出一个享元
// 使多个对象共享，从而节省内存以及减少对象数量
type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

////////////////////////////////////////////////////////////////////////////////
var imageFactory *ImageFlyweightFactory

func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			maps: make(map[string]*ImageFlyweight),
		}
	}
	return imageFactory
}

func (this *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := this.maps[filename]
	if image == nil {
		image = NewImageFlyweight(filename)
		this.maps[filename] = image
	}

	return image
}

type ImageFlyweight struct {
	data string
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	// Load image file
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data: data,
	}
}

func (this *ImageFlyweight) Data() string {
	return this.data
}

type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{
		ImageFlyweight: image,
	}
}

func (this *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", this.Data())
}

func main() {
	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")
	fmt.Println(viewer1 != viewer2) // => true
	viewer1.Display()
	viewer2.Display()
}
