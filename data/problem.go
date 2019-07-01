package data

import (
	"fmt"

	"github.com/cloudfoundry-community/go-cfclient"
)

type Reason interface {
	GetReason(app cfclient.App) string
}

func InvalidBuildpack(buildpack string) Reason {
	return &invalidBuildpack {
		buildpack: buildpack,
	}
}

type invalidBuildpack struct {
	buildpack string
}

func (r *invalidBuildpack) GetReason(app cfclient.App) string {
	return fmt.Sprintf("%s is using an invalid buildpack: %s", app.Name, r.buildpack)
}

func MissingDroplet() Reason {
	return &missingDroplet{}
}

type missingDroplet struct {}

func (r *missingDroplet) GetReason(app cfclient.App) string {
	return fmt.Sprintf("%s has not been staged in awhile and will need to rebuild its droplet", app.Name)
}

type ProblemData struct {
	App cfclient.App
	Reason Reason
}

func (p *ProblemData) GetReason() string {
	return p.Reason.GetReason(p.App)
}

type ProblemSet []ProblemData
