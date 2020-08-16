/*
 * @author: stair-go
 * @date: 2020/5/27
 */
package server

import (
	"daydaytest/acl"
	"daydaytest/share"
)

func QueryAllStock() (allFs []share.StockInfo, err error) {
	var fundInfo share.StockInfo
	for _, value := range AllStock() {
		fundInfo, err = acl.QueryStock(value)
		if err != nil {
			return
		}
		allFs = append(allFs, fundInfo)
	}
	return
}
