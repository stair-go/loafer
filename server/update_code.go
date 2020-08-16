/*
 * @author: stair-go
 * @date: 2020/5/27
 */
package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

var f = &GlobalFundCode{}

type FundCodeData struct {
	CodeAttr         []string "json:`CodeAttr`" // 基金代码数组
	RobotCallbackUrl []string                   // 机器人回调地址
	Exponent         []string                   // 大盘指数
	Stock         []string                   // 大盘指数
}

type GlobalFundCode struct {
	sync.RWMutex          // 读写锁
	FundCode FundCodeData // 基金代码数组
	Path     string       // 配置路径
}

func SetFundCodeData(globalFundCode *GlobalFundCode) {
	f = globalFundCode
}

// 写入持久化文件
func WriteConfigFile() error {
	//f.Lock()
	//defer f.Unlock()

	data, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		return err
	}
	var newf = &GlobalFundCode{}
	err = json.Unmarshal(data, &newf)
	if err != nil {
		return err
	}
	newf.FundCode = f.FundCode
	data, err = json.Marshal(newf)
	if err != nil {
		return err
	}
	f = newf
	fmt.Println(string(data))
	return ioutil.WriteFile(f.Path, []byte(data), 0777)
}

// 查询所有基金代码
func AllFundCode() (allF []string) {
	f.Lock()
	defer f.Unlock()
	// 创建新的slice,实现深拷贝
	allF = make([]string, len(f.FundCode.CodeAttr), len(f.FundCode.CodeAttr))
	copy(allF, f.FundCode.CodeAttr)
	return
}

// 查询所有指数代码
func AllExponent() (allF []string) {
	f.Lock()
	defer f.Unlock()
	// 创建新的slice,实现深拷贝
	allF = make([]string, len(f.FundCode.Exponent), len(f.FundCode.Exponent))
	copy(allF, f.FundCode.Exponent)
	return
}

func AllStock() (allF []string) {
	f.Lock()
	defer f.Unlock()
	// 创建新的slice,实现深拷贝
	allF = make([]string, len(f.FundCode.Stock), len(f.FundCode.Stock))
	copy(allF, f.FundCode.Stock)
	return
}

// 添加基金
func AddFundCode(code string) {
	f.Lock()
	defer f.Unlock()
	for _, value := range f.FundCode.CodeAttr {
		if value == code {
			return
		}
	}
	f.FundCode.CodeAttr = append(f.FundCode.CodeAttr, code)
	WriteConfigFile()
}

// 删除删除基金
func DeleteFundCode(code string) {
	f.Lock()
	defer f.Unlock()
	for i := 0; i < len(f.FundCode.CodeAttr); i++ {
		if f.FundCode.CodeAttr[i] == code {
			f.FundCode.CodeAttr = append(f.FundCode.CodeAttr[0:i], f.FundCode.CodeAttr[i+1:]...)
			i--
		}
	}
	WriteConfigFile()
}

// 添加股票代码
func AddStock(code string) {
	f.Lock()
	defer f.Unlock()
	for _, value := range f.FundCode.Stock {
		if value == code {
			return
		}
	}
	f.FundCode.Stock = append(f.FundCode.Stock, code)
	WriteConfigFile()
}

// 删除股票代码
func DeleteStock(code string) {
	f.Lock()
	defer f.Unlock()
	for i := 0; i < len(f.FundCode.Stock); i++ {
		if f.FundCode.Stock[i] == code {
			f.FundCode.Stock = append(f.FundCode.Stock[0:i], f.FundCode.Stock[i+1:]...)
			i--
		}
	}
	WriteConfigFile()
}