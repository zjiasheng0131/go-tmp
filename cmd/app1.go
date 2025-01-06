package main

import "fmt"

type App struct {
	name    string
	servers []int
}

type Option func(a *App)

func NewApp(opts ...Option) *App {
	a := &App{}
	for _, opt := range opts {
		opt(a)
	}
	return a
}

func WithServer(servers ...int) Option {
	return func(a *App) {
		a.servers = servers
	}
}
func WithName(name string) Option {
	return func(a *App) {
		a.name = name
	}
}

func main() {
	fmt.Println(1111)
	a := NewApp(WithServer(123, 456), WithName("test-cloud"))
	fmt.Println(a)

}
