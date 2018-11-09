package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/kr/pty"
	"golang.org/x/crypto/ssh/terminal"
)

func test() error {
	// Create arbitrary command.
	c := exec.Command("docker", "run", "-it", "--rm", "alpine", "sh")

	// Start the command with a pty.
	ptmx, err := pty.Start(c)
	if err != nil {
		return err
	}
	// Make sure to close the pty at the end.
	defer func() { _ = ptmx.Close() }() // Best effort.

	// Handle pty size.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGWINCH)
	go func() {
		for range ch {
			if err := pty.InheritSize(os.Stdin, ptmx); err != nil {
				log.Printf("error resizing pty: %s", err)
			} else {
				s, _ := pty.GetsizeFull(ptmx)
				log.Printf("resizing pty %v\n", s)
				ws := &pty.Winsize{Rows: s.Rows, Cols: s.Cols - 10}
				pty.Setsize(ptmx, ws)
				log.Printf("resize to %v", ws)
			}
		}
	}()
	ch <- syscall.SIGWINCH // Initial resize.

	wssource, err := pty.GetsizeFull(os.Stdin)
	log.Printf("WS source %v", wssource)

	// Set stdin in raw mode.
	oldState, err := terminal.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = terminal.Restore(int(os.Stdin.Fd()), oldState) }() // Best effort.

	// Copy stdin to the pty and the pty to stdout.
	go func() { _, _ = io.Copy(ptmx, os.Stdin) }()
	_, _ = io.Copy(os.Stdout, ptmx)

	return nil
}

func main() {
	if err := test(); err != nil {
		log.Fatal(err)
	}
}
