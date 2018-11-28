# Web Terminal

[![CircleCI](https://circleci.com/gh/overmike/webterminal.svg?style=svg)](https://circleci.com/gh/overmike/webterminal)
[![Build Status](https://travis-ci.org/overmike/webterminal.svg?branch=master)](https://travis-ci.org/overmike/webterminal)
[![codecov](https://codecov.io/gh/overmike/webterminal/branch/master/graph/badge.svg)](https://codecov.io/gh/overmike/webterminal)

```
This is an experimental repo, use at your own risk
```

# How to use
Assume you have golang workspace setup
```
go get -u github.com/overmike/webterminal
```
Run it with ${GOPATH}/bin in your ${PATH}
```
    webterminal serve
```
Run it without GOPATH setup
```
    ${GOPATH}/bin/webterminal
```

You can then access browser 8081 port for web terminal


# Development
## go gettable package
In order to make it go gettable, we bundle the webpack build into go binary through github.com/GeertJohan/go.rice

## golang Vendoring
We are opinionated, we choose to commit vendor directory in order to ensure reproducible build.

details see https://github.com/golang/dep/blob/master/docs/FAQ.md#should-i-commit-my-vendor-directory

## Makefile
....
