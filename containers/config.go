package containers

import "fmt"

type ContainerConfig struct {
	NonVMContext string
	VmContext    string
	VmName       string
	VmType       string
}

func (cc ContainerConfig) String() string {
	return fmt.Sprintf("[ \"%s\", \"%s\", \"%s\", \"%s\" ]", cc.NonVMContext, cc.VmContext, cc.VmName, cc.VmType)
}

//at some point, load from a file...
func LoadContainerConfig() ContainerConfig {

	return ContainerConfig{
		NonVMContext: "default",
		VmContext:    "docker-vm",
		VmName:       "DockerLinux",
		VmType:       "parallels",
	}
}
