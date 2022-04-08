package elevator

type Building struct {
	Elevator Elevator
	Floors   []Floor
}

type Floor struct {
	Panel  ElevatorOutPanel
	Up     bool
	Down   bool
	Number int
}

func (f *Floor) CallUp() {
	if !f.Up {
		f.Panel.CallUp()
		f.Up = true
	}
}

func (f *Floor) CallDown() {
	if !f.Down {
		f.Panel.CallDown()
		f.Down = true
	}
}
