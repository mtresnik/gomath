# gomath
[![build status](https://github.com/mtresnik/gomath/actions/workflows/go.yml/badge.svg)](https://github.com/mtresnik/gomath/actions/workflows/go.yml/)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://github.com/mtresnik/gomath/blob/main/LICENSE)
[![version](https://img.shields.io/badge/version-1.1.7-blue)](https://github.com/mtresnik/gomath/releases/tag/v1.1.7)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-green.svg?style=flat-square)](https://makeapullrequest.com)
<hr>

gomath (pronounced go-math) is a Go implementation of various linear algebra / geometrical structures.


### Sample Code

In your project run:
```
go mod download github.com/mtresnik/goutils
go mod download github.com/mtresnik/gomath 
```

Your `go.mod` file should look like this:
```go 
module mymodule

go 1.23.3

require github.com/mtresnik/gomath v1.1.7
```


Then in your go files you should be able to run different common linear algebra operations:

```go 
package main

import "github.com/mtresnik/gomath/pkg/gomath"

func main() {
	p1 := gomath.NewPoint(0, 1.0, 2.0, 3.0, 4.0)
	p2 := gomath.NewPoint(5, 5, 5, 5)
	v1 := p1.Subtract(*p2)
	println(v1.String())
	
	norm := v1.Normalize()
	println(norm.String())
	mag := norm.Magnitude()
	println(mag)
}
```