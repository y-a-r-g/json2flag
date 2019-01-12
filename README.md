# json2flag
Go config helper. Reads all flags from json file as alternative configuration method.

###Usage:

config.json
```json
{
  "test-flag": "works"
}
```

main.go
```go
package main

import (
	"flag"
	"github.com/y-a-r-g/json2flag"
)

func main() {
	testFlag := flag.String("test-flag", "", "...")
	
	flag.Parse()
	if err := json2flag.ReadConfigFile("config.json"); err != nil {
		panic(err)
	}
	
	println(*testFlag) //-> works
}
```

Also you can read config from `string` and `[]byte` using `ReadConfigString` and `ReadConfigData` correspondingly.
