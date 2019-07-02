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
		// DryRun: mOpts.DryRun,
		// Threads: mOpts.Threads,
		Foundations: mOpts.Foundations,
		Orgs: mOpts.Orgs,
		PluginPath: mOpts.PluginPath,
	}
}

type ManagerOptions struct {
	Config  string `short:"c" long:"config" description:"The configuration file to load" default:"cf.json"`
	Foundations []string `short:"f" long:"foundation" description:"Limit the foundations that are targeted from the config file"`
	Orgs []string `short:"o" long:"org" description:"Limit the orgs that are targeted from the foundations"`
	PluginPath  string `short:"p" long:"plugin-path" description:"The path to the stack-auditor plugin"`
}

func getDefaultManagerOptions() *ManagerOptions {
	return &ManagerOptions{}
}
