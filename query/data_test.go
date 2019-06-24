package query_test

import (
	"fmt"
	"io/ioutil"
	"encoding/json"

	cfclient "github.com/cloudfoundry-community/go-cfclient"

	fakes "github.com/mike-carey/change-all-stacks/query/fakes"
)

type Storage struct {
	Apps       []cfclient.App       `json:"apps"`
	Spaces     []cfclient.Space     `json:"spaces"`
	Orgs       []cfclient.Org       `json:"orgs"`
	Buildpacks []cfclient.Buildpack `json:"buildpacks"`
	Stacks     []cfclient.Stack     `json:"stacks"`
}

func GetStorage() *Storage {
	s := Storage{}

	b, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &s)
	if err != nil {
		panic(err)
	}

	return &s
}

func GetFakeClient(storage *Storage) *fakes.FakeCFClient {
	fakeClient := new(fakes.FakeCFClient)

	fakeClient.ListAppsReturns(storage.Apps, nil)
	fakeClient.GetAppByGuidStub = func(guid string) (cfclient.App, error) {
		for _, app := range storage.Apps {
			if app.Guid == guid {
				return app, nil
			}
		}

		return cfclient.App{}, fmt.Errorf("Cannot find app by guid: %s", guid)
	}

	fakeClient.ListSpacesReturns(storage.Spaces, nil)
	fakeClient.GetSpaceByGuidStub = func(guid string) (cfclient.Space, error) {
		for _, space := range storage.Spaces {
			if space.Guid == guid {
				return space, nil
			}
		}

		return cfclient.Space{}, fmt.Errorf("Cannot find space by guid: %s", guid)
	}

	fakeClient.ListOrgsReturns(storage.Orgs, nil)
	fakeClient.GetOrgByGuidStub = func(guid string) (cfclient.Org, error) {
		for _, org := range storage.Orgs {
			if org.Guid == guid {
				return org, nil
			}
		}

		return cfclient.Org{}, fmt.Errorf("Cannot find org by guid: %s", guid)
	}

	fakeClient.ListBuildpacksReturns(storage.Buildpacks, nil)
	fakeClient.GetBuildpackByGuidStub = func(guid string) (cfclient.Buildpack, error) {
		for _, buildpack := range storage.Buildpacks {
			if buildpack.Guid == guid {
				return buildpack, nil
			}
		}

		return cfclient.Buildpack{}, fmt.Errorf("Cannot find buildpack by guid: %s", guid)
	}

	fakeClient.ListStacksReturns(storage.Stacks, nil)

	return fakeClient
}
