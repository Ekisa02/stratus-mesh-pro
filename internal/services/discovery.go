package services

import (
	"sync"

	"github.com/example/stratus-mesh/internal/models"
)

// Discovery is a very small in-memory registry used for simulation and tests.
// In a real system this would be replaced by a gossip or service discovery mechanism.
type Discovery struct {
	mu sync.RWMutex
	peers map[uint]models.Node
}

func NewDiscovery() *Discovery {
	return &Discovery{peers: make(map[uint]models.Node)}
}

func (d *Discovery) Register(n models.Node) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.peers[n.ID] = n
}

func (d *Discovery) Unregister(id uint) {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.peers, id)
}

func (d *Discovery) List() []models.Node {
	d.mu.RLock()
	defer d.mu.RUnlock()
	out := make([]models.Node, 0, len(d.peers))
	for _, v := range d.peers {
		out = append(out, v)
	}
	return out
}
