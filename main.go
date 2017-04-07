/*
 * 1.Get system time stamp
 * 2.Change the time stamp to unix time
 * 3.Change system time by param
 */
package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var tsConvert *int64 = flag.Int64("c", 0, "Convert time stamp param to normal string.")
var tsAdditon *string = flag.String("a", "0Y:0M:0D:0h:0m:0s", "Add time to system time,string format(1Y:1M:1D:1h:1m:1s)")
var tsSecond *int64 = flag.Int64("d", 0, "Add second to system time.")

func main() {
	flag.Parse()
	if flag.NFlag() == 0 {
		fmt.Print("Current system time stamp:")
		fmt.Println(time.Now().Unix())
		return
	}
	flag.Visit(func(flag *flag.Flag) {
		//fmt.Println(flag)
		switch flag.Name {
		case "c":
			{
				fmt.Println(time.Unix(*tsConvert, 0))
			}
		case "a":
			{
				args := strings.Split(*tsAdditon, ":")
				var year, month, day, hour, mini, sec int
				for _, value := range args {
					if len(value) < 2 {
						continue
					}
					switch value[len(value)-1] {
					case 'Y':
						year, _ = strconv.Atoi(value[:len(value)-1])
					case 'M':
						month, _ = strconv.Atoi(value[:len(value)-1])
					case 'D':
						day, _ = strconv.Atoi(value[:len(value)-1])
					case 'h':
						hour, _ = strconv.Atoi(value[:len(value)-1])
					case 'm':
						mini, _ = strconv.Atoi(value[:len(value)-1])
					case 's':
						sec, _ = strconv.Atoi(value[:len(value)-1])
					}
				}
				modifyTime := time.Now().AddDate(year, month, day)
				modifyTime = modifyTime.Add(time.Hour*time.Duration(hour) + time.Minute*time.Duration(mini) +
					time.Second*time.Duration(sec))
				err := exec.Command("date", "-s", modifyTime.Format("01/02/2006 15:04:05.999999999")).Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Print("Set system time to:")
					fmt.Println(modifyTime.String())
				}
			}
		case "d":
			{
				modifyTime := time.Now().Add(time.Second * time.Duration(*tsSecond))
				err := exec.Command("date", "-s", modifyTime.Format("01/02/2006 15:04:05.999999999")).Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Print("Set system time to:")
					fmt.Println(modifyTime.String())
				}
			}
		}
	})
}
