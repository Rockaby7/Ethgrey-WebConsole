package utils

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

func CommandWriteBack(ctx context.Context, cmd string, outChan chan string) error {
	cmdCtx := exec.CommandContext(ctx, "sh", "-c", cmd)
	stdout, err := cmdCtx.StdoutPipe()
	if err != nil {
		close(outChan)
		return err
	}

	go func() {
		reader := bufio.NewReader(stdout)
		for {
			select {
			case <-ctx.Done():
				if ctx.Err() != nil {
					fmt.Println("error: ", ctx.Err())
				}
				close(outChan)
				return
			default:
				readString, err := reader.ReadString('\n')
				if err != nil {
					close(outChan)
					return
				}

				readString = strings.TrimSpace(readString)
				if readString != "" {
					outChan <- readString
				}
			}
		}
	}()
	return cmdCtx.Start()
}

func Command(cmd string) (string, error) {
	output, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), err
}

func SignalFn() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	<-done
	fmt.Println("exiting")
	os.Exit(0)
}
