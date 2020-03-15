# AskJeevesSecBot

AskJeevesSecBot is golang project that has a rule engine that triggers Slack notifications to users based on pre-defined conditions. For example, a user is located in the US but suddenly there is a VPN login from Russia. This project will will send a Slack notification to the user asking if that was them, if so nothing happens, if not an incident response ticket is generated.

https://manios.org/2018/08/18/logstash-geoip-json-logs-maxmint-geolite-docker
`logger -n 10.150.100.210 -P 1514 --rfc3164 -t 'openvpn' "1.1.1.1:56555 [spartan2194] Peer Connection Initiated with [AF_INET]1.1.1.1:56555"`

## Init project
1. `go mod init github.com/CptOfEvilMinions/AskJeevesSecBot`

## Download GeoIP database
1. Go to https://www.maxmind.com/en/geoip-demo and login
1. Download GeoLite2-City
1. Place GeoLite2-City.mmdb in `app/data/GeoLite2-City.mmdb`

## Spin up Docker stack
1. `docker-compose build`
1. `docker-compose up -d`

## References
* [Rsyslog Expressions](https://www.rsyslog.com/doc/v8-stable/rainerscript/expressions.html)
* [Index JSON Messages over Syslog](https://sematext.com/docs/logs/json-messages-over-syslog/)
* [Rsyslog field()](https://www.rsyslog.com/doc/v8-stable/rainerscript/functions/rs-field.html)
* [23.8. STRUCTURED LOGGING WITH RSYSLOG](https://access.redhat.com/documentation/en-us/red_hat_enterprise_linux/7/html/system_administrators_guide/s1-structured_logging_with_rsyslog)
* [IP to City Lite database](https://db-ip.com/db/download/ip-to-city-lite)
* [Format a Go string without printing?](https://stackoverflow.com/questions/11123865/format-a-go-string-without-printing)
* [Go by Example: SHA1 Hashes](https://gobyexample.com/sha1-hashes)
* [Package md5](https://golang.org/pkg/crypto/md5/)
* [Using Go Modules](https://blog.golang.org/using-go-modules)
* [Slack API in Go](https://github.com/slack-go/slack)
* [Marshal and unMarshal of Struct to JSON in Golang](https://www.restapiexample.com/golang-tutorial/marshal-and-unmarshal-of-struct-data-using-golang/)
* [Assigning null to JSON fields instead of empty strings](https://stackoverflow.com/questions/31048557/assigning-null-to-json-fields-instead-of-empty-strings)
* [How convert a string into json or a struct?](https://forum.golangbridge.org/t/how-convert-a-string-into-json-or-a-struct/3457)
* [Kafka Go Client](https://docs.confluent.io/current/clients/go.html)
* [package maxminddb](https://pkg.go.dev/github.com/oschwald/maxminddb-golang?tab=doc#example-Reader.Lookup-Interface)
* [KSQL Problem while - print <topic> from beginning #2386](https://github.com/confluentinc/ksql/issues/2386)
* []()
* []()
* []()