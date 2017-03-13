package dump

import (
	"flag"
	"fmt"
	"os"
)

// setupFlags initializes the instance configuration
func (c *Command) setupFlags(args []string) error {

	// init config if needed
	if c.config == nil {
		c.config = new(config)
	}

	// init flagset
	cmdFlags := flag.NewFlagSet("dump", flag.ContinueOnError)
	cmdFlags.Usage = func() { fmt.Fprint(os.Stdout, c.Help()); os.Exit(0) }

	// declare flags
	cmdFlags.StringVar(&c.config.fileName, "file", "consul.bak",
		"Destination file target")
	cmdFlags.StringVar(&c.config.cryptKey, "key", "password",
		"Passphrase for data encryption and signature validation")
	cmdFlags.BoolVar(&c.config.plainDump, "plain", false,
		"Dump a reduced set of information")
	cmdFlags.BoolVar(&c.config.acls, "acls", false,
		"Specified file is an ACL token backup file")
	cmdFlags.BoolVar(&c.config.queries, "queries", false,
		"Specified file is a prepared query backup file")

	// parse flags and ignore error
	if err := cmdFlags.Parse(args); err != nil {
		return nil
	}

	// always okay
	return nil
}
