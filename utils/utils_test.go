package utils

import (
	"testing"
)

// 定义两个结构体
type StructA struct {
	Name  string
	Age   int
	Email string
}

type StructB struct {
	Name  string
	Age   int
	Phone string
}

func TestCopyStruct(t *testing.T) {
	// 初始化两个结构体
	a := &StructA{Name: "John", Age: 30, Email: "john@example.com"}
	b := &StructB{}

	// 使用反射赋值
	AssignSimilarFields(a, b)

	// 打印赋值后的结构体
	t.Logf("StructB after assignment: %+v\n", b)

}
