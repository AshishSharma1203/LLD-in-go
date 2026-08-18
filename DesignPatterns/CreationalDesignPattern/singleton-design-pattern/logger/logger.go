package logger

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Logger is the exported API clients depend on.
//
// Go vs Java (production enforcement):
//
//	Java: private constructor → outside code cannot `new Logger()`.
//	Go: no private constructors. Instead:
//	  1) keep the concrete type unexported (`logger`)
//	  2) export an interface (`Logger`)
//	  3) only expose GetInstance() as the access point
//	Other packages cannot write `logger{}` — the type name is not visible.
//	They also cannot write `Logger{}` because Logger is an interface.
type Logger interface {
	Log(message string)
}

// // logger is the real singleton type (unexported).
type logger struct{}

var instance *logger

func (l *logger) Log(message string) {
	fmt.Println("Log:", message)
}

// // GetInstance returns the single Logger (lazy, not thread-safe yet).
// //
// // Prefer GetInstance over NewLogger here: "Get" means return the existing
// // shared instance (create once if missing). New* would imply a fresh object.
func GetInstance() Logger {
	if instance == nil {
		instance = &logger{}
	}
	return instance
}

// thread safe logger

var (
	threadSafeInstance atomic.Pointer[logger]
	mu                 sync.Mutex
)

func GetThreadSafeInstance() Logger {
	if p := threadSafeInstance.Load(); p != nil {
		return p
	}

	mu.Lock()
	defer mu.Unlock()

	if p := threadSafeInstance.Load(); p != nil {
		return p
	}
	l := &logger{}
	threadSafeInstance.Store(l)
	return l
}

var (
	onceInstance *logger
	once         sync.Once
)

func GetInstanceOnce() Logger {
	once.Do(func() {
		onceInstance = &logger{}
	})
	return onceInstance
}
