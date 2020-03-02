package interfaces

import (
	"fmt"
	"reflect"
)

type A struct {
	Parent interface{}
}

func (self A) Run()  {
	c := reflect.ValueOf(self.Parent)
	method := c.MethodByName("Test")
	fmt.Println(method.IsValid())
}

type B struct {
	A
}

func (self B) Test(s string)  {
	fmt.Println("b")
}
// 子类中注册父类的方法
func (self B) Run()  {
	self.A.Run()
}
