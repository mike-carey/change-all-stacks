package data

import (
	"fmt"

	"github.com/cloudfoundry-community/go-cfclient"
)

const (
	InvalidBuildpackId = "InvalidBuildpack"
	MissingDropletId = "MissingDroplet"
)

type Reason interface {
	GetReason(app cfclient.App) string
	GetId() string
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

func (r *invalidBuildpack) GetId() string {
	return InvalidBuildpackId
}

func MissingDroplet() Reason {
	return &missingDroplet{}
}

type missingDroplet struct {}

func (r *missingDroplet) GetReason(app cfclient.App) string {
	return fmt.Sprintf("%s has not been staged in awhile and will need to rebuild its droplet", app.Name)
}

func (r *missingDroplet) GetId() string {
	return MissingDropletId
}

type ProblemData struct {
	Foundation string
	Org cfclient.Org
	Space cfclient.Space
	App cfclient.App
	LatestAuthor cfclient.User
	LatestUpload string
	Reason Reason
}

func NewProblemData(foundation string, org cfclient.Org, space cfclient.Space, app cfclient.App, latestUpload string, latestAuthor cfclient.User, reason Reason) *ProblemData {
	return &ProblemData{
		Foundation: foundation,
		Org: org,
		Space: space,
		App: app,
		LatestUpload: latestUpload,
		LatestAuthor: latestAuthor,
		Reason: reason,
	}
}


func (p *ProblemData) GetReason() string {
	return p.Reason.GetReason(p.App)
}

type ProblemSet []ProblemData
