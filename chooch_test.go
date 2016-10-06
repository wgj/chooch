package main

import "testing"

var filetests = []string{
	"http://localhost:6060/",
	"localhost",
}

func TestReadHosts(t *testing.T) {
	readHosts()
	for i, f := range filetests {
		if hosts[i].name != f {
			t.Errorf("hosts[%d]: got host{name: %s}, want host{name: %s}", i, hosts[i].name, f)
		}

	}
}
