/*
 * @author: stair-go
 * @date: 2020/5/27
 */
package server

import (
	"daydaytest/acl"
	"daydaytest/share"
)

func SendFund(allFs []share.FundInfo) (err error) {
	return acl.SendFund(allFs, f.FundCode.RobotCallbackUrl)
}

func SendExponentInfo(allFs []share.ExponentInfo) (err error) {
	return acl.SendExponentInfo(allFs, f.FundCode.RobotCallbackUrl)
}

func SendStockInfo(allStock []share.StockInfo) (err error) {
	return acl.SendStockInfo(allStock, f.FundCode.RobotCallbackUrl)
}
