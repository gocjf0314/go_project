### go 설치

url: 'https://go.dev/dl/'
-> 각 OS와 사양에 맞는 설치 파일 다운로드 및 실행

### 환경설정(in MAC OS)

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

### Create Go-Module

- project root path로 이동

```zsh
$ cd $GOPATH/src/project_name
```

- go module 생성 명령어 실행(go.mod 파일이 없는 상태에서 실행)

```zsh
$ go mod init example.com/project_name
go: creating new go.mod: module example.com/project_name
```

```go
// go.mod 생성 초기 코드 셋
module example.com/project_name

go 1.20 // go version
```

### Run Go

1. go파일이(main.go) 있는 directory로 이동
2. $ go build
   -> 해당 디렉토리 이름의 실행 파일 생성
3. $ go run main.go
   -> go 파일 실행
