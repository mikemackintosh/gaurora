package main

/*
#cgo CFLAGS: -I.CUESDK/lib/x64 -I./CUESDK/include -I./CUESDK/redist/x64
#cgo LDFLAGS: -L.\CUESDK\lib\x64 -L./CUESDK/include -L./CUESDK/redist/x64 -lCUESDK.x64_2015
#include <stdbool.h>
#include <CorsairKeyIdEnum.h>
#include <CorsairLedIdEnum.h>
#include <CUESDK.h>
*/
import "C"

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", C.CorsairGetLedPositions())
}

func abort(msg string, err interface{}) {
	fmt.Printf("%#v, %#v\n", msg, err)

	os.Exit(1)
}
