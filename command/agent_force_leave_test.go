package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestAgentForceLeaveCommand_Implements(t *testing.T) {
	var _ cli.Command = &AgentForceLeaveCommand{}
}
