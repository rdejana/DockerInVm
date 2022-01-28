package main

import (
	"DockerVm/containers"
	"fmt"
)

func main() {

	cc := containers.LoadContainerConfig()
	fmt.Println(cc)
	vm := containers.CreateVm(cc)
	vm.StartAndWaitVm()

	/*
		//name of the VM will be DockerLinux
		//ok, some work is needed to collect the response and info from the command
		fmt.Println("Welcome to the machine...")
		virtualmachines.StartVM()
		fmt.Println("VMStarted")
		virtualmachines.VMStatus()
		time.Sleep(60 * time.Second)
		virtualmachines.VMStatus()
		virtualmachines.CheckVM()
		//	time.Sleep(30 * time.Second)
		//	virtualmachines.CheckVM()
		//	virtualmachines.StartVM()
		fmt.Println("starting shutdown")
		virtualmachines.StopVM()
	*/

}
