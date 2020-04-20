# Go BinaryPack

[![Build Status](https://travis-ci.org/roman-kachanovsky/go-binary-pack.svg?branch=master)](https://travis-ci.org/roman-kachanovsky/go-binary-pack)
[![GoDoc](https://godoc.org/github.com/roman-kachanovsky/go-binary-pack/binary-pack?status.svg)](http://godoc.org/github.com/roman-kachanovsky/go-binary-pack/binary-pack)

BinaryPack is a simple Golang library which implements some functionality of Python's [struct](https://docs.python.org/2/library/struct.html) package.

**Install**

`go get github.com/roman-kachanovsky/go-binary-pack/binary-pack`

**Go mod**

require github.com/roman-kachanovsky/go-binary-pack v0.0.0-20170214094030-e260e0dc6732 // indirect


**How to use**

```go
package main

/*
require github.com/roman-kachanovsky/go-binary-pack v0.0.0-20170214094030-e260e0dc6732 // indirect
 */
import (
	"fmt"
	"reflect"

	. "github.com/roman-kachanovsky/go-binary-pack/binary-pack"
)

func main() {
	// 要编码的结构体字段的数据格式
	format := []string{"I", "?", "d", "6s"}

	// 要打包的数据字段
	values := []interface{}{4, true, 3.14, "Golang"}

	// 创建二进制数据包
	bp := new(BinaryPack)

	// 封装数据到数据包中
	data, _ := bp.Pack(format, values)
	fmt.Printf("编码后:%+v\n", data)


	// 拆包数据到[]interface{}
	unpacked_values, _ := bp.UnPack(format, data)

	if !reflect.DeepEqual(unpacked_values, values) {
		fmt.Printf("Unpacked %v != original %v", unpacked_values, values)
	}
	fmt.Printf("解码后:%+v\n", unpacked_values)

	// 按照格式包来计算二进制包的大小
	size, _ := bp.CalcSize(format)

	if size != len(data) {
		fmt.Printf("Size(%v) != %v", size, len(data))
	}
	fmt.Printf("二进制包字节数:%+v\n", size)
}
/*
编码后:[4 0 0 0 1 31 133 235 81 184 30 9 64 71 111 108 97 110 103]
解码后:[4 true 3.14 Golang]
二进制包字节数:19

 */



```
