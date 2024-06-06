package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

func main() {
	// Allow the current process to lock memory for eBPF resources.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatalf("Failed to remove memlock: %s", err)
	}

	// Load the compiled BPF program.
	spec, err := ebpf.LoadCollectionSpec("hello_bpf.o")
	if err != nil {
		log.Fatalf("Failed to load collection spec: %s", err)
	}

	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		log.Fatalf("Failed to create collection: %s", err)
	}
	defer coll.Close()

	prog := coll.Programs["hello"]
	if prog == nil {
		log.Fatalf("Program 'hello' not found")
	}

	// Attach the eBPF program to the kprobe
	kp, err := link.Kprobe("sys_execve", prog, nil)
	if err != nil {
		log.Fatalf("Failed to attach kprobe: %s", err)
	}
	defer kp.Close()

	// Set up a channel to receive signals for graceful termination
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Open the trace pipe
	tracePipe, err := os.Open("/sys/kernel/debug/tracing/trace_pipe")
	if err != nil {
		log.Fatalf("Failed to open trace pipe: %s", err)
	}
	defer tracePipe.Close()

	// Read and print trace messages
	reader := bufio.NewReader(tracePipe)
	go func() {
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				log.Printf("Failed to read from trace pipe: %s", err)
				continue
			}
			fmt.Print(string(line))
		}
	}()

	// Wait for a termination signal
	<-sigs
	fmt.Println("Exiting...")
}
