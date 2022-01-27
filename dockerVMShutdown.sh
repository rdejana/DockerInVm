#!/bin/sh

echo "Stopping Docker in a VM"
echo "Stopping VM"
prlctl stop DockerLinux
prlctl status DockerLinux
echo "Resetting context"
docker context use desktop-linux 
