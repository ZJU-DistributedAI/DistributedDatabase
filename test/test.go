package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func main(){

	cmds := strings.Split("docker exec ipfs"," ")
	cmd := exec.Command(cmds[0], cmds[1], cmds[2],"ipfs", "add", "-r","/upload2/test.txt")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil{
		fmt.Println("执行出错: ", err)
	}
	fmt.Println(out.String())
}
