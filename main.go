package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Bool("no-wechat", false, "don't intercept any wechat request")
	flag.Bool("no-douyin", false, "don't intercept any douyin request")
	flag.Bool("no-wechat-block", false, "don't block wechat user who posts images or videos in red color")
	flag.Bool("no-douyin-block", false, "don't block douyin user who posts images or videos in red color")
	flag.Int("red", 10, `treat any image or video (first frame) having more than
this percentage of red-like colors as "in red color"`)
	flag.Bool("blank", false, "turn any image or video into blank black image or video")
	flag.String("c", "~/.unred.conf", "location of the config file")
	flag.Bool("C", false, "create or update config file and exit")
	flag.Bool("v", false, "verbosive")
	flag.Bool("version", false, "print version and exit")
	flag.Usage = func() {
		fmt.Println("Usage of unred [OPTIONS]")
		fmt.Println(`
This utility intercepts HTTP requests and turns red color in any image or video
into black color.

OPTIONS:
`)
		flag.PrintDefaults()
	}
	flag.Parse()
	fmt.Println("WORK IN PROGRESS!")
}
