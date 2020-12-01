package net

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strings"
)

type ActiveLanInterface struct {
}

func (this *ActiveLanInterface) Active() {
	inters, _ := net.Interfaces()
	// Run 和 Start只能用一个
	// Run starts the specified command and waits for it to complete.
	for _, value := range inters {
		fmt.Println(value)
	}
}

type ActiveWanInterface struct {

}

func (this *ActiveWanInterface) Active() {
	commandString := "ifconfig"
	cmd := exec.Command("sh", "-c", commandString)
	stdout, err := cmd.StdoutPipe()
	defer stdout.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 1024*5)
	stdout.Read(buf)

	str := string(buf)

	res := strings.Contains(str, "flags")

	fmt.Println(res)
}