#!/bin/bash

if [ $(id -u) != "0" ];then
	echo "ERROR,Please use su or sudo su to switch to root user and execute this script^_^"
	exit 1
fi

if [ -d "./fyne-cross" ];then
    rm -rf ./fyne-cross
fi


os_all='linux windows'
# arch_all='386 amd64 arm arm64 mips64 mips64le mips mipsle riscv64'
arch_all='386 amd64'

export GOFLAGS=-buildvcs=false

for os in $os_all; do
    for arch in $arch_all; do
        echo "package: OS:${os},arch: ${arch}"
        dockerimg='fyneio/fyne-cross-images:linux'
        if [ "x${os}" = x"windows" ];then
            dockerimg='fyneio/fyne-cross-images:1.1.0-windows'
        fi
        fyne-cross ${os} -app-id net.hlinfo.gouidemo -arch ${arch} -app-version 1.0.1 -image ${dockerimg} -icon ./assets/static/logo.png 
        echo ""
        echo ""
    done
done


