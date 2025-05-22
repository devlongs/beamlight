package node

import (
	"context"
	"fmt"
	"sync"
)

// Service is the interface that any service in the node must implement.
type Service interface {
	// Start starts the service.
	Start(ctx context.Context) error
	// Stop stops the service.
	Stop() error
	// Name returns the service name.
	Name() string
}

// Node represents the beamlight node.
type Node struct {
	services []Service
	ctx      context.Context
	cancel   context.CancelFunc
	wg       sync.WaitGroup
	mu       sync.Mutex
	started  bool
}

// New creates a new Node instance.
func New() *Node {
	ctx, cancel := context.WithCancel(context.Background())
	return &Node{
		ctx:    ctx,
		cancel: cancel,
	}
}

// RegisterService registers a service to the node.
func (n *Node) RegisterService(service Service) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.started {
		panic(fmt.Sprintf("cannot register service %s after node has started", service.Name()))
	}

	n.services = append(n.services, service)
}

// Start starts all registered services.
func (n *Node) Start() error {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.started {
		return fmt.Errorf("node already started")
	}

	fmt.Println("Starting beamlight node with", len(n.services), "services")

	for _, service := range n.services {
		serviceName := service.Name()
		fmt.Printf("Starting service: %s\n", serviceName)

		n.wg.Add(1)
		go func(s Service) {
			defer n.wg.Done()
			if err := s.Start(n.ctx); err != nil {
				fmt.Printf("Error starting service %s: %v\n", s.Name(), err)
			}
		}(service)
	}

	n.started = true
	return nil
}

// Stop stops all registered services.
func (n *Node) Stop() error {
	n.mu.Lock()
	defer n.mu.Unlock()

	if !n.started {
		return fmt.Errorf("node not started")
	}

	fmt.Println("Stopping beamlight node...")

	// Cancel context to signal all services to stop
	n.cancel()

	// Stop all services in reverse order
	for i := len(n.services) - 1; i >= 0; i-- {
		service := n.services[i]
		fmt.Printf("Stopping service: %s\n", service.Name())
		if err := service.Stop(); err != nil {
			fmt.Printf("Error stopping service %s: %v\n", service.Name(), err)
		}
	}

	// Wait for all goroutines to finish
	n.wg.Wait()
	n.started = false

	fmt.Println("All services stopped")
	return nil
}

// Wait blocks until all services have stopped.
func (n *Node) Wait() {
	n.wg.Wait()
}
