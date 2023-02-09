package proc

type IPFSProcess struct {
	*ProcessManager
}

func NewIPFSProcess() *ProcessManager {
	// return &IPFSProcess{
	// ProcessManager: &ProcessManager{},
	// }
	return nil
}

func (ip *IPFSProcess) Init() error {
	return nil
}
