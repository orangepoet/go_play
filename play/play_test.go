package play

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"reflect"
	"slices"
	"sync"
	"testing"
	"time"

	"testing/quick"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/spf13/cast"
)

// TestX 基础测试方法，用于验证测试框架是否正常工作
func TestX(t *testing.T) {
	t.Log("test")
}

// TestWxURLLink 测试微信小程序URL链接生成接口的HTTP请求
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

// GenerateWxURLLinkWithClient 生成微信小程序URL链接，支持自定义HTTP客户端
// 这是 GenerateWxURLLink 的版本，接受自定义的 HTTP 客户端
func GenerateWxURLLinkWithClient(accessToken string, request WxURLLinkRequest, client *http.Client) (*WxURLLinkResponse, error) {
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

// TestTimeTruncate 测试时间截断功能，验证时间精度处理
func TestTimeTruncate(t *testing.T) {
	n, _ := time.ParseInLocation("2006-01-02 15:04:05.000", "2025-06-11 10:00:00.999", time.Local)
	m := n.Truncate(time.Second)

	fmt.Println(n)
	fmt.Println(m)
}

// TestLoopParallel 测试并行循环中的闭包变量捕获问题
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

// TestSort 测试切片排序功能，比较不同排序方法的实现
func TestSort(t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}
	slices.SortFunc(arr, func(a, b int) int {
		ret := cast.ToInt(a > b)
		return ret
	})
	fmt.Println(arr)

	arr2 := []int{5, 4, 3, 2, 1}
	slices.SortFunc(arr2, func(a, b int) int {
		ret := cmp.Compare(a, b)
		return ret
	})
	fmt.Println(arr2)
}

// Test001 测试map零值访问，验证未初始化map的行为
func Test001(t *testing.T) {
	var m map[int]Sth
	s := m[1]
	fmt.Println(s.Name)
}

// Test002 测试切片越界访问，演示切片边界检查
func Test002(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	arr1 := arr[:10]
	fmt.Println(arr1)
}

// Test003 测试时间计算功能，验证time.Until方法的使用
func Test003(t *testing.T) {
	var t1 time.Time = time.Now().AddDate(0, 0, 1)
	u := time.Until(t1)
	fmt.Println(u.Seconds())

	var t2 time.Time = time.Now().AddDate(0, 0, -1)
	u2 := time.Until(t2)
	fmt.Println(u2.Seconds())
}

// Test004 测试闭包缓存模式，实现函数结果的记忆化存储
func Test004(t *testing.T) {
	handler := func() func(i int) bool {
		m := make(map[int]bool)

		return func(i int) bool {
			if v, ok := m[i]; ok {
				return v
			} else {
				v := getResult(i)
				m[i] = v
				return v
			}
		}
	}

	h := handler()
	for i := 0; i < 10; i++ {
		ret := h(1)
		fmt.Println(ret)
	}
}

// getResult 计算函数，用于测试缓存机制
func getResult(i int) bool {
	fmt.Printf("计算 getResult(%d)\n", i) // 添加日志来验证是否重复计算
	return i%2 == 0
}

// Test005 测试切片动态增长，验证append操作的行为
func Test005(t *testing.T) {
	var arr []int

	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}

	fmt.Println(arr)
}

// Test006 测试函数调用优化，比较range和传统for循环的性能差异
func Test006(t *testing.T) {
	fmt.Println("=== 使用 range (只执行一次) ===")
	for _, item := range getSlice().arr {
		fmt.Println(item)
	}

	fmt.Println("\n=== 使用传统 for 循环 (会重复执行) ===")
	for i := 0; i < getSlice().len; i++ {
		fmt.Println(getSlice().arr[i])
	}

	fmt.Println("\n=== 正确的传统 for 循环写法 (只执行一次) ===")
	result := getSlice()
	for i := 0; i < result.len; i++ {
		fmt.Println(result.arr[i])
	}
}

// Result 结果结构体，包含数组和长度信息
type Result struct {
	arr []int
	len int
}

// getSlice 创建并返回包含数组和长度的结果对象
func getSlice() *Result {
	fmt.Println("调用 getSlice()") // 添加日志验证调用次数
	arr := make([]int, 10)
	for i := 0; i < 10; i++ {
		arr[i] = i
	}
	return &Result{
		arr: arr,
		len: len(arr),
	}
}

// Test007 测试简单超时等待，使用time.Sleep实现固定时间等待
func Test007(t *testing.T) {
	// 方式1：如果只是想等待30秒后结束，直接用 time.Sleep
	fmt.Println("开始等待...")
	time.Sleep(30 * time.Second)
	fmt.Println("30秒后结束")
}

// Test007WithTicker 测试带定时器的超时控制，使用context和ticker实现
func Test007WithTicker(t *testing.T) {
	// 方式2：如果需要定期执行某些操作，用 context 控制超时
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("超时，程序结束")
			return
		case <-ticker.C:
			fmt.Println("ticker")
		}
	}
}

// Test007Simple 测试简单的超时模式，使用channel和select实现
func Test007Simple(t *testing.T) {
	// 方式3：最简单的超时模式
	done := make(chan bool)

	// 启动一些工作
	go func() {
		time.Sleep(5 * time.Second) // 模拟工作
		done <- true
	}()

	// 等待完成或超时
	select {
	case <-done:
		fmt.Println("工作完成")
	case <-time.After(30 * time.Second):
		fmt.Println("工作超时")
	}
}

