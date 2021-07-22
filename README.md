## Simple terminal spinner with go

![gif example](https://github.com/sagarsearce/gospin/blob/master/ezgif.com-gif-maker.gif)

### Sample code 

```go
package main

import (
	"time"

	"github.com/sagarsearce/gospin"
)

func main() {

	myLoder := gospin.CreateLoder("500ms", "-|_|-|_|-|_")
	go myLoder.StartLoading("stated tasl 1..")
	time.Sleep(5 * time.Second)
	myLoder.StopLoading()
}

```


### API

1. CreateLoder 

#### Args: 
    - dalay in ms string
    - pattern that you want to iterate on string
    
after you create your spinner it has function like start and stop, start will take task string as argument.stope will stop the loder.
