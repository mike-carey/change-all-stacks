package commands

type ChangeOptions struct {
	Plugin  string `short:"p" long:"plugin" description:"The path to the stack-auditor plugin"`
	Interactive bool `short:"i" long:"interactive" description:"Print the dry run before apply"`

	Stacks struct {
		From string
		To string
	} `positional-args:"yes" required:"yes"`
}

type ChangeCommand struct {}

func (c *ChangeCommand) Execute([]string) error {
	// Grab the global manager
	// manager.QueryServices()
	return nil


	// Grab all apps that meet our criteria

}
