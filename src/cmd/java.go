package main

import (
	"jvm/src/class_file"
	"jvm/src/utils"
)

func main() {
	// for hello world
	classBytes := utils.ReadFile("/Users/e-postovalov/GolandProjects/jvm/src/test/src/resources/Hello.class")
	class, err := class_file.ParseClassFile(classBytes)
	if err != nil {
		panic(err)
	}
	print(class.ThisClass)
}
