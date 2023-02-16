package proc

import (
	"os"
	"os/exec"

	"github.com/SuperclusterLabs/supercluster-client/util"

	"github.com/google/uuid"
	ma "github.com/multiformats/go-multiaddr"
)

type IPFSClusterProcess struct {
	id       uuid.UUID
	svcPort  string
	httpPort string
	secret   string

	*ProcessManager
}

var _ ManagedProcess = (*IPFSClusterProcess)(nil)

var clSvcPath string = util.GetConfDir() + "/ipfs-cluster/ipfs-cluster-service"
var clCtlPath string = util.GetConfDir() + "/ipfs-cluster/ipfs-cluster-ctl"
var clsDirPath string = util.GetConfDir() + "/clusters"

func NewIPFSClusterProcess(id uuid.UUID, svcPort, httpPort, secret string) (ManagedProcess, error) {
	if _, err := os.Stat(clCtlPath); err != nil {
		return nil, err
	}
	if _, err := os.Stat(clSvcPath); err != nil {
		return nil, err
	}

	return &IPFSClusterProcess{
		id:             id,
		ProcessManager: NewProcessManager(clSvcPath, []string{"-c", id.String(), "daemon"}, ""),
	}, nil
}

func (icp *IPFSClusterProcess) Init() error {
	// check if dir already exists, if not the daemon should start without issues
	if _, err := os.Stat(clsDirPath + "/" + icp.id.String()); err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		// initialize dir for this cluster
		cmd := exec.Command(kuboPath, []string{"-c", icp.id.String(), "init", "--randomports"}...)
		return cmd.Run()

		// TODO: update config with details provided
		// if icp.svcPort != "" && icp.httpPort != "" && icp.secret != "" {
		// ...
		// }
	}

	return nil
}

func (icp *IPFSClusterProcess) AddPeer(m ma.Multiaddr) {
	// TODO: add peer to peerstore
}
