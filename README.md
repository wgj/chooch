# chooch
[![Build Status](https://travis-ci.org/wgjohnson/chooch.svg?branch=master)](https://travis-ci.org/wgjohnson/chooch) [![Coverage Status](https://coveralls.io/repos/github/wgjohnson/chooch/badge.svg?branch=master)](https://coveralls.io/github/wgjohnson/chooch?branch=master) 
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

## Background
`chooch` is the result of a conversation between a sysadmin and a data scientist about a spotty network in the middle of the desert.

We share responsibility for IT operations at Burning Man, including network uptime, and we were having a hard time pin pointing on a particular failure case. 
We had a lot of theories about the cause of the issue (unreliable physical media, network latency, application load), but didn't have enough evidence to support any of our claims.
`chooch` is an attempt to solve this problem.
Or more accurately, point us in the right direction. 
`chooch` will give us enough information to know the _where_ and _when_ for detailed inspections. 
Currently, we're dependent on a user's first hand accounts, which may not have enough detail ("it happened on Friday" or "it doesn't work").

`chooch` was suggested for two reasons: 

* We didn't want to introduce large networking monitoring tools designed for enterprises. 
Our problem is complex, but not that complex.
Introducing a behemoth monitoring suite would create more problems than it solves (installation, configuration, maintenance).
* We wanted an excuse to program in [go](https://golang.org/). This was an interesting enough problem, and allowed us to grow beyond sandboxes like [Exercism.io](http://exercism.io/).
In addition to an interesting problem space, writing our own project meant exploring the greater software development life cycle.
Writing our own unit tests, and using [Travis CI](https://travis-ci.org/) lets us exploring professional programming practices.