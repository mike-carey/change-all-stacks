package services

import (
	"fmt"
	"io/ioutil"
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/mike-carey/change-all-stacks/data"
	"github.com/mike-carey/change-all-stacks/logger"
	"github.com/mike-carey/change-all-stacks/query"

	"github.com/cloudfoundry-community/go-cfclient"
)

const (
	DropletNotFoundCode = 10010
)

//go:generate counterfeiter -o fakes/fake_problem_service.go ProblemService
type ProblemService interface {
	FindProblems(apps []cfclient.App, fromStack string, toStack string) (data.ProblemSet, error)
}

func NewProblemService(inquisitor query.Inquisitor) ProblemService {
	return &problemService{
		inquisitor: inquisitor,
	}
}

type problemService struct {
	inquisitor query.Inquisitor
}

func (s *problemService) FindProblems(apps []cfclient.App, fromStack string, toStack string) (data.ProblemSet, error) {
	set := make(data.ProblemSet, 0)

	for _, app := range apps {
		// Check that buildpack it has is valid
		bp, bs, err := s.getBuildpackForApp(app, fromStack, toStack)
		if err != nil {
			return nil, err
		}

		if bp == nil {
			set = append(set, data.ProblemData{
				App: app,
				Reason: data.InvalidBuildpack(bs),
			})
		}

		// Check that the current droplet is available
		droplet, err := s.GetCurrentDroplet(app)
		if err != nil {
			return nil, err
		}

		if droplet == "" {
			set = append(set, data.ProblemData{
				App: app,
				Reason: data.MissingDroplet(),
			})
		}
	}

	return set, nil
}

func (s *problemService) getBuildpackForApp(app cfclient.App, fromStack string, toStack string) (*cfclient.Buildpack, string, error) {
	bps, err := s.inquisitor.GetAllBuildpacks()
	if err != nil {
		return nil, "", err
	}

	if app.Buildpack != "" {
		logger.Debugf("App(%s) explicitly specified a buildpack(%s)", app.Name, app.Buildpack)

		for _, b := range bps {
			if b.Name == app.Buildpack {
				logger.Debugf("Found the explicit buildpack(%s)", app.Buildpack)
				if b.Stack == fromStack {
					logger.Debugf("The buildpack is using the older stack(%s)", fromStack)
					continue
				}

				if b.Stack == toStack {
					logger.Debugf("The buildpack is using the newer stack(%s)", fromStack)
					return &b, app.Buildpack, nil
				}
			}
		}

		logger.Debugf("Did not find the proper buildpack(%s) with stack(%s) for app(%s)", app.Buildpack, toStack, app.Name)
		return nil, app.Buildpack, nil
	}

	if app.DetectedBuildpackGuid != "" {
		logger.Debugf("App(%s)'s buildpack was detected as a guid %s", app.Name, app.DetectedBuildpackGuid)
		buildpack, err := s.inquisitor.GetBuildpackByGuid(app.DetectedBuildpackGuid)
		if err != nil {
			return nil, "", err
		}

		if buildpack.Stack == fromStack {
			logger.Debugf("App(%s)'s buildpack(%s) is using %s, checking if there is %s version of the buildpack", app.Name, buildpack.Name, fromStack, toStack)

			for _, b := range bps {
				if b.Guid == buildpack.Guid {
					continue
				}

				if b.Name == buildpack.Name && b.Stack == toStack {
					logger.Debugf("Found a %s buildpack with stack(%s)", buildpack.Name, toStack)
					return &buildpack, buildpack.Name, nil
				}
			}

			return nil, buildpack.Name, nil
		}

		return &buildpack, buildpack.Name, nil
	}

	if app.DetectedBuildpack != "" {
		logger.Debugf("App(%s)'s buildpack was detected as %s", app.Name, app.DetectedBuildpack)
		var buildpack cfclient.Buildpack
		for _, b := range bps {
			if app.DetectedBuildpack == b.Name {
				if b.Stack == fromStack {
					logger.Debugf("Found the detected buildpack, but it is %s", fromStack)
					continue
				}

				if b.Stack == toStack {
					buildpack = b
					break
				}
			}
		}

		_b := cfclient.Buildpack{}
		if buildpack == _b {
			return nil, app.DetectedBuildpack, nil
		}

		return &buildpack, app.DetectedBuildpack, nil
	}

	logger.Warningf("Buildpack for app(%s) could not be determined", app.Name)

	return nil, "", nil
}

type DropletResponse struct {
	Errors []struct {
		Detail string `json:"detail"`
		Title string `json:"title"`
		Code int `json:"code"`
	} `json:"errors"`
	Data struct {
		Guid string `json:"guid"`
	} `json:"data"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Related struct {
			Href string `json:"href"`
		} `json:"related"`
	} `json:"links"`
}

func (s *problemService) GetCurrentDroplet(app cfclient.App) (string, error) {
	c := s.inquisitor.Client()
	req := c.NewRequest("GET", fmt.Sprintf("/v3/apps/%s/relationships/current_droplet", app.Guid))
	resp, err := c.DoRequest(req)
	if err != nil {
		if e, ok := err.(cfclient.CloudFoundryError); ok {
			if e.Code == DropletNotFoundCode {
				return "", nil
			}
		}

		return "", errors.Wrap(err, "Error requesting current_droplet")
	}
	defer resp.Body.Close()
	b, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return "", e
	}

	var res DropletResponse
	err = json.Unmarshal(b, &res)
	if err != nil {
		return "", errors.Wrap(err, "Error unmarshaling response %v")
	}

	if len(res.Errors) > 0 {
		// Assume there is one for now
		for _, e := range res.Errors {
			if e.Code == DropletNotFoundCode {
				return "", nil
			}

			return "", errors.New(e.Detail)
		}
	}

	return res.Data.Guid, nil
}
