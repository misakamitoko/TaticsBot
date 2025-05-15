package fetchtaticsdata

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func saveHtml(reader io.ReadCloser) error {
	outputFile, err := os.Create("tantic.html")
	if err != nil {
		fmt.Printf("create html file failed\n")
		return err
	}
	defer outputFile.Close()
	_, err = io.Copy(outputFile, reader)
	if err != nil {
		fmt.Printf("write to file failed\n")
		return err
	}
	return err
}

func fetchData() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://op.gg/zh-cn/tft/meta-trends/comps", nil)
	if err != nil {
		fmt.Printf("create new request failed\n")
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) "+
		"Chrome/114.0.0.0 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("get html failed\n")
		return err
	}
	defer res.Body.Close()
	err = saveHtml(res.Body)
	if err != nil {
		fmt.Printf("save html failed\n")
		return err
	}
	return nil
}
