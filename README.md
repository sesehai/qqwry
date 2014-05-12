QQWry
=====

纯真IP库 golang 版

>有关纯真IP库的文档详见 http://lumaqq.linuxsir.org/article/qqwry_format_detail.html

使用golang面向对象开发

## 1. 依赖
* iconv: libiconv for go (请确保先装好 `hg`)
```bash
go get github.com/qiniu/iconv
```

## 2. 使用
* 下载
```bash
go get github.com/sesehai/qqwry
```
* 在项目中引入
```go
import (
	"github.com/sesehai/qqwry"
	"fmt"
)

func main() {
	var ip = "112.0.91.210"
	var qqwryfile = "./qqwry.dat"
	fmt.Println(ip)
	file, _ := qqwry.Getqqdata(qqwryfile)
	country, _ := qqwry.Getlocation(file, ip)
	fmt.Println(country)
}

// 输出结果：
// 112.0.91.210
// 江苏省镇江市
```