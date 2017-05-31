# flagcmpl

Package flagcmpl adds completion to flag package.

## Usage

Use `flagcmpl.Parse()` instead of `flag.Parse()`.

    package main

    import "flag"
    import "github.com/sago35/go-flagcmpl"

    var verbose = flag.Bool("verbose", false, "Verbose mode.")

    func main() {
        flagcmpl.Parse()
    }

Add your bash_profile (or equivalent).

    eval "$(your-cli-tool --completion-script-bash)"

By ending your argv with `--`, hints for flags will be shown.

## Install

    go get github.com/sago35/go-flagcmpl

## Licence

[MIT](http://opensource.org/licenses/mit-license.php)

## Author

[sago35](https://github.com/sago35)

