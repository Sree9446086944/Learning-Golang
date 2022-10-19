package main

import (
	"fmt"
	"time"
)

//go mod init mytime
//https://pkg.go.dev/time#pkg-index

func main() {
	fmt.Println("Welcome to time package in golang.")
	presentTime := time.Now() //2022-10-19 02:24:59.6441101 -0500 CDT m=+0.003585101
	fmt.Println(presentTime)

	//format
	// for format , always use "01-02-2006 15:04:05 Monday" time for parsing any time to format

	fmt.Println(presentTime.Format("01-02-2006"))        //10-19-2022
	fmt.Println(presentTime.Format("01-02-2006 Mon"))    //10-19-2022 Wed
	fmt.Println(presentTime.Format("01-02-2006 Monday")) //10-19-2022 Wednesday

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //10-19-2022 02:28:56 Wednesday

	//create time from values

	//(year,month-time.Month type,day,hr,min,sec,nanaosec,location-*time.Location type)
	createdDate := time.Date(2020, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)                             //2020-08-12 23:23:00 +0000 UTC
	fmt.Println(createdDate.Format("01-02-2006 Monday")) //08-12-2020 Wednesday

	//go env -> gives all env variables
	// GOOS=windows -> this env gives the type os based on system
	// this is needed for creating executable for windows/linux,mac etc
	// using GOOS env we can build for different os
	//go build -> will build executable using go tools , this command automatically finds main.go file and other files and compile it

	//change GOOS value for build app based on os
	// go build -> builds based on system
	// GOOS="darwin" go build  -> builds app executable for mac os
	// GOOS="linux" go build -> builds app executable for linux os
	//then in cmd , .\mytime.exe
	// if error in cmd, use gitbash to execute the command

}
