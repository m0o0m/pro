### 一些工具的集合

##### 支付的排序工具
 - 首先定义一个结构体,并给结构体字段打上tag
 - 设置完值后,调用sort函数,返回的就是排序好的字符串了
 
```go
type UnifyOrderReq struct {
	Nonce_str        string `sort:"nonce_str"`
	Appid            string `sort:"appid"`
	Body             string `sort:"body"`
	Trade_type       string `sort:"trade_type"`
	Total_fee        int    `sort:"total_fee"`
	Mch_id           string `sort:"mch_id"`
	Notify_url       string `sort:"notify_url"`
	Out_trade_no     string `sort:"out_trade_no"`
	Spbill_create_ip string `sort:"spbill_create_ip"`
	Sign             string `sort:"sign"`
}

```
```go
func TestSort(t *testing.T) {
	var yourReq UnifyOrderReq
	AppSecret := "12345678910111213141516171819202"
	yourReq.Appid = "wxxxxxxxxxxxxxxxxx" //微信开放平台我们创建出来的app的app id
	yourReq.Body = "shangpinmiaoshu"
	yourReq.Mch_id = "1111111111" //商户号
	yourReq.Nonce_str = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXx"
	yourReq.Notify_url = "http://www.baidu.com/"
	yourReq.Trade_type = "NATIVE"
	yourReq.Trade_type = "JSAPI"
	yourReq.Spbill_create_ip = "0.0.0.0"
	yourReq.Total_fee = 1                           //单位是分，这里是1毛钱

	currentTime := time.Now().UnixNano()
	r := rand.New(rand.NewSource(currentTime))
	yourReq.Out_trade_no = strconv.Itoa(int(currentTime)) + strconv.Itoa(r.Intn(1000000)) //后台系统单号
	pp := Sort(&yourReq, "sort") + "&key=" + AppSecret
	log.Println(pp)
}

```
输出结果
```
appid=wxxxxxxxxxxxxxxxxx&body=shangpinmiaoshu&mch_id=1111111111&nonce_str=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXx&notify_url=http://www.baidu.com/&out_trade_no=1501735009014394202563796&spbill_create_ip=0.0.0.0&total_fee=1&trade_type=JSAPI&key=12345678910111213141516171819202
```
如果使用的是NoSort()方法,结果则是按照filed顺序排列
```
nonce_str=XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXx&appid=wxxxxxxxxxxxxxxxxx&body=shangpinmiaoshu&trade_type=JSAPI&total_fee=1&mch_id=1111111111&notify_url=http://www.baidu.com/&out_trade_no=1501735116528896644227487&spbill_create_ip=0.0.0.0&key=12345678910111213141516171819202
```

##### 一个带Close功能的 WaitGroupPlus
```go
func main() {
	fmt.Println(wgF())
	//var wg sync.WaitGroup
	//wg.Add(1)
	//wg.Wait()
	select {
	case <-time.After(1<<63 - 1):

	}
}
func wgF() string {
	var wg = lutils.WaitGroupPlus{}
	wg.Add(13)
	go func() {
		//<-time.After(3e9)
		time.Sleep(2e9)
		go func() {
			defer wg.Close()
			time.Sleep(4e9)
			fmt.Println("线程还活着")
		}()
		wg.Done()
	}()
	wg.Wait()
	return "exit"
}
```
##### 线程安全的err