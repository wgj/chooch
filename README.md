# chooch
![go report card](https://goreportcard.com/badge/github.com/wgjohnson/chooch)

Ranger Chooch snitches on latency so you don't have to.

`chooch` continuously sends out a ping to a list of hosts, and stores results in a Google Sheets spread sheet.

## Usage
```bash
chooch -f hosts
```
where `hosts` is a list of hostnames to monitor latency.
 
If an entry in `hosts` is a hostname, `chooch` will send ICMP `Echo`. 

If an entry is a URL, `chooch` will send an HTTP `GET`.