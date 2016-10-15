# chooch
![Build Status](https://travis-ci.org/wgjohnson/chooch.svg?branch=master) [![Coverage Status](https://coveralls.io/repos/github/wgjohnson/chooch/badge.svg?branch=master)](https://coveralls.io/github/wgjohnson/chooch?branch=master) 
[![Code Climate](https://codeclimate.com/github/wgjohnson/chooch/badges/gpa.svg)](https://codeclimate.com/github/wgjohnson/chooch) ![go report card](https://goreportcard.com/badge/github.com/wgjohnson/chooch)

Ranger Chooch snitches on latency so you don't have to.

`chooch` continuously sends out a ping to a list of hosts, and stores results in a Google Sheets spread sheet.

## Usage
```bash
chooch -f hosts
```
Where `hosts` is a list of hostnames to monitor latency. If an entry in `hosts` is a hostname, `chooch` will send ICMP `Echo`.  If an entry is a URL, `chooch` will send an HTTP `GET`.

## Configuration
Running on Linux requires kernel edits to allow unprivileged ICMP requests. [1](https://godoc.org/golang.org/x/net/icmp#ListenPacket)[2](https://sturmflut.github.io/linux/ubuntu/2015/01/17/unprivileged-icmp-sockets-on-linux/)[3](http://man7.org/linux/man-pages/man7/icmp.7.html)
```bash
sudo /sbin/sysctl -w net.ipv4.ping_group_range="0 2147483647"
```