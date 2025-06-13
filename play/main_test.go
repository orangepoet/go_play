package play

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestX(t *testing.T) {
	t.Log("test")
}

func TestWxURLLink(t *testing.T) {
	// 定义请求的URL
	url := "https://api.weixin.qq.com/wxa/generate_urllink?access_token=ACCESS_TOKEN"

	// 定义要发送的JSON数据
	jsonData := []byte(`{"key1":"value1", "key2":"value2"}`)

	// 创建一个新的POST请求
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// 打印响应状态和响应体
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}

// generateWxURLLinkWithClient is a version of GenerateWxURLLink that accepts a custom HTTP client
func generateWxURLLinkWithClient(accessToken string, request WxURLLinkRequest, client *http.Client) (*WxURLLinkResponse, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/generate_urllink?access_token=%s", accessToken)

	// Convert request to JSON
	jsonData, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request using the provided client
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse response
	var response WxURLLinkResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	// Check for API errors
	if response.Errcode != 0 {
		return &response, fmt.Errorf("API error: %d %s", response.Errcode, response.Errmsg)
	}

	return &response, nil
}

type WxURLLinkRequest struct {
	Path           string `json:"path"`
	Query          string `json:"query"`
	IsExpire       bool   `json:"is_expire"`
	ExpireType     int    `json:"expire_type"`
	ExpireTime     int64  `json:"expire_time"`
	ExpireInterval int    `json:"expire_interval"`
}

type WxURLLinkResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	URLLink string `json:"url_link"`
}

func TestTimeTruncate(t *testing.T) {
	n, _ := time.ParseInLocation("2006-01-02 15:04:05.000", "2025-06-11 10:00:00.999", time.Local)
	m := n.Truncate(time.Second)

	fmt.Println(n)
	fmt.Println(m)
}

func TestLoopParallel(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		// j := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("i: %d\n", i)
		}()
	}
	wg.Wait()
}
