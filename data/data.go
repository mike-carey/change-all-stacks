package data

import (
	"fmt"
	"strings"
	"github.com/cloudfoundry-community/go-cfclient"
)

type DataEntry struct {
	Foundation string
	Org cfclient.Org
	Space cfclient.Space
	App cfclient.App
	LatestUpload string
	LatestAuthor cfclient.User
}
type Data []DataEntry

func (d Data) String() string {
	s := make([]string, len(d))
	for i, entry := range d {
		s[i] = entry.String()
	}

	return strings.Join(s, "\n")
}

func NewDataEntry(foundation string, org cfclient.Org, space cfclient.Space, app cfclient.App, latestUpload string, latestAuthor cfclient.User) *DataEntry {
	return &DataEntry{
		Foundation: foundation,
		Org: org,
		Space: space,
		App: app,
		LatestUpload: latestUpload,
		LatestAuthor: latestAuthor,
	}
}

func (d *DataEntry) String() string {
	return fmt.Sprintf("{ foundation: %s, org: %s, space: %s, app: %s, latestUpload: %s, latestAuthor: %s }", d.Foundation, d.Org.Name, d.Space.Name, d.App.Name, d.LatestUpload, d.LatestAuthor)
}
