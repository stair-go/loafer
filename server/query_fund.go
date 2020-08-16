/*
 * @author: stair-go
 * @date: 2020/5/27
 */
package server

import (
	"daydaytest/acl"
	"daydaytest/share"
)

func QueryAllFund() (allFs []share.FundInfo, err error) {
	var fundInfo share.FundInfo
	for _, value := range AllFundCode() {
		fundInfo, err = acl.QueryFund(value)
		if err != nil {
			return
		}
		allFs = append(allFs, fundInfo)
	}
	return
}

func QueryAllExponent() (allEs []share.ExponentInfo, err error) {
	var exponentInfo share.ExponentInfo
	for _, value := range AllExponent() {
		exponentInfo, err = acl.QueryBroadMarket(value)
		if err != nil {
			return
		}
		allEs = append(allEs, exponentInfo)
	}
	return
}
