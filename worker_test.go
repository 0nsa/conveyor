package conveyor

import (
	"errors"
	"io/ioutil"
	"testing"

	"golang.org/x/net/context"

	"github.com/remind101/conveyor/builder"
	"github.com/remind101/conveyor/logs"
	"github.com/stretchr/testify/assert"
)

func TestWorker(t *testing.T) {
	b := new(mockBuilder)
	q := make(chan BuildRequest, 1)
	w := &Worker{
		Builder:       b,
		Logger:        logs.Discard,
		buildRequests: q,
	}

	done := make(chan struct{})
	go func() {
		w.Start()
		close(done)
	}()

	b.On("Build", ioutil.Discard, builder.BuildOptions{}).Return("", nil)

	q <- BuildRequest{
		Ctx:          context.Background(),
		BuildOptions: builder.BuildOptions{},
	}
	close(q)

	<-done
}

func TestWorker_Shutdown(t *testing.T) {
	b := new(mockBuilder)
	q := make(chan BuildRequest, 1)
	w := &Worker{
		Builder:       b,
		Logger:        logs.Discard,
		buildRequests: q,
		shutdown:      make(chan struct{}),
		done:          make(chan error),
	}

	done := make(chan struct{})
	go func() {
		w.Start()
		close(done)
	}()

	err := w.Shutdown()

	<-done

	assert.NoError(t, err)
}

func TestWorker_Shutdown_Cancel(t *testing.T) {
	b := new(mockCancelBuilder)
	q := make(chan BuildRequest, 1)
	w := &Worker{
		Builder:       b,
		Logger:        logs.Discard,
		buildRequests: q,
		shutdown:      make(chan struct{}),
		done:          make(chan error),
	}

	done := make(chan struct{})
	go func() {
		w.Start()
		close(done)
	}()

	b.On("Cancel").Return(nil)
	err := w.Shutdown()

	<-done

	assert.NoError(t, err)
}

func TestWorker_Shutdown_Cancel_Error(t *testing.T) {
	b := new(mockCancelBuilder)
	q := make(chan BuildRequest, 1)
	w := &Worker{
		Builder:       b,
		Logger:        logs.Discard,
		buildRequests: q,
		shutdown:      make(chan struct{}),
		done:          make(chan error),
	}

	done := make(chan struct{})
	go func() {
		w.Start()
		close(done)
	}()

	boom := errors.New("Failed to cancel")
	b.On("Cancel").Return(boom)
	err := w.Shutdown()

	<-done

	assert.Equal(t, boom, err)
}
