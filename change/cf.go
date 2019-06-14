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

type ChangerWrapper interface {
	ChangeStack(app string, stack string) (string, error)
	GetOrg() (*cfclient.Org, error)
	GetSpace() (*cfclient.Space, error)
}

type changerWrapper struct {
	Changer Changer
	Inquisitor query.Inquisitor
	SpaceGuid string
}

func NewChangerWrapper(ch Changer, inquisitor query.Inquisitor, spaceGuid string) ChangerWrapper {
	return &changerWrapper {
		Changer: ch,
		Inquisitor: inquisitor,
		SpaceGuid: spaceGuid,
	}
}

func (c *changerWrapper) GetOrg() (*cfclient.Org, error) {
	s, e := c.GetSpace()
	if e != nil {
		return &cfclient.Org{}, e
	}

	return c.Inquisitor.GetOrgByGuid(s.OrganizationGuid)
}

func (c *changerWrapper) GetSpace() (*cfclient.Space, error) {
	return c.Inquisitor.GetSpaceByGuid(c.SpaceGuid)
}

func (c *changerWrapper) ChangeStack(app string, stack string) (string, error) {
	return c.Changer.ChangeStack(app, stack)
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
