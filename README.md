# Drdos Framework

Drdos Framework is my outcome of drdos learning. I try to write this like a real framework.
The tools has been completed. I'll release it later. In fact, it's easy to write this tool. Just forge UDP raw pack and insert the payload.

## Declaration

**This tool is for learning only. Not for illegal use.**

## Prepare

1. Linux server
2. Golang env

## Usage

1. `go get github.com/google/gopacket`
2. Just run `go run main.go --help`

### Some help

`--type` should be one of `dns`,`mem`,`ntp`,`snmp`,`ssdp`,`portmap`,`ldap`
`-o` the outputfile is saved in `/data/results`

## Intruduction

Drdos Framework can check the vulner of drdos iplist. Also it's got the attack mode to start drdos attack.

## Protocol Supported

|Port|Protocol|
|:-:|:-:|
|53|dns|
|111|portmap|
|123|ntp|
|161|snmp|
|389|ldap|
|1900|ssdp|
|11211|memcache|

## Update plan

It will be deleted once I achieve

1. HTTP API
2. SHODAN | FOFA API
3. ~~Blacklist of attack~~
4. Improve response check
5. Support more protocol
6. Use sqlite to save data
