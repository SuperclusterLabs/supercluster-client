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

var clSvc string = util.GetConfDir() + "/ipfs-cluster/ipfs-cluster-service"
var clCtl string = util.GetConfDir() + "/ipfs-cluster/ipfs-cluster-ctl"
var clsDir string = util.GetConfDir() + "/clusters"
var logDir string = util.GetConfDir() + "/logs"

func NewHostIPFSClusterProcess(id uuid.UUID) (*IPFSClusterProcess, error) {
	if _, err := os.Stat(clCtl); err != nil {
		return nil, err
	}
	if _, err := os.Stat(clSvc); err != nil {
		return nil, err
	}

	return &IPFSClusterProcess{
		id:             id,
		ProcessManager: NewProcessManager(clSvc, []string{"-c", logDir + "/" + id.String(), "daemon"}, logDir+"/"+id.String()),
	}, nil
}

func NewJoinIPFSClusterProcess(id uuid.UUID, svcPort, httpPort, secret string) (ManagedProcess, error) {
	if _, err := os.Stat(clCtl); err != nil {
		return nil, err
	}
	if _, err := os.Stat(clSvc); err != nil {
		return nil, err
	}

	return &IPFSClusterProcess{
		id:             id,
		ProcessManager: NewProcessManager(clSvc, []string{"-c", clsDir + "/" + id.String(), "daemon"}, logDir+"/"+id.String()),
	}, nil
}

func (icp *IPFSClusterProcess) Init() error {
	// check if dir already exists, if not the daemon should start without issues
	if _, err := os.Stat(clsDir + "/" + icp.id.String()); err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		// initialize dir for this cluster
		cmd := exec.Command(clSvc, []string{"-c", clsDir + "/" + icp.id.String(), "init", "--randomports"}...)
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
