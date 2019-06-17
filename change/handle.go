package change

import (
	"io"
	"os"
	"fmt"
)

//go:generate counterfeiter -o fakes/fake_changer.go handle.go Changer
type Changer interface {
	ChangeStack(app string, stack string) (string, error)
}

type Handler interface {
	Handle(org string, space string, app string, stack string) error
	HandleDryRun(org string, space string, app string, stack string) error
}

//go:generate counterfeiter -o fakes/fake_writer.go io.Writer
type handler struct {
	Changer Changer
	Writer io.Writer
}

func NewHandler(ch Changer, writer io.Writer) *handler {
	return &handler{
		Changer: ch,
		Writer: writer,
	}
}

func NewHandlerWithStdout(ch Changer) *handler {
	return &handler{
		Changer: ch,
		Writer: os.Stdout,
	}
}

func (h *handler) printf(msg string, args ...interface{}) {
	h.Writer.Write([]byte(fmt.Sprintf(msg, args...)))
}

func (h *handler) Handle(org string, space string, app string, stack string) error {
	h.printf("Changing stack in org: %s, space: %s, app: %s, to %s", org, space, app, stack)
	str, err := h.Changer.ChangeStack(app, stack)
	if err == nil {
		h.printf(str)
	} else {
		h.printf("%v", err)
	}

	return err
}

func (h *handler) HandleDryRun(org string, space string, app string, stack string) error {
	h.printf("cf target -o %s -s %s\ncf change-stack %s %s", org, space, app, stack)

	return nil
}
