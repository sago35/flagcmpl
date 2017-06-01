/*
	Package flagcmpl adds completion to flag package.

	Usage:

	Use `flagcmpl.Parse()` instead of `flag.Parse()`.

		package main

		import "flag"
		import "github.com/sago35/flagcmpl"

		var verbose = flag.Bool("verbose", false, "Verbose mode.")

		func main() {
			flagcmpl.Parse()
		}

	Or you can use `flag.FlagSet()`.

		package main

		import (
			"flag"
			"github.com/sago35/flagcmpl"
			"os"
		)

		func main() {
			flags := flag.NewFlagSet("sample2", flag.ExitOnError)
			flags.Bool("verbose", false, "Verbose mode.")

			flagcmpl.ParseFlagSet(os.Args[0], flags, os.Args)
		}

	Add your bash_profile (or equivalent).

		eval "$(your-cli-tool --completion-script-bash)"

	By ending your argv with `--`, hints for flags will be shown.
*/
package flagcmpl

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Parse parses the command-line flags from os.Args[1:].
// Generate completion bash script if os.Args[1:] has `--completion-script-bash`.
func Parse() {
	ParseFlagSet(os.Args[0], flag.CommandLine, os.Args)
}

// ParseFlagSet parses the command-line flags from appName, fs and argv.
// Generate completion bash script if args[1:] has `--completion-script-bash`.
func ParseFlagSet(appName string, fs *flag.FlagSet, args []string) {

	var argsNew []string
	for _, arg := range args[1:] {
		if arg == `--completion-script-bash` {
			fmt.Println(makeCompletionBash(args[0], fs))
			os.Exit(0)
		} else {
			argsNew = append(argsNew, arg)
		}
	}

	argsNew = append([]string{args[0]}, argsNew...)
	fs.Parse(argsNew[1:])
}

func makeCompletionBash(app string, fs *flag.FlagSet) string {
	var flags []string
	fs.VisitAll(func(f *flag.Flag) {
		flags = append(flags, fmt.Sprintf("--%s", f.Name))
	})

	app = filepath.Base(app)

	w := new(bytes.Buffer)
	fmt.Fprintf(w, "_%s()\n", app)
	fmt.Fprintf(w, "{\n")
	fmt.Fprintf(w, "    local cur=${COMP_WORDS[COMP_CWORD]}\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "    case \"$cur\" in\n")
	fmt.Fprintf(w, "        --*)\n")
	fmt.Fprintf(w, "            COMPREPLY=( $( compgen -W \"%s\" -- $cur ) )\n", strings.Join(flags, " "))
	fmt.Fprintf(w, "            ;;\n")
	fmt.Fprintf(w, "        *)\n")
	fmt.Fprintf(w, "            COMPREPLY=( $( compgen -f -- $cur ) )\n")
	fmt.Fprintf(w, "            ;;\n")
	fmt.Fprintf(w, "    esac\n")
	fmt.Fprintf(w, "}\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "complete -F _%s %s\n", app, app)

	return w.String()

}
