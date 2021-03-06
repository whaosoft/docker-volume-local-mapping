package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/volume"
)

const socketAddress = "/run/docker/plugins/local-mapping.sock"

const stateDir = "/mnt/state/"
const rootMountPoint = "/mnt/root/"

func main() {
	driver, _ := newLocalPersistDriver(rootMountPoint, stateDir)

	handler := volume.NewHandler(driver)
	logrus.Infof("Listening on %s", socketAddress)
	fmt.Println(handler.ServeUnix(socketAddress, 0))
}
