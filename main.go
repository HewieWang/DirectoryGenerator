package main

import (
	"fmt"
	"strings"
	"flag"
	"os"
)

func permutation(S string) []string {
    m := make(map[string]bool)
    m[S] = true
    l := len(S)
    for i := 0; i < l-1; i ++ {
        for j := i + 1; j < l; j ++ {
            for s, _ := range m {
                t := []byte(s)
                t[i], t[j] = t[j], t[i]
                m[string(t)] = true
            }
        }
    }
    ret := make([]string, len(m))
    i := 0
    for s, _ := range m {
        ret[i] = s
        i ++
    }
    return ret
}
func Add2file(filename string,str_content string)  {
    fd,_:=os.OpenFile(filename,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
    fd_content:=strings.Join([]string{str_content,"\n"},"")
    buf:=[]byte(fd_content)
    fd.Write(buf)
    fd.Close()
}
func main()  {
	var left    string
	var right   string
	var content string
	var output  string
	flag.StringVar(&left, "l", "", "left")
	flag.StringVar(&right, "r", "", "right")
	flag.StringVar(&content, "c", "abc", "content")
	flag.StringVar(&output, "o", "dictionary.txt", "output")
	flag.Parse()

	if len(left)>0 {
		left_arr := strings.Split(left, `,`)
		for _, l := range left_arr {
			for _, p := range permutation(content) {
					if len(right)>0{
						right_arr := strings.Split(right, `,`)
						for _, r := range right_arr {
									fmt.Println(l+p+r)
									Add2file(output,l+p+r)
						}
					}else{
						fmt.Println(l+p)
						Add2file(output,l+p)
					}

			}
		}
	}else{
		right_arr := strings.Split(right, `,`)
		for _, r := range right_arr {
			for _, p := range permutation(content) {
					fmt.Println(p+r)
					Add2file(output,p+r)
			}
		}
	}
}
