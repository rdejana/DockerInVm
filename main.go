package main

import (
	"DockerVm/virtualmachines"
	"fmt"
	"time"
)

func main() {
	//name of the VM will be DockerLinu
	//ok, some work is needed to collect the response and info from the command
	fmt.Println("Welcome to the machine...")
	virtualmachines.StartVM()
	virtualmachines.VMStatus()
	virtualmachines.CheckVM()
	time.Sleep(30 * time.Second)
	virtualmachines.CheckVM()
	virtualmachines.StartVM()
	virtualmachines.StopVM()

}
