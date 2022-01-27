#!/bin/sh

echo "Setting up Docker in a VM"
echo "Starting VM"
prlctl start DockerLinux
prlctl status DockerLinux
echo "Setting up context"
docker context use docker-vm


