package main

import "fmt"
import "time"
import "braingoo/wellknown"
import "braingoo/api"

func main() {
    fmt.Println("Hello, world. ", time.Now(), api.Instance());

	wellknown.NodeInfo();
}

