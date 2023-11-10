package main

import "os"

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (l ConsoleLogger) Log(message string) {
	println(message)
}

type FileLogger struct{}

func (l FileLogger) Log(message string) {
	// write to file
	err := os.WriteFile("log.txt", []byte(message), 0644)
	if err != nil {
		panic(err)
	}

}

type MyService struct {
	logger Logger
}

func (s *MyService) DoSomething() {
	s.logger.Log("Doing something")
}

func main() {
	//using console logger
	consoleLogger := ConsoleLogger{}
	service1 := MyService{logger: consoleLogger}
	service1.DoSomething()

	//using file logger
	fileLogger := FileLogger{}
	service2 := MyService{logger: fileLogger}
	service2.DoSomething()

}
