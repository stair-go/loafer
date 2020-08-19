/*
 * @author: stair-go
 * @date: 2020/5/27
 */
package run

import (
	"github.com/stair-go/loafer/control"
	"github.com/stair-go/loafer/server"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"time"
)

func Run() error {
	// 配置初始化
	err := configInit("./config.json")
	if err != nil {
		return err
	}
	go Polling()
	// 表现层初始化
	control.ControlInit(":8189")
	return nil
}

// 初始化 GlobalFundCode ---并不优雅
func configInit(path string) error {
	var f = &server.GlobalFundCode{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &f)
	if err != nil {
		return err
	}
	f.Path = path
	server.SetFundCodeData(f)
	return nil
}

func Polling() {
	ticker := time.NewTicker(time.Minute * 30)
	var err error
	for {
		select {
		case <-ticker.C:
			if !whetherTheCurrentTimeIsOpen(time.Now()) {
				continue
			}
			allFund, err := server.QueryAllFund()
			if err != nil {
				continue
			}
			err = server.SendFund(allFund)
			if err != nil {
				continue
			}

			allExponent, err := server.QueryAllExponent()
			if err != nil {
				continue
			}
			err = server.SendExponentInfo(allExponent)
			if err != nil {
				continue
			}
			allStock, err := server.QueryAllStock()
			if err != nil {
				continue
			}
			err = server.SendStockInfo(allStock)
			if err != nil {
				continue
			}
		}
		if err != nil {
			fmt.Println("轮询出错:", err)
		}
	}
}

func whetherTheCurrentTimeIsOpen(t time.Time) bool {
	yearInt := t.Year()
	monthInt := t.Month() //time.Now().Month().String()
	dayInt := t.Day()
	if zellerFunction2Week(yearInt, int(monthInt), dayInt) == 6 || zellerFunction2Week(yearInt, int(monthInt), dayInt) == 7 {
		return false
	}
	if t.Hour() < 9 || t.Hour() > 15 {
		return false
	}
	return true
}

//获取当天是周几
func zellerFunction2Week(year, month, day int) int {
	var weekday = [7]int{7, 1, 2, 3, 4, 5, 6}
	var y, m, c int
	fmt.Println("ppppppppppp")
	fmt.Println(reflect.TypeOf(month))
	if month >= 3 {
		m = month
		y = year % 100
		c = year / 100
	} else {
		m = month + 12
		y = (year - 1) % 100
		c = (year - 1) / 100
	}
	week := y + (y / 4) + (c / 4) - 2*c + ((26 * (m + 1)) / 10) + day - 1
	if week < 0 {
		week = 7 - (-week)%7
	} else {
		week = week % 7
	}
	whichWeek := int(week)
	return weekday[whichWeek]
}
