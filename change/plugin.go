package change

import (
	"fmt"
	"os"
	"strings"
	"strconv"

	"code.cloudfoundry.org/cli/cf/commandregistry"
	"code.cloudfoundry.org/cli/cf/configuration/coreconfig"
	"code.cloudfoundry.org/cli/plugin"
	"code.cloudfoundry.org/cli/plugin/rpc"
	// "code.cloudfoundry.org/cli/cf/terminal"
	"code.cloudfoundry.org/cli/cf/trace"

	netrpc "net/rpc"
)

var Writer = os.Stdout

//go:generate counterfeiter -o fakes/fake_cli_connection.go code.cloudfoundry.org/cli/plugin.CliConnection
type ChangeAllStacksPlugin struct {
	connection plugin.CliConnection
}

func parseVersion(version string) plugin.VersionType {
	vs := strings.Split(version, ".")
	is := make([]int, len(vs))
	for i, v := range vs {
		j, err :=  strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		is[i] = j
	}

	return plugin.VersionType{
		Major: is[0],
		Minor: is[1],
		Build: is[2],
	}
}

func (p *ChangeAllStacksPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "change-all-stacks",
		Version: parseVersion(Version[1:]),
		MinCliVersion: parseVersion("6.3.0"),
		Commands: []plugin.Command{
			plugin.Command {
				Name: "change-all-stacks",
				Alias: "cas",
				HelpText: "Change all stacks",
				UsageDetails: plugin.Usage{
					Usage: "change-all-stacks",
					Options: make(map[string]string, 0),
				},
			},
		},
	}
}

func (p *ChangeAllStacksPlugin) Run(cliConnection plugin.CliConnection, _ []string) {
	p.connection = cliConnection
}

func (p *ChangeAllStacksPlugin) GetConnection() plugin.CliConnection {
	return p.connection
}

func closeConfig(conf coreconfig.Repository) {
	fmt.Println("Closing config")
	conf.Close()
}

func (p *ChangeAllStacksPlugin) withRpcService(do func(rpcService *rpc.CliRpcService) error) error {
	traceLogger := trace.NewLogger(os.Stdout, true)
	deps := commandregistry.NewDependency(Writer, traceLogger, "6000")
	defer closeConfig(deps.Config)

	server := netrpc.NewServer()
	rpcService, err := rpc.NewRpcService(deps.TeePrinter, deps.TeePrinter, deps.Config, deps.RepoLocator, rpc.NewCommandRunner(), deps.Logger, Writer, server)
	if err != nil {
		return err
	}

	err = do(rpcService)

	return err
}

func stopRpcService(rpcService *rpc.CliRpcService) {
	fmt.Println("Stopping rpc service")
	rpcService.Stop()
}

func (p *ChangeAllStacksPlugin) WithConnection(do func(cliConnection plugin.CliConnection) error) error {
	err := p.withRpcService(func (rpcService *rpc.CliRpcService) error {
		defer stopRpcService(rpcService)
		rpcService.Start()

		fmt.Printf("Started rpc service on port %q", rpcService.Port())

		os.Args = []string{os.Args[0], rpcService.Port()}
		plugin.Start(p)

		err := do(p.connection)

		return err
	})

	return err
}

func WithCliConnection(do func(cliConnection plugin.CliConnection) error) error {
	p := &ChangeAllStacksPlugin{}
	e := p.WithConnection(do)

	return e
}
