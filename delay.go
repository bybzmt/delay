/*
 * 延时1s启动程序
 */
package main

import (
	"os"
	"os/signal"
	"os/exec"
	"time"
	"log"
	"syscall"
)


func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Option Error! usage:", os.Args[0], "program")
	}

	time.Sleep(1 * time.Second)

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatalln(err.Error())
	}

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, os.Kill)

		for {
			err := cmd.Process.Signal(<-c)
			if err != nil {
				log.Fatalln(err.Error())
			}
		}
	}()

	p, err := cmd.Process.Wait()
	if err != nil {
		log.Fatalln(err.Error())
	}

	os.Exit(p.Sys().(syscall.WaitStatus).ExitStatus())
}
