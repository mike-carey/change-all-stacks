package data

import (
	"fmt"
	"strings"
)

const (
	DefaultFormat = "Foundation: %{foundation}, Org: %{org}, Space: %{space}, App: %{app}, LastUpload: %{latestUpload}, LatestAuthor: %{latestAuthor}"
)

func Tsprintf(format string, params map[string]interface{}) string {
	for key, val := range params {
		format = strings.Replace(format, "%{"+key+"}", fmt.Sprintf("%s", val), -1)
	}

	return fmt.Sprintf(format)
}

func TDataMap(foundation string, org string, space string, app string, latestUpload string, latestAuthor string) map[string]interface{} {
	return map[string]interface{}{
		"foundation": foundation,
		"org": org,
		"space": space,
		"app": app,
		"latestUpload": latestUpload,
		"latestAuthor": latestAuthor,
	}
}

func FormatData(formatter Formatter, data Data) (string, error) {
	d := make([]string, len(data))
	for i, entry := range data {
		s, e := formatter.Format(entry)
		if e != nil {
			return "", e
		}

		d[i] = s
	}

	return strings.Join(d, "\n"), nil
}

//go:generate counterfeiter -o fakes/fake_formatter.go Formatter
type Formatter interface {
	Format(entry DataEntry) (string, error)
}

type formatter struct {
	FormatString string
}

func NewFormatter(format string) Formatter {
	return &formatter{
		FormatString: format,
	}
}

func NewDefaultFormatter() Formatter {
	return NewFormatter(DefaultFormat)
}

func (f *formatter) Format(entry DataEntry) (string, error) {
	return Tsprintf(f.FormatString, TDataMap(entry.Foundation, entry.Org.Name, entry.Space.Name, entry.App.Name, "00/00/00 00:00:00", entry.LatestAuthor.Username)), nil
}
