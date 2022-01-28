package containers

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

//vm for parallels  still don't know if I want a struct

type ParallelsVm struct {
	Name string
}

func CreateVm(config ContainerConfig) ParallelsVm {
	return ParallelsVm{
		Name: config.VmName,
	}
}

/*
internal method to run commands
*/
func runParallelsCmd(arg ...string) (int, string, error) {
	cmd := exec.Command("prlctl", arg...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		rc := 255 //default
		if exitError, ok := err.(*exec.ExitError); ok {
			rc = exitError.ExitCode()
		}

		return rc, string(out), err
	}
	return 0, string(out), err
}

func (pvm ParallelsVm) IsRunning() (bool, error) {
	started, err := pvm.IsStarted()
	if err != nil {
		return false, err
	}
	if started == true {
		rc, msg, err := runParallelsCmd("exec", pvm.Name, "uptime")
		//if not booted, we'll get an error back
		//fmt.Println(rc, " ", msg, " ", err)
		if err == nil {
			if rc == 0 { //command ran...
				//fmt.Println(msg)
				return true, nil
			}

		}
		return false, errors.New(msg)

	}
	return false, nil
}

func (pvm ParallelsVm) internalRunning() bool { //this means booted...
	rc, msg, err := runParallelsCmd("exec", pvm.Name, "uptime")
	fmt.Println(msg, " ", rc, " ", err)
	return false
}

func (pvm ParallelsVm) IsStarted() (bool, error) { //this means it is powered own
	//rc 0 >> command ran
	// "exist running" -> running
	// "exist stopped" -> stopped

	rc, msg, err := runParallelsCmd("status", pvm.Name)

	if err == nil {
		if rc == 0 {
			if strings.Contains(msg, "stopped") {
				return false, nil
			} else if strings.Contains(msg, "running") {
				return true, nil
			} else {
				fmt.Println("Unknown return message: ", msg)
				return false, errors.New("Unknown status message: " + msg)
			}
		}
	}
	//we have an error
	return false, errors.New(msg)
	//fmt.Println(msg, " ", rc, " ", err)
	//return false,nil
}

func (pvm ParallelsVm) StartVM() error { //need a good return...

	running, err := pvm.IsStarted()
	if err != nil {
		return err
	}

	if running == true {
		fmt.Printf("VM \"%s\" is already started.", pvm.Name)
		return nil
	}

	rc, msg, err := runParallelsCmd("start", pvm.Name)
	if err != nil {
		//need better errors
		return errors.New(msg)
	}
	//already running
	//doesn't exit
	//other

	fmt.Println(msg, " ", rc, " ", err)
	return nil
}

func (pvm ParallelsVm) StartAndWaitVm() {
	pvm.StartVM()
	fmt.Print("VM is starting")
	//timeout would be nice
	keepSleeping := true
	for keepSleeping == true {
		running, _ := pvm.IsRunning()
		if running == true {
			break
		}
		fmt.Print(".")
		time.Sleep(15 * time.Second)
	}
	fmt.Println("\nVM started")

}
