package change

type Options struct {
	Config  string `short:"c" long:"config" description:"The configuration file to load" default:"cf.json"`
	DryRun  bool   `short:"d" long:"dry-run" description:"Does not actually do the stack change, but instead prints what it would do"`
	Verbose bool   `short:"v" long:"verbose" description:"Prints more output"`

	Stacks struct {
		From string
		To string
	} `positional-args:"yes" required:"yes"`
}

func Go(opts *Options) error {
	return NewRunner(opts.Config, opts.Verbose, opts.DryRun, opts.Stacks.From, opts.Stacks.To).Run()
}
