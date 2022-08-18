package one

import "fmt"

var Global int = 1

type ICommand interface {
	Execute() error
}

type StartCommand struct{}

func NewStartCommand() *StartCommand {
	return &StartCommand{}
}

func (c *StartCommand) Execute() error {
	fmt.Println("game start")
	return nil
}

type ArchiveCommand struct{}

func NewArchiveCommand() *ArchiveCommand {
	return &ArchiveCommand{}
}

func (c *ArchiveCommand) Execute() error {
	fmt.Println("game archive")
	return nil
}
