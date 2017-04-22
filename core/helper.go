package core

import (
	"os"
	"bufio"
	"io"
	"fmt"
	"strings"
)

var (
	evnItems map[string]string
)

func init() {
	evnItems = make(map[string]string)
	f,err := os.Open("./.env");
	defer f.Close()
	fmt.Print(err)
	if nil == err {
		buf := bufio.NewReader(f)
		for {
			line,err := buf.ReadString('\n');
			if io.EOF == err {
				break
			}
			keyValue := strings.Split(line,"=")
			if len(keyValue) != 2 {
				continue
			}
			evnItems[strings.TrimSpace(keyValue[0])] = strings.TrimSpace(keyValue[1])
		}
	} else {
		fmt.Print(err.Error())
	}
}


func Env(name,def string) string  {
	value,exist := evnItems[name]
	if !exist {
		return def
	}else {
		return value
	}
}