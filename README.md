# [Go Guide]

## Install Go

- 각 OS와 사양에 맞는 설치 파일 다운로드 및 실행 https://go.dev/dl/

## Setting(in MAC OS)

- 터미널 실행 후 각 shell 종류에 해당하는 파일로 접근(해당 예시는 zsh환경에서 실행)

```zsh
$ vi ~/.zshrc
```

- '.zshrc'에서 환경변수 설정

```vim
# GO lang
# GOPATH: go 언어 사용을 위한 환경 변수
export GOPATH="/Users/user_name/development/golang"
# PATH: go 파일 실행 및 빌드를 위한 환경 변수
export PATH="$PATH:GOPATH/bin"
```

- 변경 된 환경변수 적용

```zsh
$ source ~/.zshrc
```

- 변경 된 환경 설정이 제대로 적용되었는지 확인(go전체 환경 설정 확인)

```zsh
$ go env
```

## Go-Module: go.mod

### Create go.mod

- project root path로 이동

```zsh
$ cd $GOPATH/src/project_name
```

- go module 생성 명령어 실행(go.mod 파일이 없는 상태에서 실행)

```zsh
$ go mod init example.com/project_name
# go: creating new go.mod: module example.com/project_name
```

```go
module example.com/project_name

go 1.20 // go version
```

### Add module to go.mod

```zsh
# Example code
go get -u google.golang.org/grpc
```

## Run Go

1. go파일이(main.go) 있는 directory로 이동
2. $ go build
   -> 해당 디렉토리 이름의 실행 파일 생성
3. $ go run main.go
   -> go 파일 실행

## gRPC & ProtoBuf

### gRPC

```zsh
# Add to module
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

# Install module
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### Protocol Compiler

- Linux

```zsh
apt install -y protobuf-compiler
```

- Mac:

```zsh
brew install protobuf
```

- Check Version

```zsh
protoc —version
```

### Protobuf Plugin

```zsh
# Add to mod
go get -u google.golang.org/protobuf/cmd/protoc-gen-go

# Install module
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

### Create \*.pb.go

```zsh
# Create
protoc -I=. \ --go_out . --go_opt paths=source_relative \ --go-grpc_out . --go-grpc_opt paths=source_relative \ service.proto
```
