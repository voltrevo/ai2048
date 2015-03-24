package ai2048

type Direction int

const (
    Up Direction = iota
    Down
    Left
    Right
)

func (d Direction) String() string {
    switch d {
        case Up:
            return "Up"
        case Down:
            return "Down"
        case Left:
            return "Left"
        case Right:
            return "Right"
    }

    return "Invalid Direction"
}
