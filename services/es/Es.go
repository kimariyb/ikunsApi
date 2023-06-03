package es

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego/httplib"
)

var esUrl string

func init() {
	esUrl = "http://127.0.0.1:9200/"
}

func EsSearch(indexName string, query map[string]interface{}, from int, size int, sort []map[string]string) HitsData {
	searchQuery := map[string]interface{}{
		"query": query,
		"from":  from,
		"size":  size,
		"sort":  sort,
	}

	req := httplib.Post(esUrl + indexName + "/_search")
	req.JSONBody(searchQuery)

	str, err := req.String()
	fmt.Println(str)
	if err != nil {
		fmt.Println(err)
	}
	var stb ReqSearchData
	err = json.Unmarshal([]byte(str), &stb)

	return stb.Hits
}

//解析获取到的值
type ReqSearchData struct {
	Hits HitsData `json:"hits"`
}
type HitsData struct {
	Total TotalData     `json:"total"`
	Hits  []HitsTwoData `json:"hits"`
}
type TotalData struct {
	Value    int
	Relation string
}
type HitsTwoData struct {
	Source json.RawMessage `json:"_source"`
}

//添加
func EsAdd(indexName string, id string, body map[string]interface{}) bool {
	req := httplib.Post(esUrl + indexName + "/_doc/" + id)
	req.JSONBody(body)

	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	return true
}

//修改
func EsEdit(indexName string, id string, body map[string]interface{}) bool {
	bodyData := map[string]interface{}{
		"doc": body,
	}

	req := httplib.Post(esUrl + indexName + "/_doc/" + id + "/_update")
	req.JSONBody(bodyData)

	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	return true
}

//删除
func EsDelete(indexName string, id string) bool {
	req := httplib.Delete(esUrl + indexName + "/_doc/" + id)
	str, err := req.String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)
	return true
}
