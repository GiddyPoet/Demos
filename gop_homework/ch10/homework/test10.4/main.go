//练习 10.4： 创建一个工具，根据命令行指定的参数，报告工作区所有依赖包指定的其它包集合。提示：
// 你需要运行go list命令两次，一次用于初始化包，一次用于所有包。你可能需要用encoding/json（§4.5）包来分析输出的JSON格式的信息。
package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

var (
	RestMap map[string]int
)

type Result struct {
	Imports []string
}

func main() {
	result := []string{}
	RestMap = make(map[string]int)
	TT("github.com/googege/goo", 0)
	for k, _ := range RestMap {
		result = append(result, k)
	}
	fmt.Println(result)
}
func TT(s string, old int) {
	re := new(Result)
	cmd := exec.Command("/usr/local/bin/go", "list", "-e", "-json", s)
	data, err := cmd.Output()
	if err != nil {
		fmt.Print(err)
	}
	json.Unmarshal(data, re)
	for _, v := range re.Imports {
		RestMap[v]++
		if old == len(RestMap) {
			return
		}
		TT(v, len(RestMap))
	}

}
