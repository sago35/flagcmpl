package flagcmpl

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Parse() {

	var args []string
	for _, arg := range os.Args[1:] {
		if arg == `--completion-script-bash` {
			fmt.Println(makeCompletionBash(os.Args[0], flag.CommandLine))
			os.Exit(0)
		} else {
			args = append(args, arg)
		}
	}

	os.Args = append([]string{os.Args[0]}, args...)
	flag.Parse()
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
