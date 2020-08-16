/*
 * @author: stair-go
 * @date: 2020/5/27
 */
package util

import (
	"bytes"
	"github.com/axgle/mahonia"
	"golang.org/x/net/html/charset"
	"io/ioutil"
	"strconv"
	"strings"
)

func ParseStr(hqStr string) []string {
	baseData := strings.Split(hqStr, "\"")
	base := strings.Split(baseData[1], ",")
	return base
}

// 编码
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

//转换字符编码格式，例如文本是gb2312的，现在转换为utf-8：ConvretCharacterEncoding(str, "gb2312", "utf-8")
func ConvertCharacterEncoding(msg string) (result string,err error){
	r := bytes.NewReader([]byte(msg))
	d, err := charset.NewReader(r, "gb2312")
	if err != nil {
		return
	}
	content, err := ioutil.ReadAll(d)
	if err != nil {
		return
	}
	return string(content),nil
}

func StringToInt(number string) int {
	num, _ := strconv.Atoi(number)
	return num
}

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func Float64ToString(number float64, pointSize int) string {
	num := strconv.FormatFloat(number, 'f', pointSize, 64) // -1可改为7，即保留7位小数
	return num
}

func StringToFloat64(number string) float64 {
	num, _ := strconv.ParseFloat(number, 64)
	return num
}