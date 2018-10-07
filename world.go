package gravmet

type World struct {
	entities  map[string]Entity
	positions map[string]uint
}

func NewWorld() *World {
	return &World{
		entities:  make(map[string]Entity),
		positions: make(map[string]uint),
	}
}

func (w *World) Tick(delta uint) {
	for _, entity := range w.entities {
		entity.Tick(w, delta)
	}
}

func (w *World) RemoveEntity(id string) {
	delete(w.entities, id)
}

func (w *World) AddEntity(e Entity) string {

	id := ""
	w.entities[id] = e
	return id
}
