# 📦 Box - is your dependency helping hand
[![Build Status](https://travis-ci.org/cryptopay-dev/pony-box.svg?branch=master)](https://travis-ci.org/cryptopay-dev/pony-box)
[![codecov](https://codecov.io/gh/cryptopay-dev/pony-box/branch/master/graph/badge.svg)](https://codecov.io/gh/cryptopay-dev/pony-box)

## Usage
```go
package main

import (
	"fmt"
	
	"github.com/cryptopay-dev/pony-box"
)

type A struct {}
type B struct {}

func (a A) Ping() string {
	return "ping"
}


func (b B) Pong() string {
	return "pong"
}

func main() {
    b := box.New()
    if err := b.Provide(
    	box.NewProvider(A{}), 
    	box.NewProvider(B{}),
    ); err != nil {
    	panic(err)
    }
    
    if err := b.Invoke(func(a A, b B) error {
    	fmt.Printf("%s - %s\n", a.Ping(), b.Pong())
    	return nil
    }); err != nil {
    	panic(err)
    }
}
```