// Test008RetryTask 测试任务重试机制，失败时重试直到成功或超时
func Test008RetryTask(t *testing.T) {
	// 模拟一个可能失败的任务
	attemptCount := 0
	task := func() error {
		attemptCount++
		fmt.Printf("尝试第 %d 次...\n", attemptCount)

		// 模拟前几次失败，第4次成功
		if attemptCount < 4 {
			return fmt.Errorf("任务失败 (第%d次尝试)", attemptCount)
		}
		return nil // 成功
	}

	// 执行带重试的任务
	err := retryWithTimeout(task, 30*time.Second, 2*time.Second)
	if err != nil {
		fmt.Printf("任务最终失败: %v\n", err)
	} else {
		fmt.Println("任务成功完成!")
	}
}

// retryWithTimeout 通用的重试函数，支持超时和重试间隔控制
func retryWithTimeout(task func() error, timeout, retryInterval time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ticker := time.NewTicker(retryInterval)
	defer ticker.Stop()

	// 立即执行第一次
	if err := task(); err == nil {
		return nil // 第一次就成功
	} else {
		fmt.Printf("任务失败: %v, 将重试...\n", err)
	}

	// 如果第一次失败，开始定期重试
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("任务超时: %w", ctx.Err())
		case <-ticker.C:
			if err := task(); err == nil {
				return nil // 重试成功
			} else {
				fmt.Printf("任务失败: %v, 继续重试...\n", err)
			}
		}
	}
}

// Test009RetryWithBackoff 测试带指数退避的高级重试机制
func Test009RetryWithBackoff(t *testing.T) {
	attemptCount := 0
	task := func() error {
		attemptCount++
		fmt.Printf("尝试第 %d 次...\n", attemptCount)

		if attemptCount < 3 {
			return fmt.Errorf("任务失败 (第%d次尝试)", attemptCount)
		}
		return nil
	}

	err := retryWithBackoff(task, 30*time.Second, 1*time.Second, 2.0, 5)
	if err != nil {
		fmt.Printf("任务最终失败: %v\n", err)
	} else {
		fmt.Println("任务成功完成!")
	}
}

// retryWithBackoff 带指数退避的重试函数，支持最大重试次数和退避倍数
func retryWithBackoff(task func() error, timeout, initialInterval time.Duration, backoffMultiplier float64, maxRetries int) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	interval := initialInterval

	for attempt := 0; attempt < maxRetries; attempt++ {
		if err := task(); err == nil {
			return nil // 成功
		} else {
			fmt.Printf("任务失败: %v\n", err)
		}

		// 检查是否还有时间重试
		select {
		case <-ctx.Done():
			return fmt.Errorf("任务超时，共尝试 %d 次", attempt+1)
		default:
		}

		// 如果不是最后一次尝试，等待后重试
		if attempt < maxRetries-1 {
			fmt.Printf("等待 %v 后重试...\n", interval)

			timer := time.NewTimer(interval)
			select {
			case <-ctx.Done():
				timer.Stop()
				return fmt.Errorf("任务超时，共尝试 %d 次", attempt+1)
			case <-timer.C:
				// 指数退避：下次等待时间翻倍
				interval = time.Duration(float64(interval) * backoffMultiplier)
			}
		}
	}

	return fmt.Errorf("达到最大重试次数 %d", maxRetries)
}

// Test010 测试JSON序列化，验证结构体嵌套的JSON处理
func Test010(t *testing.T) {
	a := A{Name: "test", B: &B{P1: 1, P2: "2"}}
	// a.P1 = 1
	// a.P2 = "2"
	bytes, _ := json.Marshal(a)
	ret := string(bytes)
	fmt.Println(ret)
}

// Test012 测试channel操作，验证关闭channel后的读取行为
func Test012(t *testing.T) {
	c := make(chan int, 2)

	c <- 1
	c <- 2
	close(c)

	fmt.Println("=== 直接读取 ===")
	fmt.Println(<-c) // 1
	fmt.Println(<-c) // 2
	fmt.Println(<-c) // 0 (零值)

	fmt.Println("\n=== 使用 ok 检查 ===")
	c2 := make(chan int, 2)
	c2 <- 10
	c2 <- 20
	close(c2)

	for {
		value, ok := <-c2
		if !ok {
			fmt.Println("channel 已关闭，退出")
			break
		}
		fmt.Printf("读取到值: %d\n", value)
	}

	fmt.Println("\n=== 使用 range 自动处理 ===")
	c3 := make(chan int, 3)
	c3 <- 100
	c3 <- 200
	c3 <- 300
	close(c3)

	for value := range c3 {
		fmt.Printf("range 读取: %d\n", value)
	}
	fmt.Println("range 自动结束")
}

// Test013 测试gomonkey方法打桩，验证反射方法替换功能
func Test013(t *testing.T) {
	dao := NewDao()
	patch := gomonkey.ApplyMethod(reflect.TypeOf(dao), "GetData", func() (int, error) {
		return 101, nil
	})
	defer patch.Reset()

	service := NewService()
	data, err := service.GetData()
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("data->", data)
}

// Add 简单的加法函数，用于测试quick.Check属性测试
func Add(a, b int) int {
	c := a + b
	fmt.Println("a->", a, "b->", b, "c->", c)
	return c
}

// TestAdd 测试quick.Check属性测试，验证加法函数的数学性质
func TestAdd(t *testing.T) {
	fn := func(a, b int) bool {
		return Add(a, b) == a+b
	}
	fmt.Println("start....")
	if err := quick.Check(fn, nil); err != nil {
		t.Error(err)
	}
}
