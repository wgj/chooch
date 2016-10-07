package main

import (
	"bufio"
	"flag"
	"log"
	"net"
	"os"
	"runtime"
	"strings"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type host struct {
	name     string
	protocol string
	endpoint string
	addrs    []string
}

var hosts []host

// readHosts reads in hosts from a file and populates hosts []host.
func readHosts() {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		hosts = append(hosts, host{name: scanner.Text()})
		// TODO: parse for urls, set host.protocol and host.endpoint. net/url.Parse seems like a good fit.
	}
	if err := scanner.Err(); err != nil {
		log.Printf("error reading hosts from %s:%s\n", filename, err)
	}
}

var filename string

func init() {
	// Check for '-f' flag for host file
	flag.StringVar(&filename, "f", "hosts", "File with hosts and urls to check.")
	flag.Parse()
}

func (h *host) htoi() error {
	if len(h.addrs) == 0 {
		addrs, err := net.LookupHost(h.name)
		for _, i := range addrs {
			// Skip if address is ipv6.
			if strings.Contains(i, ":") {
				continue
			}
			h.addrs = append(h.addrs, i)
		}
		return err
	}
	return nil
}

func (h host) ping() {
	switch runtime.GOOS {
	case "darwin":
	case "linux":
		log.Println("you may need to adjust the net.ipv4.ping_group_range kernel state")
	default:
		log.Println("not supported on", runtime.GOOS)
		return
	}

	c, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("Ranger-Chooch"),
		},
	}
	wb, err := wm.Marshal(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Cowardly taking the first address.
	if _, err := c.WriteTo(wb, &net.UDPAddr{IP: net.ParseIP(h.addrs[0])}); err != nil {
		log.Fatal(err)
	}

	rb := make([]byte, 1500)
	n, peer, err := c.ReadFrom(rb)
	if err != nil {
		log.Fatal(err)
	}
	rm, err := icmp.ParseMessage(1, rb[:n])
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("got reflection from %v", peer)
	log.Printf("got %+v", rm)
	log.Printf("got %s", rm.Body)
	log.Printf("rm.Body type: %T", rm.Body)
	log.Printf("os.Getpid(): %d", os.Getpid()&0xffff)

}

func main() {
	// if an entry is a url, send a GET request
	// if an entry is a hostname, send an ICMP ping
	// TODO: host method for GET
	// TODO: host method for ICMP
	// TODO: figure out how to represent responses.
	// TODO: store responses in google sheets.
	// TODO: cache writes to google sheets if network is unavailable.
	// TODO: rewrite host request methods as goroutines.
}
