package main

import b64 "encoding/base64"
import "fmt"

func main() {

	var str string
	fmt.Println("输入需要进行Base64编码的字符：")
	fmt.Scanf("%s",&str)
	// 读入将要编解码的字符串。

	// Go 同时支持标准的和 URL 兼容的 base64 格式。编码需要
	// 使用 `[]byte` 类型的参数，所以要将字符串转成此类型。
	sEnc := b64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println("编码结果：",sEnc)
	fmt.Println()

	// 解码可能会返回错误，如果不确定输入信息格式是否正确，
	// 那么，你就需要进行错误检查了。
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println("解码结果：",string(sDec))
	fmt.Println()

	/* 使用 URL 兼容的 base64 格式进行编解码。
	uEnc := b64.URLEncoding.EncodeToString([]byte(str))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
	*/
}