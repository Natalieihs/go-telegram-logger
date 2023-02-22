package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime/debug"
)

func sendLogToTerraform(err error, token string, chatId string) error {

	//获取异常消息的堆栈和跟踪
	stackTrace := string(debug.Stack())
	message := "An exception occurred: " + err.Error() + "+stackTrace: " + stackTrace

	requestUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", token, chatId, url.QueryEscape(message))
	//创建http客户端并发送请求
	client := &http.Client{}
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return fmt.Errorf("http status code is %d", response.StatusCode)
	}
	return nil
}

func main() {
	botToken := os.Getenv("BOT_TOKEN")
	chatId := os.Getenv("CHAT_ID")
	err := fmt.Errorf("test error")
	if err != nil && botToken != "" && chatId != "" {
		err = sendLogToTerraform(err, botToken, chatId) //发送异常日志到telegram
		if err != nil {
			fmt.Println("send error log to telegram failed, error: ", err)
		}
	}
}
