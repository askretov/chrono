# Mutable
[![Go Report Card](https://goreportcard.com/badge/github.com/askretov/chrono)](https://goreportcard.com/report/github.com/askretov/chrono)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/1c52c7899e544969b1d83896dbc2b9c4)](https://www.codacy.com/app/askretov/go-mutable?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=askretov/mutable&amp;utm_campaign=Badge_Grade)
[![codecov](https://codecov.io/gh/askretov/chrono/branch/master/graph/badge.svg)](https://codecov.io/gh/askretov/chrono)
[![Build Status](https://travis-ci.org/askretov/chrono.svg?branch=master)](https://travis-ci.org/askretov/chrono)
[![GoDoc](https://godoc.org/github.com/askretov/chrono?status.svg)](https://godoc.org/github.com/askretov/chrono)
[![Licenses](https://img.shields.io/badge/license-mit-brightgreen.svg)](https://opensource.org/licenses/BSD-3-Clause)

## Introduction
Package chrono provides basic chronometric features to measure time elapsed in various cases.
Package usage is as simple as native clock timer in your phone but also has some useful features like cumulative measurements.

## Usage
### Installation
```go
go get github.com/askretov/chrono
```
### Basic example
```go
package main

import "github.com/askretov/chrono"

func main() {
    chrono.Start("test")
    for i := 0; i < 1000000000; i++ {}
    chrono.Stop("test")
}
```
### Laps example
```go
    chrono.Start("test")
    for i := 0; i < 1000000000; i++ {}
    chrono.Lap("test", "lap 1")
    for i := 0; i < 1000000000; i++ {}
    chrono.Lap("test", "lap 2")
    for i := 0; i < 1000000000; i++ {}
    chrono.Lap("test", "lap 3")
    chrono.Stop("test")
```
### Func time capture
```go
	Capture("capture", false, func(){
		for i := 0; i < 1000000000; i++ {}
	})
```