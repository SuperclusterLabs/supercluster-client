package proc

import (
	"os"
	"os/exec"

	"github.com/SuperclusterLabs/supercluster-client/util"
)

// TODO: unit tests!!!
type IPFSProcess struct {
	*ProcessManager
}

var _ ManagedProcess = (*IPFSProcess)(nil)

var kuboPath string = util.GetConfDir() + "/kubo/ipfs"

func NewIPFSProcess() (ManagedProcess, error) {
	if _, err := os.Stat(kuboPath); err != nil {
		return nil, err
	}
	return &IPFSProcess{
		NewProcessManager(kuboPath, []string{"daemon"}, util.GetConfDir()+"/kubo.log"),
	}, nil
}

func (ip *IPFSProcess) Init() error {
	// skip ipfs init if dir already exists
	hDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	_, err = os.Stat(hDir + "/.ipfs")
	if os.IsNotExist(err) {
		cmd := exec.Command(kuboPath, []string{"daemon"}...)
		return cmd.Wait()
	}
	if err != nil {
		return err
	}

	// run any outstanding migrations
	cmd := exec.Command(kuboPath, []string{"repo", "migrate"}...)
	return cmd.Wait()
}
