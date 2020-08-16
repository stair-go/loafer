/*
 * @author: stair-go
 * @date: 2020/8/16
 */
package acl

import (
	"bytes"
	"daydaytest/share"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

// 重写该接口,则可换取到发送,dd.微信....(现只实现了dd)
type SendMessage interface {
	// 发送基金
	SendFund(allFs []share.FundInfo, robotCallbackUrl []string) (err error)
	// 发送指数
	SendExponentInfo(allFs []share.ExponentInfo, robotCallbackUrl []string) (err error)
	// 发送股票
	SendStockInfo(allStock []share.StockInfo, robotCallbackUrl []string) (err error)
}

// 获取钉钉发送机器人消息模板
func GetTemplate() map[string]interface{} {
	var ddMarkdown = make(map[string]interface{})
	ddMarkdown["msgtype"] = "markdown"
	ddMarkdown["markdown"] = map[string]string{
		"title": "加!加特么的",
		"text":  "",
	}
	ddMarkdown["at"] = map[string]interface{}{
		"isAtAll": false,
	}
	return ddMarkdown
}

func SendFund(allFs []share.FundInfo, robotCallbackUrl []string) (err error) {
	sort.Slice(allFs, func(i, j int) bool {
		f1, _ := strconv.ParseFloat(allFs[i].Gszzl, 64)
		f2, _ := strconv.ParseFloat(allFs[j].Gszzl, 64)
		return f1 > f2
	})
	var fundInfo share.FundInfo
	var mdBody = "|  **基金名**   |  **估值**  |\n|  ----  | ----  |\n"
	for _, fundInfo = range allFs {
		mdBody += "| -------------------------- | ------ |\n"
		mdBody += "| " + fundInfo.Name + "  | " + fundInfo.Gszzl + " | \n"
	}

	template := GetTemplate()
	msgBody := template["markdown"].(map[string]string)
	msgBody["text"] = mdBody
	ddDate, err := json.Marshal(template)
	if err != nil {
		return
	}

	for _, value := range robotCallbackUrl {
		resp, err := http.Post(value, "application/json;charset=UTF-8", bytes.NewReader(ddDate))
		if resp != nil {
			defer resp.Body.Close()
			respBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("ioutil.ReadAll%v %s", err, respBytes)
				return err
			}
			fmt.Println("dingding的响应:", string(respBytes))
		}
		if err != nil {
			fmt.Printf("client.Do%v", err)
			return err
		}
	}
	return
}

func SendExponentInfo(allFs []share.ExponentInfo, robotCallbackUrl []string) (err error) {
	var exponentInfo share.ExponentInfo
	var mdBody = "|  **大盘指数**   |  **点数**  |  **成交额**  |  **估值**  | \n |  ----  |   ----  |  ----  |----  | \n"
	for _, exponentInfo = range allFs {
		mdBody += "| ---------- |   ----  |  ----------------   | ------ |\n"
		mdBody += "| " + exponentInfo.Name + "  | " + exponentInfo.CurrentPoints + "_  | " + exponentInfo.Turnover + "  | " + exponentInfo.Gszzl + " | \n"
	}

	template := GetTemplate()
	msgBody := template["markdown"].(map[string]string)
	msgBody["text"] = mdBody
	ddDate, err := json.Marshal(template)
	if err != nil {
		return
	}

	for _, value := range robotCallbackUrl {
		resp, err := http.Post(value, "application/json;charset=UTF-8", bytes.NewReader(ddDate))
		if resp != nil {
			defer resp.Body.Close()
			respBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("ioutil.ReadAll%v %s", err, respBytes)
				return err
			}
			fmt.Println("dingding的响应:", string(respBytes))
		}
		if err != nil {
			fmt.Printf("client.Do%v", err)
			return err
		}
	}
	return
}

// 发送基金信息到dingding
func SendStockInfo(allStock []share.StockInfo, robotCallbackUrl []string) (err error) {
	var exponentInfo share.StockInfo
	var mdBody = "|  **股票名称**  |  **当前价**  |  **幅度**  | \n |  ----  |   ----  |  ----  | \n"
	for _, exponentInfo = range allStock {
		mdBody += "| ---------- |   ----  |  ----------------  |\n"
		mdBody += "| " + exponentInfo.Name + "  | " + exponentInfo.CurrentPrice + " | " + exponentInfo.Increase + " | \n"
	}

	template := GetTemplate()
	msgBody := template["markdown"].(map[string]string)
	msgBody["text"] = mdBody
	ddDate, err := json.Marshal(template)
	if err != nil {
		return
	}

	for index, value := range robotCallbackUrl {
		// 暂时跳过1
		if index == 0 {
			continue
		}
		resp, err := http.Post(value, "application/json;charset=UTF-8", bytes.NewReader(ddDate))
		if resp != nil {
			defer resp.Body.Close()
			respBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("ioutil.ReadAll%v %s", err, respBytes)
				return err
			}
			fmt.Println("dingding的响应:", string(respBytes))
		}
		if err != nil {
			fmt.Printf("client.Do%v", err)
			return err
		}
	}
	return
}
