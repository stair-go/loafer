/*
 * acl:防腐层,调用外部服务接口,此处粗糙的借鉴了领域设计,实际这些方法均可抽象为go接口,便于项目的维护
 */
package acl

import (
	"github.com/stair-go/loafer/share"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// 此处粗糙的借鉴了领域设计,将需要的外部方法均可抽象为go接口
// 使用方调用实现了,因项目简单,并未使用结构体实现
// 便于后续维护,若修改爬取逻辑,可无需修改server层的调用
type Fund interface {
	QueryFund(code string) (fundInfo share.FundInfo, err error)
}

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
