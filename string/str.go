package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func strDemo() {

	p("Contains:  ", s.Contains("unitTest", "es"))
	p("Count:     ", s.Count("unitTest", "t"))
	p("HasPrefix: ", s.HasPrefix("unitTest", "te"))
	p("HasSuffix: ", s.HasSuffix("unitTest", "st"))
	p("Index:     ", s.Index("unitTest", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1)) // replace all
	p("Replace:   ", s.Replace("foo", "o", "0", 1))  // replace one times
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("unitTest"))
	p("Len:		  ", len("hello"))
	p("Char:	  ", "hello"[1], string("hello"[1]))

}

func main() {
	strDemo()
}
