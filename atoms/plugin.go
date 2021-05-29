package atoms

import (
	"context"

	"atomizer.io/engine"
)

// Ensure compile time adherence to interface
var _ engine.Atom = (*Plugin)(nil)

type Plugin struct{
	
}

func (t *Plugin) Process(
	ctx context.Context,
	conductor engine.Conductor,
	electron *engine.Electron,
) (result []byte, err error) {

	return result, err
}
