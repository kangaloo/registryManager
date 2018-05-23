package docker

import "testing"

func TestNew(t *testing.T) {
	_, err := New("/etc/docker/daemon.json")

	if err != nil {
		t.Errorf("%v", err)
	}
}
