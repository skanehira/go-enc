package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/skanehira/go-enc/pkg/urlenc"
)

var encodeURL = flag.String("e", "", "encode URL")
var decodeURL = flag.String("d", "", "decode URL")

func init() {
	flag.Parse()
}

func run() (int, error) {
	if flag.NFlag() == 0 {
		flag.Usage()
		return 0, nil
	} else if *encodeURL != "" {
		if result, err := urlenc.Encode(*encodeURL); err != nil {
			return 1, err
		} else {
			fmt.Println(result)
		}
	} else if *decodeURL != "" {
		if result, err := urlenc.Decode(*decodeURL); err != nil {
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
