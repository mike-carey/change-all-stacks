package commands

import (
	"sync"

	"github.com/mike-carey/change-all-stacks/services"
)

var manager services.Manager
var once sync.Once

func GetManager() services.Manager {
	if manager == nil {
		panic("Manager hasn't been initialized!")
	}
	return manager
}

func TranslateManagerOptions(mOpts *ManagerOptions) *services.ManagerOptions {
	return &services.ManagerOptions{
		Config: mOpts.Config,
		DryRun: mOpts.DryRun,
		Threads: mOpts.Threads,
		Foundations: mOpts.Foundations,
		Orgs: mOpts.Orgs,
	}
}

type ManagerOptions struct {
	Config  string `short:"c" long:"config" description:"The configuration file to load" default:"cf.json"`
	DryRun  bool   `short:"d" long:"dry-run" description:"Does not actually do the stack change, but instead prints what it would do"`
	Threads int `short:"t" long:"threads" description:"The number of threads to run" default:"10"`
	Foundations []string `short:"f" long:"foundation" description:"Limit the foundations that are targeted from the config file"`
	Orgs []string `short:"o" long:"org" description:"Limit the orgs that are targeted from the foundations"`
}

func getDefaultManagerOptions() *ManagerOptions {
	return &ManagerOptions{}
}
