package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type CommonResult struct {
	ErrorCode    int         `json:"ErrorCode"`
	ErrorMessage string      `json:"ErrorMessage"`
	Data         interface{} `json:"Data"`
}

type SearchResult struct {
	Word        string `json:"Word"`
	Translation string `json:"Translation"`
}

//const ServerUrl = "http://localhost:4000/search/"
const ServerUrl = "https://www.wordcounter007.com/search/"

func main() {
	args := os.Args[1:]
	var wordSB strings.Builder
	argsLength := len(args)
	for i := 0; i < argsLength; i++ {
		wordSB.WriteString(args[i])
		if i != argsLength-1 {
			wordSB.WriteString(" ")
		}
	}
	escapedQuery := url.QueryEscape(wordSB.String())
	responseString, err := requestSearchingWord(escapedQuery)
	if err != nil {
		fmt.Println("服务器返回错误 :", err.Error())
	}
	var responseStruct CommonResult
	err = json.Unmarshal([]byte(responseString), &responseStruct)
	if err != nil {
		fmt.Println("JSON解析错误:", err)
		return
	}
	if responseStruct.ErrorCode != 0 {
		fmt.Println("error message :", responseStruct.ErrorMessage)
		return
	}
	searchResultData := responseStruct.Data.(map[string]interface{})
	fmt.Println(strings.ReplaceAll(searchResultData["Translation"].(string), "\\n", "\n"))
}

func requestSearchingWord(word string) (string, error) {
	response, err := http.Get(ServerUrl + word)
	if err != nil {
		fmt.Println("HTTP请求错误:", err)
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应错误:", err)
		return "", err
	}
	return string(body), nil
}