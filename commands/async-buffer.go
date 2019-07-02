package commands
import (
	"bytes"
	"sync"
)

// Buffer is a goroutine safe bytes.Buffer
type AsyncBuffer struct {
	buffer bytes.Buffer
	mutex  sync.Mutex
}

func NewAsyncBuffer(b []byte) *AsyncBuffer {
	return &AsyncBuffer{
		buffer: *bytes.NewBuffer(b),
		mutex: sync.Mutex{},
	}
}

// Write appends the contents of p to the buffer, growing the buffer as needed. It returns
// the number of bytes written.
func (s *AsyncBuffer) Write(p []byte) (n int, err error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.buffer.Write(p)
}

// String returns the contents of the unread portion of the buffer
// as a string.  If the Buffer is a nil pointer, it returns "<nil>".
func (s *AsyncBuffer) String() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.buffer.String()
}

func (s *AsyncBuffer) CopyFrom(b *bytes.Buffer) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	b.WriteTo(&s.buffer)
}

func (s *AsyncBuffer) Buffer() bytes.Buffer {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.buffer
}
