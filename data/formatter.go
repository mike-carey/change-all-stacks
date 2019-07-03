package data

import (
	"fmt"
	"strings"
)

const (
	DataDefaultFormat = "Foundation: %{foundation}, Org: %{org}, Space: %{space}, App: %{app}, LastUpload: %{latestUpload}, LatestAuthor: %{latestAuthor}"

	DataCsvHeader = "Foundation, Org, Space, App, LastUpload, LatestAuthor"
	DataCsvFormat = "%{foundation}, %{org}, %{space}, %{app}, %{latestUpload}, %{latestAuthor}"

	ProblemDataDefaultFormat = "Foundation: %{foundation}, Org: %{org}, Space: %{space}, App: %{app}, LastUpload: %{latesUpload}, LatestAuthor: %{latestAuthor}, Reason: %{reason}"

	ProblemDataCsvHeader = "Foundation, Org, Space, App, LastUpload, LatestAuthor, Reason"
	ProblemDataCsvFormat = "%{foundation}, %{org}, %{space}, %{app}, %{latestUpload}, %{latestAuthor}, %{reason}"
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

func TProblemDataMap(foundation string, org string, space string, app string, latestUpload string, latestAuthor string, reason string) map[string]interface{} {
	return map[string]interface{}{
		"foundation": foundation,
		"org": org,
		"space": space,
		"app": app,
		"latestUpload": latestUpload,
		"latestAuthor": latestAuthor,
		"reason": reason,
	}
}

//go:generate counterfeiter -o fakes/fake_formatter.go Formatter
type Formatter interface {
	FormatData(entry Data) (string, error)
	FormatProblemSet(set ProblemSet) (string, error)
}

type formatter struct {
	DataFormatString string
	ProblemDataFormatString string
}

func NewFormatter(dataFormat string, problemDataFormat string) Formatter {
	return &formatter{
		DataFormatString: dataFormat,
		ProblemDataFormatString: problemDataFormat,
	}
}

func NewDefaultFormatter() Formatter {
	return NewFormatter(DataDefaultFormat, ProblemDataDefaultFormat)
}

func (f *formatter) FormatData(entries Data) (string, error) {
	strs := make([]string, len(entries))
	for i, entry := range entries {
		strs[i] = Tsprintf(f.DataFormatString, TDataMap(entry.Foundation, entry.Org.Name, entry.Space.Name, entry.App.Name, entry.App.UpdatedAt, entry.LatestAuthor.Username))
	}

	return strings.Join(strs, "\n"), nil
}

func (f *formatter) FormatProblemSet(set ProblemSet) (string, error) {
	strs := make([]string, len(set))
	for i, entry := range set {
		strs[i] = Tsprintf(f.ProblemDataFormatString, TProblemDataMap(entry.Foundation, entry.Org.Name, entry.Space.Name, entry.App.Name, entry.App.UpdatedAt, entry.LatestAuthor.Username, entry.Reason.GetId()))
	}

	return strings.Join(strs, "\n"), nil
}

var _ Formatter = &formatter{}

type csvFormatter struct {
	DataHeader string
	DataFormat string
	ProblemDataHeader string
	ProblemDataFormat string
}

func NewCsvFormatter() Formatter {
	return &csvFormatter{
		DataHeader: DataCsvHeader,
		DataFormat: DataCsvFormat,
		ProblemDataHeader: ProblemDataCsvHeader,
		ProblemDataFormat: ProblemDataCsvFormat,
	}
}

func (f *csvFormatter) FormatData(entries Data) (string, error) {
	strs := make([]string, len(entries)+1)
	strs[0] = f.DataHeader
	for i, entry := range entries {
		strs[i+1] = Tsprintf(f.DataFormat, TDataMap(entry.Foundation, entry.Org.Name, entry.Space.Name, entry.App.Name, entry.App.UpdatedAt, entry.LatestAuthor.Username))
	}

	return strings.Join(strs, "\n"), nil
}

func (f *csvFormatter) FormatProblemSet(set ProblemSet) (string, error) {
	strs := make([]string, len(set)+1)
	strs[0] = f.ProblemDataHeader
	for i, entry := range set {
		strs[i+1] = Tsprintf(f.ProblemDataFormat, TProblemDataMap(entry.Foundation, entry.Org.Name, entry.Space.Name, entry.App.Name, entry.App.UpdatedAt, entry.LatestAuthor.Username, entry.Reason.GetId()))
	}

	return strings.Join(strs, "\n"), nil
}

var _ Formatter = &csvFormatter{}
