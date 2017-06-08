package flagcmpl_test

import (
	"flag"
	"fmt"
	"os"

	"github.com/sago35/flagcmpl"
)

func ExampleParse() {
	var verbose = flag.Bool("verbose", false, "Verbose mode.")

	err := flagcmpl.Parse()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("verbose flag :", *verbose)
}

func ExampleParseFlagSet() {
	flags := flag.NewFlagSet("sample2", flag.ExitOnError)
	flags.Bool("verbose", false, "Verbose mode.")

	err := flagcmpl.ParseFlagSet(os.Args[0], flags, os.Args)
	if err != nil {
		panic(err.Error())
	}
}
