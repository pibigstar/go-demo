package main

import (
	"sync"
	"strings"
	"os/exec"
	"log"
)

func exec_cmd(cmd string, wg *sync.WaitGroup)  {

	parts := strings.Fields(cmd)

	out, err := exec.Command(parts[0], parts[1]).Output()
	if err!=nil {
		log.Println("exec command failed:",err)
	}

	log.Println("out:",string(out))

	wg.Done()
}


func main() {

	cmds := []string{"echo Hello","echo World"}
	wg := new(sync.WaitGroup)
	for _,cmd := range cmds{
		wg.Add(1)
		go exec_cmd(cmd,wg)
	}

	wg.Wait()

}
