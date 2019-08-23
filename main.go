/*
动态生成JAVA动态链接库
*/
// GoLibrarySO project main.go
package main

import "C"

func main() {

}

//export RunGoLibrarySo
func RunGoLibrarySo(_str *C.char) *C.char {
	//TODO 接收参数
	str := C.GoString(_str)

	//TODO 方法调用并返回值
	results := C.CString(goMethod(str))

	return results
}

func goMethod(_value string) string {
	//TODO 业务逻辑处理

	return _value
}
