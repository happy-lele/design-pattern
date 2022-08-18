package main

import "fmt"

// 基类：Bird
type Bird struct {
	Type string
}

// 鸟的类别
func (bird *Bird) Class() string {
	return bird.Type
}

// 定义了一个鸟类
type Birds interface {
	Name() string
	Class() string
}

// 鸟类：金丝雀
type Canary struct {
	Bird  // 通过匿名字段实现了继承Bird 类的 Class 属性和方法
	name string
}

func (c *Canary) Name() string {
	return c.name
}

// 鸟类：乌鸦
type Crow struct {
	Bird     // 通过匿名字段实现了继承Bird 类的 Class 属性和方法
	name string
}

func (c *Crow) Name() string {
	return c.name
}

func NewCrow(name string) *Crow {
	return &Crow{
		Bird: Bird{
			Type: "Crow",
		},
		name: name,
	}
}

func NewCanary(name string) *Canary {
	return &Canary{
		Bird: Bird{
			Type: "Canary",
		},
		name: name,
	}
}

func BirdInfo(birds Birds) {
	fmt.Printf("I'm %s, I belong to %s bird class!\n", birds.Name(), birds.Class())
}

func main() {
	canary := NewCanary("CanaryA")
	crow := NewCrow("CrowA")
	// 通过接口实现了多态
	BirdInfo(canary)
	BirdInfo(crow)
}