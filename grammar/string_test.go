package main

import "testing"

func TestString(t *testing.T)  {
	var s1, s2, s3 string
	t.Log(s1)		//  初始化默认为零值——空字符串""

	s1 = "hello"
	t.Log(len(s1))	// 输出hello 的byte 数，而不是字符数
	// s1[1] = "3"   // string 是不可变的 byte slice
	// t.Logf("hello 的UTF8：%x", s1)

	s2 = "\xE4\xB8\xA5"
	// s = "\xE4\xB8\xBB\xA5"		// 可以存储任何二进制数据，尽管它是一个错误的编码，所以会出现乱码
	t.Log(s2)
	t.Log(len(s2))

	// rune 这个数据类型可以取出字符串中的 Unicode 编码
	s3 = "中"
	// byte 这个数据类型可以取出字符串的 UTF8 存储
	r := []rune(s3)
	b := []byte(s3)
	t.Log(b)			// [228 184 173]
	t.Logf("中 的Unicode 编码：%x", r[0])
	t.Logf("中 的UTF8 存储：%X", s3)		// [0xE4, 0xB8, 0xAD]
}