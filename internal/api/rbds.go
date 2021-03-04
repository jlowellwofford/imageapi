package api

// API operations on rbd maps

import (
	"fmt"
	"sync"

	"github.com/bensallen/rbd/pkg/krbd"
	"github.com/jlowellwofford/imageapi/models"
)

type RbdsType struct {
	next  models.ID
	rbds  map[models.ID]*models.Rbd
	mutex *sync.Mutex
}

func (r *RbdsType) Init() {
	r.next = 1 // starting from 1 means 0 == unspecified
	r.rbds = make(map[models.ID]*models.Rbd)
	r.mutex = &sync.Mutex{}
}

func (r *RbdsType) List() (result []*models.Rbd) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	for _, m := range r.rbds {
		result = append(result, m)
	}
	return
}

func (r *RbdsType) Map(rbd *models.Rbd) (m *models.Rbd, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	// sanity check
	if len(rbd.Monitors) == 0 || *rbd.Pool == "" || *rbd.Image == "" || rbd.Options.Name == "" || rbd.Options.Secret == "" {
		return nil, fmt.Errorf("The following are required: 1 or more monitors, pool, image, options/name, options/secret")
	}
	w, err := krbd.RBDBusAddWriter()
	if err != nil {
		return nil, fmt.Errorf("krbd error: %v", err)
	}
	defer w.Close()

	// We allow this because we get free IPv4 format checking
	mons := []string{}
	for _, m := range rbd.Monitors {
		mons = append(mons, m.String())
	}

	i := krbd.Image{
		Monitors: mons,
		Pool:     *rbd.Pool,
		Image:    *rbd.Image,
		Snapshot: rbd.Snapshot,
		Options: &krbd.Options{
			ReadOnly:  rbd.Options.Ro,
			Name:      rbd.Options.Name,
			Secret:    rbd.Options.Secret,
			Namespace: rbd.Options.Namespace,
		},
	}
	// make sure ID doesn't already exist
	dev := krbd.Device{Image: i.Image, Pool: i.Pool, Namespace: i.Options.Namespace, Snapshot: i.Snapshot}

	if err := dev.Find(); err == nil {
		return nil, fmt.Errorf("rbd device already exists")
	}
	// map the rbd
	if err := i.Map(w); err != nil {
		return nil, fmt.Errorf("krbd error: %v", err)
	}

	// now go find our ID
	if err := dev.Find(); err != nil {
		return nil, fmt.Errorf("could not find device ID: %v", err)
	}
	rbd.DeviceID = dev.ID
	rbd.DeviceFile = dev.DevPath()
	rbd.ID = r.next
	r.next++
	r.rbds[rbd.ID] = rbd

	return rbd, err
}

func (r *RbdsType) Get(id models.ID) (m *models.Rbd, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	var ok bool
	if m, ok = r.rbds[id]; ok {
		return
	}
	return nil, ERRNOTFOUND
}

func (r *RbdsType) Unmap(id models.ID) (m *models.Rbd, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var rbd *models.Rbd
	var ok bool

	if rbd, ok = r.rbds[id]; !ok {
		return nil, ERRNOTFOUND
	}

	// should we be able to force this?
	if rbd.Refs > 0 {
		return nil, fmt.Errorf("device %d is in use, cannot unmap", id)
	}

	wc, err := krbd.RBDBusRemoveWriter()
	defer wc.Close()

	i := krbd.Image{
		DevID: int(rbd.DeviceID),
		Options: &krbd.Options{
			Force: rbd.Options.Force,
		},
	}

	if err := i.Unmap(wc); err != nil {
		return nil, fmt.Errorf("krbd error: %v", err)
	}
	// remove from our map
	delete(r.rbds, id)

	return rbd, nil
}

// add/subtract from ref counter
// silently fails if id doesn't exist
func (r *RbdsType) RefAdd(id models.ID, n int64) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if rbd, ok := r.rbds[id]; ok {
		rbd.Refs += n
	}
}
