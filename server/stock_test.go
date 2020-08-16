/*
 * @author: stair-go
 * @date: 2020/7/26
 */
package server

import (
	"fmt"
	"testing"
)

func TestGenShortID(t *testing.T) {
	xx, _ := QueryStock("sh600000")
	fmt.Println(xx)
}