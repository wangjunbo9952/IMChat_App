package rabbit

import (
	"github.com/streadway/amqp"
	"log"
	"sync"
)

var pool *ConnectionPool

type ConnectionPool struct {
	connections []*amqp.Connection
	channels    []*amqp.Channel
	mu          sync.Mutex
	maxSize     int
}

// 创建新的连接池
func NewConnectionPool(maxSize int) {
	pool = &ConnectionPool{
		connections: make([]*amqp.Connection, 0, maxSize),
		channels:    make([]*amqp.Channel, 0, maxSize),
		maxSize:     maxSize,
	}
	return
}

// 获取连接
func (p *ConnectionPool) GetConnection() *amqp.Connection {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.connections) > 0 {
		conn := p.connections[len(p.connections)-1]
		p.connections = p.connections[:len(p.connections)-1]
		return conn
	}

	// 如果没有可用连接，创建一个新的连接
	conn, err := amqp.Dial("amqp://guest:guest@122.51.195.185:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	return conn
}

// 获取通道
func (p *ConnectionPool) GetChannel(conn *amqp.Connection) *amqp.Channel {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.channels) > 0 {
		ch := p.channels[len(p.channels)-1]
		p.channels = p.channels[:len(p.channels)-1]
		return ch
	}

	// 如果没有可用通道，创建一个新的通道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %s", err)
	}
	return ch
}

// 归还连接
func (p *ConnectionPool) ReturnConnection(conn *amqp.Connection) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.connections) < p.maxSize {
		p.connections = append(p.connections, conn)
	} else {
		conn.Close() // 超过最大连接数，关闭连接
	}
}

// 归还通道
func (p *ConnectionPool) ReturnChannel(ch *amqp.Channel) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.channels) < p.maxSize {
		p.channels = append(p.channels, ch)
	} else {
		ch.Close() // 超过最大通道数，关闭通道
	}
}
