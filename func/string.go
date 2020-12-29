package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main()  {
	//  这些都是strings 的包函数，这意味着我们需要在调用函数时，将字符串作为第一个参数进行传递
	p("Contains: ", s.Contains("boo", "o"))   // true
	p("Count: ", s.Count("boo", ""))			 //
	p("HasPrefix: ", s.HasPrefix("boo", "b"))  // true
	p("HasSuffix: ", s.HasSuffix("boo", "c"))  // false
	p("Index: ", s.Index("boo", "boo"))		  // 0
	p("Index: ", s.Index("boo", "o"))		  // 1
	p("Index: ", s.Index("boo", "c"))		  // -1
	p("Join: ", s.Join([]string{"a", "b"}, "-"))	  // a-b
	p("Split: ", s.Split("a-b", "-"))		      // [a b]
	p("Repeat: ", s.Repeat("boo", 2))			  // booboo
	p("Replace: ", s.Replace("boo", "b", "0xb", -1)) // 0xboo
	p("Replace: ", s.Replace("boo", "b", "0xb", 1)) // 0xboo
	p("ToLower: ", s.ToLower("BOO"))	// boo
	p("ToUpper: ", s.ToUpper("boo"))	// BOO
	p("Len: ", len("boo"))	// 3
	p("Char: ", "boo"[1])		// 111
}