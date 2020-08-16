/*
 * @author: stair-go
 * @date: 2020/5/27
 */
package main

import (
	"github.com/stair-go/loafer/run"
)

//v1
func main() {
	err := run.Run()
	if err != nil {
		panic(err)
	}
}
