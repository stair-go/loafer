/*
 * @author: Wu Zhihui
 * @date: 2020/8/16
 */
package control

import (
	"github.com/stair-go/loafer/server"
	"log"
	"net/http"
)

func ControlInit(addr string)  {
	http.HandleFunc("/", query)
	http.HandleFunc("/add", add)
	http.HandleFunc("/add/stock", addStock)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/delete/stock", deleteStock)
	http.HandleFunc("/query", query)
	log.Println("Starting v1 server ...")
	log.Fatal(http.ListenAndServe(addr, nil))
}

func add(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	getcode := query.Get("code")
	if getcode!="" {
		server.AddFundCode(getcode)
	}
	allFund, err := server.QueryAllFund()
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	err = server.SendFund(allFund)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, _ = w.Write([]byte("添加成功!"))
}

func addStock(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	getcode := query.Get("code")
	if getcode!="" {
		server.AddStock(getcode)
	}
	allFund, err := server.QueryAllStock()
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	err = server.SendStockInfo(allFund)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("股票代码添加成功!"))
}

func delete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	getcode := query.Get("code")
	server.DeleteFundCode(getcode)
	allFund, err := server.QueryAllFund()
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	err = server.SendFund(allFund)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("删除成功!"))
}


func deleteStock(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	getcode := query.Get("code")
	server.DeleteStock(getcode)
	allFund, err := server.QueryAllStock()
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	err = server.SendStockInfo(allFund)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("删除成功!"))
}


func query(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		return
	}
	allFund, err := server.QueryAllFund()
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	err = server.SendFund(allFund)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	allExponent, err := server.QueryAllExponent()
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	err = server.SendExponentInfo(allExponent)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	allStock, err := server.QueryAllStock()
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	err = server.SendStockInfo(allStock)
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("查询成功!"))
}
