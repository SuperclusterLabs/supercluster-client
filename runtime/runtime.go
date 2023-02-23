package runtime

import (
	"errors"
	"log"
	"os"
	"sort"

	"github.com/SuperclusterLabs/supercluster-client/db"
	"github.com/SuperclusterLabs/supercluster-client/proc"
	"github.com/SuperclusterLabs/supercluster-client/util"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

type SuperclusterRuntime struct {
	// IPFS daemon process
	IPFSDaemon proc.ManagedProcess

	// pool of cluster pinning services
	// TODO: process pool logic
	clusterCtlProcs map[string]*proc.IPFSClusterProcess
	// metadata DB for users and clusters
	AppDB db.SuperclusterDB

	*gin.Engine
}

func NewSuperclusterRuntime(ipfs proc.ManagedProcess, db db.SuperclusterDB) SuperclusterRuntime {
	return SuperclusterRuntime{
		IPFSDaemon:      ipfs,
		clusterCtlProcs: make(map[string]*proc.IPFSClusterProcess),
		AppDB:           db,
		Engine:          gin.Default(),
	}
}

var GlobalRuntime SuperclusterRuntime

func (r *SuperclusterRuntime) Init() {
	// TODO: come up with a strategy for preferred clusters.
	// Picking 10 most recent clusters for now
	cd := util.GetConfDir() + "/clusters"
	cs, err := os.ReadDir(cd)
	if err != nil {
		panic(err)
	}
	sort.Slice(cs, func(i, j int) bool {
		ii, _ := cs[i].Info()
		ji, _ := cs[j].Info()
		return ii.ModTime().Unix() > ji.ModTime().Unix()
	})
	if len(cs) > 10 {
		cs = cs[:10]
	}

	// start pinning services
	for _, c := range cs {
		u, err := uuid.Parse(c.Name())
		if err != nil {
			log.Println("Skipping cluster service for non-uuid " + c.Name())
			continue
		}
		icp, err := proc.NewHostIPFSClusterProcess(u)
		if err != nil {
			log.Println("Skipping cluster service: Incorrect install")
			continue
		}
		err = icp.Init()
		if err != nil {
			log.Println("Skipping cluster service: Init error: " + err.Error())
			continue
		}
		err = icp.Start()
		if err != nil {
			log.Println("Skipping cluster service: Start error: " + err.Error())
			continue
		}

		// TODO: separation of concerns?
		r.AddProcess(c.Name(), icp)
	}
}

// TODO: should this check if the process is running?
func (r *SuperclusterRuntime) AddProcess(clusterId string, p *proc.IPFSClusterProcess) error {
	// TODO: abstract
	if len(r.clusterCtlProcs) == 10 {
		return errors.New("Max clusters reached")
	}
	r.clusterCtlProcs[clusterId] = p
	return nil
}

func (r *SuperclusterRuntime) GetProcess(clusterId string) (*proc.IPFSClusterProcess, error) {
	if p, ok := r.clusterCtlProcs[clusterId]; ok {
		return p, nil
	}
	return nil, errors.New("No cluster running with that ID")
}

// TODO: should this check if the process is ended?
func (r *SuperclusterRuntime) RemoveProcess(clusterId string, p *proc.IPFSClusterProcess) {
	delete(r.clusterCtlProcs, clusterId)
}
