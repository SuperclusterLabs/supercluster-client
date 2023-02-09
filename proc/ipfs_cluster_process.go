package proc

type IPFSClusterProcess struct {
	*ProcessManager
}

func NewIPFSClusterProcess() *ProcessManager {
	// return &IPFSClusterProcess{
	// ProcessManager: &ProcessManager{},
	// }
	return nil
}

func (ip *IPFSClusterProcess) Init() error {
	// ipfs-cluster-service -c test1 init
	return nil
}
