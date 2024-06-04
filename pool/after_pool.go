package main

import (
	"fmt"
	"sync"
)

// Simulates a database connection
type Connection struct {
	ID int
}

// Simulates a connection pool
type ConnectionPool struct {
	connections []*Connection
	mu          sync.Mutex
}

// Creates a new connection pool
func NewConnectionPool(size int) *ConnectionPool {
	pool := &ConnectionPool{}
	for i := 0; i < size; i++ {
		pool.connections = append(pool.connections, &Connection{ID: i})
	}
	return pool
}

// Acquires a connection from the pool
func (p *ConnectionPool) AcquireConnection() *Connection {
	p.mu.Lock()
	defer p.mu.Unlock()
	if len(p.connections) == 0 {
		return nil
	}
	conn := p.connections[0]
	p.connections = p.connections[1:]
	return conn
}

// Releases a connection back to the pool
func (p *ConnectionPool) ReleaseConnection(conn *Connection) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.connections = append(p.connections, conn)
}

func main() {
	// With pool
	pool := NewConnectionPool(3)
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			conn := pool.AcquireConnection()
			if conn != nil {
				defer pool.ReleaseConnection(conn)
				fmt.Printf("Using connection %d\n", conn.ID)
			} else {
				fmt.Printf("No connection available for %d\n", id)
			}
		}(i)
	}
	wg.Wait()
}
