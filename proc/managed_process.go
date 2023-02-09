package proc

type ManagedProcess interface {
	Init() error
	Start() error
	Stop() error
}
