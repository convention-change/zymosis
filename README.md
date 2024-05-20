[![ci](https://github.com/convention-change/zymosis/workflows/ci/badge.svg?branch=main)](https://github.com/convention-change/zymosis/actions/workflows/ci.yml)

[![go mod version](https://img.shields.io/github/go-mod/go-version/convention-change/zymosis?label=go.mod)](https://github.com/convention-change/zymosis)
[![GoDoc](https://godoc.org/github.com/convention-change/zymosis?status.png)](https://godoc.org/github.com/convention-change/zymosis)
[![goreportcard](https://goreportcard.com/badge/github.com/convention-change/zymosis)](https://goreportcard.com/report/github.com/convention-change/zymosis)

[![GitHub license](https://img.shields.io/github/license/convention-change/zymosis)](https://github.com/convention-change/zymosis)
[![GitHub latest SemVer tag)](https://img.shields.io/github/v/tag/convention-change/zymosis)](https://github.com/convention-change/zymosis/tags)
[![GitHub release)](https://img.shields.io/github/v/release/convention-change/zymosis)](https://github.com/convention-change/zymosis/releases)

## for what

- project used to cli with project mark git res by git commit id

## Features

- [x] mark project res by git commit
    - [x] support golang project mark
- [ ] more perfect test case coverage
- [ ] more perfect benchmark case

## usage

```bash
# install at ${GOPATH}/bin
$ go install -v github.com/convention-change/zymosis/cmd/zymosis@latest
# install version v1.0.0
$ go install -v github.com/convention-change/zymosis/cmd/zymosis@v1.0.0
```

### golang project

```bash
# init project need code 
$ zymosis init
# if want update code, just use
$ zymosis init --coverage-exist-file

# then before CI or release binary run as
$ zymosis -g go 
```

- use at go code show res mark code

```go
    fmt.Println(zymosis.MainProgramRes())
```

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/convention-change/zymosis)](https://github.com/convention-change/zymosis/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

请阅读有关 [贡献者指南](.github/CONTRIBUTING_DOC/zh-CN/CONTRIBUTING.md) 以获取更多如何入门的信息