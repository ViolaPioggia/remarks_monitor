package main

import (
	"bufio"
	"fmt"
	"net/http"

	_ "net/http/pprof"
	"os"
	"remarks_monitor/common/tool"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(time.Now())
		for i := 0; i <= 100; i++ {

			file, err := os.Open(tool.GetWD() + "/data/remarks_monitor/input/input" + strconv.Itoa(i))
			if err != nil {
				fmt.Println("无法打开文件:", err)
				return
			}
			defer file.Close()

			// 创建映射保存字符串和行数
			counts := make(map[string]int)

			// 逐行读取文件内容
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := scanner.Text()

				// 检查映射中是否已存在该键
				counts[line]++
			}

			// 检查扫描过程中是否发生错误
			if err := scanner.Err(); err != nil {
				fmt.Println("文件扫描错误:", err)
				return
			}

			// 遍历映射，输出结果
			for line, count := range counts {
				fmt.Printf("字符串: %s，行数: %d\n", line, count)
			}
		}
		fmt.Println(time.Now())
	})
	http.ListenAndServe("localhost:9999", nil)
}
