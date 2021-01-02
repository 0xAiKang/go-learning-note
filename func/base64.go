package main

import (
	"fmt"
	b64 "encoding/base64"
)

func main()  {
	var p = fmt.Println

	data := "abc123!?$*&()'-=@~"
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	p(sEnc)
	sDnc, _ := b64.StdEncoding.DecodeString(sEnc)
	p(string(sDnc))

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	p(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	p(string(uDec))
}