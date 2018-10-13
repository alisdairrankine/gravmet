package gravmet

type Entity interface {
	Tick(*World, uint)

	Renderable() Renderable
}
