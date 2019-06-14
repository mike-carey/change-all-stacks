package change_all_stacks

import (
	"io"

	"github.com/cloudfoundry/stack-auditor/cf"
	"github.com/cloudfoundry/stack-auditor/changer"
	"github.com/cloudfoundry/stack-auditor/utils"

	"code.cloudfoundry.org/cli/plugin"
	plugin_models "code.cloudfoundry.org/cli/plugin/models"

	cfclient "github.com/cloudfoundry-community/go-cfclient"

	"github.com/mike-carey/cfquery/query"
)

const ChangeStackCmd = "change-stack"

//go:generate counterfeiter -o fakes/fake_changer.go cf.go Changer
type Changer interface {
	ChangeStack(app string, stack string) (string, error)
}

func GetCliConnection() plugin.CliConnection {
	cliConnection := plugin.NewCliConnection(ChangeStackCmd)

	return cliConnection
}

func GetSpace(space *cfclient.Space) plugin_models.Space {
	s := plugin_models.Space{}
	s.Guid = space.Guid
	s.Name = space.Name

	return s
}

func GetChanger(i query.Inquisitor, spaceGuid string) (Changer, error) {
	s, e := i.GetSpaceByGuid(spaceGuid)
	if e != nil {
		return nil, e
	}

	return &changer.Changer{
		Log: func(w io.Writer, msg string) {
			w.Write([]byte(msg))
		},
		Runner: utils.Command{},
		CF: cf.CF{
			Conn:  GetCliConnection(),
			Space: GetSpace(s),
		},
	}, nil
}
