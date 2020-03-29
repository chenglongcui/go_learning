package type_test

import "testing"

type MyInt int64

// go 不支持隐式类型转换
func TestImplicit(t *testing.T) {
	var a int64 = 1
	var b MyInt
	b = MyInt(a)
	t.Log(a, b)
}

// go 不支持指针算数运算
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// go 不支持指针算数运算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

// 字符串的初始化是空字符串，不是nil
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
