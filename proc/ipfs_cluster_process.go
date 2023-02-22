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
	config   *map[string]interface{}

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
		ProcessManager: NewProcessManager(clSvc, []string{"-c", clsDir + "/" + id.String(), "daemon"}, logDir+"/"+id.String()),
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
	cld := clsDir + "/" + icp.id.String()
	// check if dir already exists, if not the daemon should start without issues
	if _, err := os.Stat(cld); err != nil {
		if !os.IsNotExist(err) {
			return err
		}

		// initialize dir for this cluster
		cmd := exec.Command(clSvc, []string{"-c", cld, "init", "--randomports"}...)
		err = cmd.Run()
		if err != nil {
			return err
		}
	}

	// get generated config file for the cluster
	c, err := util.ReadJSONFile(cld + "/service.json")
	if err != nil {
		return err
	}
	icp.config = &c

	return nil
}

func (icp *IPFSClusterProcess) AddPeer(m ma.Multiaddr) {
	// TODO: add peer to peerstore
}

func (icp *IPFSClusterProcess) GetPort() (string, error) {
	apiData, ok := (*icp.config)["api"]
	if !ok {
		return "", util.ErrBadClConfig
	}

	// TODO: can we do better than this ugly wrangling? refer to IPFS cluster src to find out
	ipfsproxyData, ok := apiData.(map[string]interface{})["ipfsproxy"]
	if !ok {
		return "", util.ErrBadClConfig
	}

	listenAddr, ok := ipfsproxyData.(map[string]interface{})["listen_multiaddress"].(string)
	if !ok {
		return "", util.ErrBadClConfig
	}
	m, err := ma.NewMultiaddr(listenAddr)
	if err != nil {
		return "", err
	}

	port, err := m.ValueForProtocol(ma.ProtocolWithName("tcp").Code)
	if err != nil {
		return "", err
	}

	return port, nil
}
