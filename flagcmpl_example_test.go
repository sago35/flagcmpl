package flagcmpl_test

import (
	"flag"
	"fmt"
	"os"

	"github.com/sago35/flagcmpl"
)

func ExampleParse() {
	var verbose = flag.Bool("verbose", false, "Verbose mode.")

	flagcmpl.Parse()

	fmt.Println("verbose flag :", *verbose)
}

func ExampleParseFlagSet() {
	flags := flag.NewFlagSet("sample2", flag.ExitOnError)
	flags.Bool("verbose", false, "Verbose mode.")

	flagcmpl.ParseFlagSet(os.Args[0], flags, os.Args)
}
