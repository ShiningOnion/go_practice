package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//函數調用io.Copy(dst, src)會從src中讀取內容，併將讀到的結果寫入到dst中，
//使用這個函數替代掉例子中的ioutil.ReadAll來拷貝響應結構體到os.Stdout，
//避免申請一個緩衝區（例子中的b）來存儲。
//記得處理io.Copy返迴結果中的錯誤。

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		// b, err := ioutil.ReadAll(resp.Body)
		// resp.Body.Close()
		_, err = io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%v", os.Stdout)
	}
}
