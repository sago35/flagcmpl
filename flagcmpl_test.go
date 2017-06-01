package flagcmpl

import (
	"flag"
	"testing"
)

func TestNormal(t *testing.T) {
	fs := flag.NewFlagSet("sample", flag.ExitOnError)
	fs.String("species", "gopher", "the species we are studying")
	fs.String("gopher_type", "defaultGopher", "usage")
	fs.String("g", "defaultGopher", "usage")
	fs.String("deltaT", "default", "comma-separated list of intervals to use between events")

	expected := `_sample()
{
    local cur=${COMP_WORDS[COMP_CWORD]}

    case "$cur" in
        --*)
            COMPREPLY=( $( compgen -W "--deltaT --g --gopher_type --species" -- $cur ) )
            ;;
        *)
            COMPREPLY=( $( compgen -f -- $cur ) )
            ;;
    esac
}

complete -F _sample sample
`

	output := makeCompletionBash("sample", fs)

	if output != expected {
		t.Errorf("not match\n[%s]\n[%s]", output, expected)
	}
}

func ExampleParse() {
	var verbose = flag.Bool("verbose", false, "Verbose mode.")

	flagcmpl.Parse()
}

func ExampleParseFlagSet() {
	flags := flag.NewFlagSet("sample2", flag.ExitOnError)
	flags.Bool("verbose", false, "Verbose mode.")

	flagcmpl.ParseFlagSet(os.Args[0], flags, os.Args)
}
