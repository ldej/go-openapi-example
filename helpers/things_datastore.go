package helpers

import (
	"errors"
	"sync"
	"time"
)

var ErrNotFound = errors.New("not found")
var ErrEmptyUUID = errors.New("empty uuid")
var ErrAlreadyExists = errors.New("already exists")

type Thing struct {
	UUID    string
	Name    string
	Rank    int64
	Rating  float32
	Score   float64
	Type    string
	Created time.Time
}

func NewThingsDatastore() *ThingsDatastore {
	return &ThingsDatastore{
		lock:  sync.Mutex{},
		index: map[string]Thing{},
	}
}

type ThingsDatastore struct {
	lock  sync.Mutex
	index map[string]Thing
}

func (d *ThingsDatastore) StoreThing(t Thing) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if t.UUID == "" {
		return ErrEmptyUUID
	}
	if _, found := d.index[t.UUID]; found {
		return ErrAlreadyExists
	}
	d.index[t.UUID] = t
	return nil
}

func (d *ThingsDatastore) ListThings(offset int, limit int) ([]Thing, int) {
	d.lock.Lock()
	defer d.lock.Unlock()

	// NOTE: Let's assume maps are ordered for simplicity’s sake :)
	var results []Thing
	for _, thing := range d.index {
		results = append(results, thing)
		if len(results) == limit {
			break
		}
	}
	return results, len(d.index)
}

func (d *ThingsDatastore) GetThing(uuid string) (Thing, error) {
	d.lock.Lock()
	defer d.lock.Unlock()

	thing, found := d.index[uuid]
	if !found {
		return Thing{}, ErrNotFound
	}
	return thing, nil
}

func (d *ThingsDatastore) UpdateThing(t Thing) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	storedThing, found := d.index[t.UUID]
	if !found {
		return ErrNotFound
	}
	storedThing.Score = t.Score
	d.index[t.UUID] = storedThing
	return nil
}

func (d *ThingsDatastore) DeleteThing(uuid string) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if _, found := d.index[uuid]; !found {
		return ErrNotFound
	}
	delete(d.index, uuid)
	return nil
}
