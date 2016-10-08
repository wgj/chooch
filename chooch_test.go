package main

import (
	"fmt"
	"testing"
)

var h host

var filetests = []string{
	"http://localhost:6060/",
	"localhost",
}

func init() {
	h = host{name: "localhost"}
	s := make([]string, 0)
	h.addrs = append(s, "127.0.0.1")
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
	want := resp{
		seq:  0,
		code: 0,
		to:   "127.0.0.1",
		from: "127.0.0.1",
		body: fmt.Sprintf("Ranger-Chooch-%s", h.name),
	}
	h.ping()
	if h.resps[0].seq != want.seq {
		t.Errorf("h.resps[0].seq got %d, want %d", h.resps[0].seq, want.seq)
	}
	if h.resps[0].code != want.code {
		t.Errorf("h.resps[0].code got %d, want %d", h.resps[0].code, want.code)
	}
	if h.resps[0].to != want.to {
		t.Errorf("h.resps[0].to got %s, want %s", h.resps[0].to, want.to)
	}
	if h.resps[0].from != want.from {
		t.Errorf("h.resps[0].from got %s, want %s", h.resps[0].from, want.from)
	}
	if h.resps[0].body != want.body {
		t.Errorf("h.resps[0].body got %s, want %s", h.resps[0].body, want.body)
	}
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

func TestUnpackUrls(t *testing.T) {
	h.name = "http://localhost/path"
	want := host{protocol: "http", name: "localhost", endpoint: "/path"}
	err := h.unpackUrls()
	if err != nil {
		t.Error(err)
	}
	if h.protocol != want.protocol {
		t.Errorf("h.protocol got %s, want %s", h.protocol, want.protocol)
	}
	if h.name != want.name {
		t.Errorf("h.name got %s, want %s", h.name, want.name)
	}
	if h.endpoint != want.endpoint {
		t.Errorf("h.endpoint got %s, want %s", h.endpoint, want.endpoint)
	}
}
