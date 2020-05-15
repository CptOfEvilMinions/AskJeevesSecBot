# Askjeeves

## Local development
### Install Golang on macOS
1. `curl -o golang.pkg https://dl.google.com/go/go1.13.3.darwin-amd64.pkg`
1. `sudo open golang.pkg`
1. Follow instructions to install GOLANG
1. `echo 'export GOROOT=/usr/local/go' >> ~/.zshrc`
1. `echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> ~/.zshrc`
1. `source ~/.zshrc`
1. `go version`

### Init project
1. `cd AskJeeves/`
1. `go mod init github.com/CptOfEvilMinions/AskJeevesSecBot`

#### macOS
1. `brew install librdkafka pkg-config`
1. `go build -o AskJeeves main.go`

#### Ubuntu
1. `apt-get install librdkafka-dev -y`
1. `go build -o AskJeeves main.go`

### Spin up stack
1. `docker-compose -f docker-compose-dev.yml build`
1. `docker-compose -f docker-compose-dev.yml up`

### Run AskJeeves
1. ``

## References
* [How To Install Go 1.13 on MacOS](https://tecadmin.net/install-go-on-macos/)
* [Installing librdkafka](https://github.com/confluentinc/confluent-kafka-go#getting-started)