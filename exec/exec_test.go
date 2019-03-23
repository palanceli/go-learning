package golexec

import (
	"bytes"
	"flag"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/golang/glog"
)

func TestMain(m *testing.M) {
	defer glog.Flush()
	flag.Set("logtostderr", "true")
	flag.Parse()
	ret := m.Run()
	os.Exit(ret)
}

func TestStartAndWait(t *testing.T) {
	cmd := exec.Command("ping", "-c", "10", "127.0.0.1")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		t.Fatalf("FAILED to start ping. err=%v", err)
	}
	cmd.Wait()
	glog.Infof("out: %s", out.String())
}

func TestStartAndKill(t *testing.T) {
	cmd := exec.Command("ping", "127.0.0.1")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		t.Fatalf("FAILED to start ping. err=%v", err)
	}
	defer cmd.Process.Kill()
	time.Sleep(time.Duration(5) * time.Second)
	glog.Infof("out: %s", out.String())
}
