package command

type Command string

const (
	Exit Command = "exit"
	Echo Command = "echo"
)

func (c Command) String() string {
	return string(c)
}
