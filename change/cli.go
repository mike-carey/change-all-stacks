package change

type Options struct {
	Config  string `short:"c" long:"config" description:"The configuration file to load" default:"cf.json"`
	DryRun  bool   `short:"d" long:"dry-run" description:"Does not actually do the stack change, but instead prints what it would do"`
	Verbose bool   `short:"v" long:"verbose" description:"Prints more output"`
	Plugin  string `short:"p" long:"plugin" description:"The path to the stack-auditor plugin"`
	Interactive bool `short:"i" long:"interactive" description:"Print the dry run before apply"`
	Threads int `short:"t" long:"threads" description:"The number of threads to run" default:"10"`
	Orgs []string `short:"o" long:"org" description:"Org names to target"`

	Stacks struct {
		From string
		To string
	} `positional-args:"yes" required:"yes"`
}

func Go(opts *Options) error {
	return NewDefaultManager(opts).Go()
}
