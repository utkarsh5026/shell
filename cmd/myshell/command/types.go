package command

type Command string

const (
	Exit Command = "exit"
	Echo Command = "echo"
	Type Command = "type"
	PWD  Command = "pwd"
)

func (c Command) String() string {
	return string(c)
}

func IsValidCommand(command string) bool {
	switch Command(command) {
	case Exit, Echo, Type, PWD:
		return true
	}
	return false
}
