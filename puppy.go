package puppy

import (
	"github.com/vsawant1608/dog"
	"github.com/vsawant1608/dog/cat"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "woof! woof! woof!"
}

func MixedFunction() string {
	return "hj" + dog.WhenGrown("Hola") + cat.Fights()
}
