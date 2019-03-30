package elevatorDirections

type Direction int

const (
	UP Direction = iota
	DOWN
	STAY
)

func (d Direction) ToString() string {
	if d == UP {
		return "UP"
	} else if d == DOWN {
		return "DOWN"
	} else if d == STAY {
		return "STAY"
	}
	return "NO STATUS"
}
