package ephemeral

import (
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestServer(t *testing.T) {
	s := New()

	s.HandleFunc("/", func(s *Server, w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test success\n"))
		s.Stop("foobar")
	})

	var wg sync.WaitGroup
	var stopMsg string

	wg.Add(1)
	go func() {
		defer wg.Done()

		ret, err := s.Listen(":8124")
		require.Nil(t, err)

		stopMsg = ret.(string)
	}()

	time.Sleep(time.Second)

	resp, err := http.Get("http://127.0.0.1:8124/")
	require.Nil(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	require.Nil(t, err)

	require.Equal(t, "test success\n", string(body))

	wg.Wait()

	require.Equal(t, "foobar", stopMsg)
}
