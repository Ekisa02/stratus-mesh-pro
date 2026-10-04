package services

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"

	"github.com/example/stratus-mesh/internal/models"
)

// Replicator demonstrates a background job that periodically scans schedules
// and attempts to replicate tasks to peer nodes via the discovery registry.
type Replicator struct {
	db *gorm.DB
	discovery *Discovery
	interval time.Duration
}

func NewReplicator(db *gorm.DB, discovery *Discovery) *Replicator {
	return &Replicator{db: db, discovery: discovery, interval: 30 * time.Second}
}

func (r *Replicator) Run(ctx context.Context) {
	tick := time.NewTicker(r.interval)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Println("replicator stopping")
			return
		case <-tick.C:
			r.replicateOnce()
		}
	}
}

func (r *Replicator) replicateOnce() {
	var schedules []models.Schedule
	if err := r.db.Preload("Node").Find(&schedules).Error; err != nil {
		log.Printf("replicator: failed to load schedules: %v", err)
		return
	}
	peers := r.discovery.List()
	if len(peers) == 0 {
		log.Print("replicator: no peers discovered")
		return
	}
	for _, s := range schedules {
		// naive replication: log an intent to replicate to all peers except origin
		for _, p := range peers {
			if p.ID == s.NodeID {
				continue
			}
			log.Printf("replicator: would replicate task '%s' for schedule %d to peer %s (%s)", s.Task, s.ID, p.Name, p.Address)
		}
	}
}
