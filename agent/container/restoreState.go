package container

import (
	"os"
	"time"

	"github.com/subutai-io/agent/config"
	"github.com/subutai-io/agent/db"
	"github.com/subutai-io/agent/lib/container"
	"github.com/subutai-io/agent/log"
)

// temporary function to provide backward compatibility with old approach
// Need to remove it in next release
func compat() {
	for _, name := range container.Containers() {
		if _, err := os.Stat(config.Agent.LxcPrefix + name + "/.start"); !os.IsNotExist(err) {
			os.Remove(config.Agent.LxcPrefix + name + "/.start")
			container.AddMetadata(name, map[string]string{"state": "RUNNING"})
		}
	}
}

// StateRestore checks container state and starting or stopping containers if required.
func StateRestore(canRestore *bool) {
	compat()

	bolt, err := db.New()
	log.Check(log.WarnLevel, "Opening database", err)
	active := bolt.ContainerByKey("state", "RUNNING")
	log.Check(log.WarnLevel, "Closing database", bolt.Close())

	for _, v := range active {
		if !*canRestore {
			return
		}
		if container.State(v) != "RUNNING" {
			log.Debug("Starting container " + v)
			startErr := container.Start(v)
			for i := 0; i < 5 && startErr != nil; i++ {
				if !*canRestore {
					return
				}
				log.Debug("Retrying container " + v + " start")
				time.Sleep(time.Second * time.Duration(5+i))
				startErr = container.Start(v)
			}
			if startErr != nil {
				container.AddMetadata(v, map[string]string{"state": "STOPPED"})
			}
		}
	}
}
