package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"runtime"
	"strings"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

type host struct {
	name     string
	protocol string
	endpoint string
	addrs    []string
	resps    []resp
}

type resp struct {
	id   int
	seq  int
	code int
	sent time.Time
	recv time.Time
	dur  time.Duration
	to   string
	from string
	body string
}

var hosts []host

// readHosts reads in hosts from filename and populates hosts []host.
func readHosts() {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		hosts = append(hosts, host{name: scanner.Text()})
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

// unpackURLs parses h.name{}'s for urls, and sets h.protocol h.endpoint.
func (h *host) unpackUrls() error {
	var isUrl bool
	// loop over host name
	for i := 0; i < len(h.name)-1; i++ {
		if string(h.name[i]) == ":" {
			isUrl = true
			break
		}
	}
	if isUrl {
		url, err := url.Parse(h.name)
		h.protocol = url.Scheme
		h.endpoint = url.Path
		h.name = url.Host
		return err
	}
	return nil
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

func (h *host) ping() {
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

	s := fmt.Sprintf("Ranger-Chooch-%s", h.name)
	var to net.IP
	var from string
	var sent, recv time.Time

	// TODO: icmp echo forever
	for i := 0; i < 1; i++ {
		wm := icmp.Message{
			Type: ipv4.ICMPTypeEcho, Code: 0,
			/* Using ID was from the example from golang.org, and /sbin/ping uses its PID as ICMP's ID.
			 * Maybe there's a reason for this, and we should keep ID as our PID.
			 */
			Body: &icmp.Echo{
				ID: os.Getpid() & 0xffff, Seq: i & 0xffff,
				Data: []byte(s),
			},
		}
		wb, err := wm.Marshal(nil)
		if err != nil {
			log.Fatal(err)
		}

		// Cowardly taking the first address.
		to = net.ParseIP(h.addrs[0])

		if _, err := c.WriteTo(wb, &net.UDPAddr{IP: to}); err != nil {
			log.Fatal(err)
		}
		sent = time.Now()
	}

	rb := make([]byte, 1500)
	// TODO: icmp echoreply forever
	for i := 0; i < 1; i++ {
		n, peer, err := c.ReadFrom(rb)
		if err != nil {
			log.Fatal(err)
		}
		recv = time.Now()
		rm, err := icmp.ParseMessage(1, rb[:n])
		if err != nil {
			log.Fatal(err)
		}

		from, _, _ = net.SplitHostPort(peer.String())
		if h.addrs[0] != from {
			log.Printf("got echo reply from %s; want %s", from, to)
		}
		if rm.Type != ipv4.ICMPTypeEchoReply {
			log.Printf("received something other than ping: %d", rm.Type)
			continue
		}

		body := string(rm.Body.(*icmp.Echo).Data)
		if s != body {
			log.Printf("got echo reply body %s; want %s", body, s)
		}
		seq := rm.Body.(*icmp.Echo).Seq
		id := rm.Body.(*icmp.Echo).ID
		dur := sent.Sub(recv)
		h.addResp(id, seq, rm.Code, sent, recv, dur, to.String(), from, body)
		time.Sleep(time.Second)

	}
}

func (h *host) addResp(id, seq, code int, sent, recv time.Time, dur time.Duration, to, from, body string) {
	r := resp{
		id:   id,
		seq:  seq,
		code: code,
		sent: sent,
		recv: recv,
		dur:  dur,
		to:   to,
		from: from,
		body: body,
	}
	h.resps = append(h.resps, r)
}

func main() {
	// if an entry is a url, send a GET request
	// if an entry is a hostname, send an ICMP ping
}
