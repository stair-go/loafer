/*
 * @author: stair-go
 * @date: 2020/8/16
 */
package acl

import (
	"github.com/stair-go/loafer/share"
	"github.com/stair-go/loafer/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Stock interface {
	// 查询指数
	QueryBroadMarket(code string) (exponentInfo share.ExponentInfo, err error)
	// 查询股票
	QueryStock(code string) (stockInfo share.StockInfo, err error)
}


// 查询指数
func QueryBroadMarket(code string) (exponentInfo share.ExponentInfo, err error) {
	res, err := http.Get("http://hq.sinajs.cn/list=" + code)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		return exponentInfo, fmt.Errorf("http statusCode:%d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return exponentInfo, err
	}
	base := util.ParseStr(string(result))
	//exponentInfo.Name = BroadMarketName[code]
	exponentInfo.Name = util.ConvertToString(base[0], "gbk", "utf-8")
	currentPoints := strings.Split(base[1], ".")
	exponentInfo.CurrentPoints = currentPoints[0]
	exponentInfo.Gszzl = base[3]
	if len(base[5]) > 5 {
		base[5] = base[5][0 : len(base[5])-4]
		base[5] += "亿"
	}
	exponentInfo.Turnover = base[5]
	return
}

// 查询股票
func QueryStock(code string) (stockInfo share.StockInfo, err error) {
	res, err := http.Get("http://hq.sinajs.cn/list=" + code)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		return stockInfo, fmt.Errorf("http statusCode:%d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return stockInfo, err
	}
	base := util.ParseStr(string(result))
	if len(base) < 3 {
		return stockInfo, fmt.Errorf("股票代码错误:%s", code)
	}
	stockInfo.Name = util.ConvertToString(base[0], "gbk", "utf-8")

	stockInfo.OpenPrice = base[1]
	stockInfo.CurrentPrice = base[3]

	distFloat, err := strconv.ParseFloat(base[2], 64)
	if err != nil {
		return
	}
	distFloat2, err := strconv.ParseFloat(base[3], 64)
	if err != nil {
		return
	}
	increase := (distFloat2 - distFloat) / distFloat * 100
	stockInfo.Increase = fmt.Sprintf("%.2f", increase)
	increaseStrings := strings.Split(stockInfo.Increase, ".")
	// 不需要转了
	if len(increaseStrings) < 2 || len(increaseStrings[1]) < 2 {
		return
	}
	stockInfo.Increase = increaseStrings[0] + "." + increaseStrings[1][:2]
	return
}
