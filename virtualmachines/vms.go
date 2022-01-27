package virtualmachines

import (
	"fmt"
	"os/exec"
)

/*
commands needed
prlctl start DockerLinux
prlctl status DockerLinux
prlctl exec DockerLinux uptime
prlctl stop DockerLinux

*/

const VM_NAME = "DockerLinux"

type ParallelsVM struct {
	Name string
}

type VM interface {
}

func cmdRun(name string, arg ...string) (int, string, error) {
	cmd := exec.Command(name, arg...)
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

func StartVM() {

	rc, msg, err := cmdRun("prlctl", "start", "DockerLinux")
	fmt.Println(msg, " ", rc, " ", err)

}

func VMStatus() {
	rc, msg, err := cmdRun("prlctl", "status", "DockerLinux")
	fmt.Println(msg, " ", rc, " ", err)
}

func CheckVM() {
	rc, msg, err := cmdRun("prlctl", "exec", "DockerLinux", "uptime")
	fmt.Println(msg, " ", rc, " ", err)
}

func StopVM() {

	rc, msg, err := cmdRun("prlctl", "stop", "DockerLinux")
	fmt.Println(msg, " ", rc, " ", err)
}
