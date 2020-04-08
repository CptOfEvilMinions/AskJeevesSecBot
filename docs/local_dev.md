# Local development

### Init project
1. `cd AskJeeves/`
1. `go mod init github.com/CptOfEvilMinions/AskJeevesSecBot`

### Build project
* [Installing librdkafka](https://github.com/confluentinc/confluent-kafka-go#getting-started)

#### macOS
1. `brew install librdkafka pkg-config`
1. `go build main.go`

#### Ubuntu
1. `apt-get install librdkafka-dev -y`
1. `go build main.go`
