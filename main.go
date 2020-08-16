/*
 * @author: stair-go
 * @date: 2020/5/27
 */
package main

import (
	"daydaytest/run"
)

//v1
func main() {
	err := run.Run()
	if err != nil {
		panic(err)
	}
}
