package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/skanehira/go-enc/pkg/base64"
)

var encodeString = flag.String("e", "", "encode to base64")
var decodeString = flag.String("d", "", "decode to string")

func init() {
	flag.Parse()
}

func run() (int, error) {
	if flag.NFlag() == 0 {
		flag.Usage()
		return 0, nil
	} else if *encodeString != "" {
		fmt.Println(base64.Encode([]byte(*encodeString)))
	} else if *decodeString != "" {
		if result, err := base64.Decode(*decodeString); err != nil {
			return 1, err
		} else {
			fmt.Println(result)
		}
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
