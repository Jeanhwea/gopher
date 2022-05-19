package main

import (
	"errors"
	"fmt"
	"math"
)

// 直接抛出错误
func circleArea0(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero!")
	}
	return math.Pi * radius * radius, nil
}

// 添加额外的错误信息
func circleArea1(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius is %0.2f!", radius)
	}
	return math.Pi * radius * radius, nil
}

func test01() {
	r1 := -20.0
	area, err := circleArea1(r1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f\n", area)
}

////////////////////////////////////////////////////////////////////////////////
// 使用结构体保存错误信息
type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleArea2(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func test02() {
	r2 := -23.9
	area, err := circleArea2(r2)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle %0.2f", area)
}

////////////////////////////////////////////////////////////////////////////////
// 使用结构体和帮助函数来处理
type areaError2 struct {
	err    string
	length float64
	width  float64
}

func (e *areaError2) Error() string {
	return e.err
}

func (e *areaError2) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError2) widthNagative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}
	if err != "" {
		return 0, &areaError2{err, length, width}
	}
	return length * width, nil
}

func test03() {
	length, width := -5.0, -9.0
	area, err := rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaError2); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)
			}
			if err.widthNagative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)
			}
		}
		return
	}
	fmt.Println("area of rect", area)
}

func main() {
	// test01()
	// test02()
	test03()
}
