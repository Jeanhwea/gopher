////////////////////////////////////////////////////////////////////////////////
// 抽象工厂模式
////////////////////////////////////////////////////////////////////////////////
package main

import "fmt"

// 定义存取数据的抽象工厂
type DataLoaderAbstractFactory interface {
	CreateLoader() DataLoader
}

type DataLoader interface {
	Load()
}

////////////////////////////////////////////////////////////////////////////////
type XMLLoader struct{}

func (this *XMLLoader) Load() {
	fmt.Println("Load data from xml file.")
}

type XMLFactory struct{}

func (this *XMLFactory) CreateLoader() DataLoader {
	return &XMLLoader{}
}

////////////////////////////////////////////////////////////////////////////////
type DBLoader struct{}

func (this *DBLoader) Load() {
	fmt.Println("Select data from database.")
}

type DBFactory struct{}

func (this *DBFactory) CreateLoader() DataLoader {
	return &DBLoader{}
}

func main() {
	var loaderFactor DataLoaderAbstractFactory

	loaderFactor = &XMLFactory{}
	loaderFactor.CreateLoader().Load()

	loaderFactor = &DBFactory{}
	loaderFactor.CreateLoader().Load()
}
