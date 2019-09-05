package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/skanehira/go-enc/pkg/base64"
)

var encodeString = flag.String("e", "", "encode to base64")
var decodeString = flag.String("d", "", "decode to string")
var b64 = base64.New()

func init() {
	flag.Parse()
}

func run() (int, error) {
	if *encodeString == "" && *decodeString == "" {
		flag.PrintDefaults()
	} else if *encodeString != "" {
		fmt.Println(b64.Encode(*encodeString))
	} else if *decodeString != "" {
		result, err := b64.Decode(*decodeString)
		if err != nil {
			return 1, err
		}
		fmt.Println(result)
	}

	return 0, nil
}

func main() {
	exitCode, err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	os.Exit(exitCode)
}
