package main

import "testing"

var filetests = []string{
	"http://localhost:6060/",
	"localhost",
	"www.google.com",
}

func TestReadHosts(t *testing.T) {
	readHosts()
	for i, f := range filetests {
		if hosts[i].name != f {
			t.Errorf("hosts[%d]: got host{name: %s}, want host{name: %s}", i, hosts[i].name, f)
		}
	}
}

func TestPing(t *testing.T) {
	// TODO: Improve this when ping() has output.
	s := make([]string, 0)
	addrs := append(s, "127.0.0.1")
	host{name: "localhost", addrs: addrs}.ping()

	//s := make([]string, 0)
	//addrs := append(s, "127.0.0.1")
	host{name: "localhost", addrs: addrs}.ping()
}

func TestHtoI(t *testing.T) {
	h := host{name: "localhost"}
	want := "127.0.0.1"
	err := h.htoi()
	if err != nil {
		t.Fatal(err)
	}
	got := h.addrs[0]
	if got != want {
		t.Errorf("h.addrs: got %s, want %s", got, want)
	}
}
