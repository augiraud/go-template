// Command go-template prints FizzBuzz from 1 to COUNT.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strconv"

	"github.com/augiraud/go-template/internal/fizzbuzz"
)

const defaultCount = 15

// version is set at build time by goreleaser and the Docker image. Without it, go build uses the git tag.
var version string

func main() {
	err := run(os.Args[1:], os.Stdout, os.Stderr)
	switch {
	case err == nil, errors.Is(err, flag.ErrHelp):
	default:
		_, _ = fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run(args []string, stdout, stderr io.Writer) error {
	flags := flag.NewFlagSet("go-template", flag.ContinueOnError)
	flags.SetOutput(stderr)
	flags.Usage = func() {
		_, _ = fmt.Fprintf(flags.Output(), "Usage: go-template [-version] [COUNT]\n\n")
		_, _ = fmt.Fprintf(flags.Output(), "Print FizzBuzz from 1 to COUNT (default: %d).\n\n", defaultCount)
		flags.PrintDefaults()
	}
	showVersion := flags.Bool("version", false, "print the version and exit")
	if err := flags.Parse(args); err != nil {
		return err
	}

	if *showVersion {
		_, err := fmt.Fprintln(stdout, buildVersion())
		return err
	}

	count := defaultCount
	switch flags.NArg() {
	case 0:
	case 1:
		n, err := strconv.Atoi(flags.Arg(0))
		if err != nil {
			return fmt.Errorf("COUNT must be a whole number, got %q", flags.Arg(0))
		}
		count = n
	default:
		return errors.New("too many arguments, see -help")
	}

	for i := range count {
		if _, err := fmt.Fprintln(stdout, fizzbuzz.Of(i+1)); err != nil {
			return err
		}
	}
	return nil
}

func buildVersion() string {
	if version != "" {
		return version
	}
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}
	return "(devel)"
}
