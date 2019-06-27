package query

import (

)

type QueryStackOptions struct {
	Stacks []string `positional-args:"yes" required:"yes" description:"The stacks to "`
}

type QueryStackCommand struct {}

func (c *QueryStackCommand) Execute([]args) error {
	
}
