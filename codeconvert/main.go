package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

// ExampleEncode string -> byte -> hex
func ExampleEncode() {
	src := []byte("Hello marvin!")
	fmt.Println(len(src), src) // 13 [72 101 108 108 111 32 109 97 114 118 105 110 33]

	dst := make([]byte, hex.EncodedLen(len(src))) // hex.EncodedLen(len(src)) = len(src) * 2
	fmt.Println(len(dst), dst)                    // 26 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]

	hex.Encode(dst, src)

	fmt.Println(string(dst)) // 48656c6c6f206d617276696e21
	fmt.Printf("%s\n", dst)  // 48656c6c6f206d617276696e21
	fmt.Printf("%x\n", 72)   // 48
}

// ExampleDecode hexString -> byte -> string
func ExampleDecode() {
	src := []byte("48656c6c6f206d617276696e21") // need even count
	fmt.Println(len(src), src)                  // 26 [52 56 54 53 54 99 54 99 54 102 50 48 54 100 54 49 55 50 55 54 54 57 54 101 50 49]

	dst := make([]byte, hex.DecodedLen(len(src))) // hex.DecodedLen(len(src)) = len(src) / 2
	fmt.Println(len(dst), dst)                    // 13 [0 0 0 0 0 0 0 0 0 0 0 0 0]

	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", dst[:n])        // Hello marvin!
	fmt.Printf("%d, %c\n", 0x48, 0x48) // 72  H
}

// ExampleEncodeToString string -> hex
func ExampleEncodeToString() {
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)

	fmt.Printf("%s\n", encodedStr)

	// Output:
	// 48656c6c6f
}

// ExampleDecodeString hexString  -> byte -> string
func ExampleDecodeString() {
	const s = "48656c6c6f20476f7068657221"
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", decoded)

	// Output:
	// Hello Gopher!
}
