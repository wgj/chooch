package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

type host struct {
	name     string
	protocol string
	endpoint string
}

var hosts []host

// readHosts reads in hosts from a file and populates hosts []host.
func readHosts() {
	f, err := os.Open(filename)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		hosts = append(hosts, host{name: scanner.Text()})
		// TODO: parse for urls, set host.protocol and host.endpoint
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

func main() {
	// if an entry is a url, send a GET request
	// if an entry is a hostname, send an ICMP ping
	// TODO: host method for GET
	// TODO: host method for ICMP
	// TODO: figure out how to represent responses.
	// TODO: store responses in google sheets.
}
