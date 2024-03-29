package basicserverhttp

import (
	"net/http"
	"os"
	"testing"

	"github.com/sirupsen/logrus"

	clarkezoneLog "github.com/clarkezone/boosted-go/log"
)

// TestMain initizlie all tests
func TestMain(m *testing.M) {
	clarkezoneLog.Init(logrus.DebugLevel)
	code := m.Run()
	os.Exit(code)
}

func Test_webhooklistening(t *testing.T) {
	// wait := make(chan bool)
	wh := BasicServer{}
	mux := http.NewServeMux()
	wh.StartMetrics(8080)
	wh.StartListen(8081, "ss", mux)
	// TODO: how to wait for server async start
	// client := &http.Client{}
	// req, err := http.NewRequestWithContext(context.Background(), "POST", "http://0.0.0.0:8090", nil)
	// if err != nil {
	//	t.Errorf("bad request")
	// }
	// req.Header.Set("X-GitHub-Event", "push")
	// resp, err := client.Do(req)
	// if err != nil {
	//	log.Fatal(err)
	// }
	// <-wait
	err := wh.Shutdown()
	if err != nil {
		t.Errorf("shutdown failed")
	}
	// err = resp.Body.Close()
	// if err != nil {
	//	t.Errorf("close failed")
	//}
}
