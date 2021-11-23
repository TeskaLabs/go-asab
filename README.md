# Asynchronous Server App Boilerplate (ASAB)

_A micro-service framework for Go_

Modeled after Pythonic https://github.com/TeskaLabs/asab

## Quick start

`main.go`

```go
package main

import (
	"github.com/teskalabs/go-asab/asab"
)

type MyApplication struct {
	asab.Application

	WebService       asab.WebService
}

func main() {
	asab.AddConfigDefaults("general", map[string]string{
		"config_file": "./etc/my.conf",
	})

	asab.AddConfigDefaults("web", map[string]string{
		"listen": "[::]:8895",
	})

	MyApp := new(MyApplication)
	MyApp.Application.Initialize()
	defer MyApp.Finalize()

	MyApp.WebService.Initialize(&MyApp.Application)
	defer MyApp.WebService.Finalize()

	MyApp.Run()
}
```
