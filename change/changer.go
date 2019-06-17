package change

import (
	"io"

	"github.com/cloudfoundry/stack-auditor/cf"
	"github.com/cloudfoundry/stack-auditor/changer"
	"github.com/cloudfoundry/stack-auditor/utils"

	"code.cloudfoundry.org/cli/plugin"
	plugin_models "code.cloudfoundry.org/cli/plugin/models"

	cfclient "github.com/cloudfoundry-community/go-cfclient"
)

func getSpace(space *cfclient.Space) plugin_models.Space {
	s := plugin_models.Space{}
	s.Guid = space.Guid
	s.Name = space.Name

	return s
}

func NewChanger(cliConnection plugin.CliConnection, space *cfclient.Space) (Changer, error) {
	return &changer.Changer{
		Log: func(w io.Writer, msg string) {
			w.Write([]byte(msg))
		},
		Runner: utils.Command{},
		CF: cf.CF{
			Conn:  cliConnection,
			Space: getSpace(space),
		},
	}, nil
}
