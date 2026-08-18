package main

import (
	"fmt"

	"github.com/ashishsharma1203/lld/singleton-design-pattern/logger"
)

func main() {
	logger1 := logger.GetInstance()
	logger1.Log("This is a log message")

	logger2 := logger.GetInstance()
	logger2.Log("another log message")

	fmt.Printf("same instance? %v\n", logger1 == logger2)

	// Lesson: other packages cannot construct the concrete singleton.
	// These do NOT compile (left here as comments on purpose):
	//   _ = logger.logger{}     // unexported type
	//   _ = logger.Logger{}     // Logger is an interface, not a struct
	// Only GetInstance() can hand out the shared instance.
}
