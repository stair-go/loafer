/*
 * acl:防腐层,调用外部接口
 */
package acl

import (
	"daydaytest/share"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 查询基金
func QueryFund(code string) (fundInfo share.FundInfo, err error) {
	res, err := http.Get("http://fundgz.1234567.com.cn/js/" + code + ".js")
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return
	}
	if res.StatusCode != http.StatusOK {
		return fundInfo, fmt.Errorf("http statusCode:%d", res.StatusCode)
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fundInfo, err
	}
	s := string(result)
	s = strings.Replace(s, "jsonpgz(", "", -1)
	s = strings.Replace(s, ");", "", -1)
	err = json.Unmarshal([]byte(s), &fundInfo)
	if err != nil {
		return fundInfo, err
	}
	return
}
