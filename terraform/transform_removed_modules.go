package terraform

import (
	"log"

	"github.com/hashicorp/terraform/configs"
)

// RemovedModuleTransformer implements GraphTransformer to add nodes indicating
// when a module was removed from the configuration.
type RemovedModuleTransformer struct {
	Config *configs.Config // root node in the config tree
	State  *State
}

func (t *RemovedModuleTransformer) Transform(g *Graph) error {
	// nothing to remove if there's no state!
	if t.State == nil {
		return nil
	}

	for _, m := range t.State.Modules {
		c := t.Module.Child(m.Path[1:])
		if c != nil {
			continue
		}

		log.Printf("[DEBUG] module %s no longer in config\n", modulePrefixStr(m.Path))
		g.Add(&NodeModuleRemoved{PathValue: m.Path})
	}
	return nil
}
