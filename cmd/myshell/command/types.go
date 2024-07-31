package command

type Command string

const (
	Exit Command = "exit"
	Echo Command = "echo"
	Type Command = "type"
)

func (c Command) String() string {
	return string(c)
}

func IsValidCommand(command string) bool {
	switch Command(command) {
	case Exit, Echo, Type:
		return true
	}
	return false
}
