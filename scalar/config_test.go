package scalar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigMatch(t *testing.T) {
	config := Config{
		Env:       []string{"FT", "Staging"},
		Resources: []string{"FT"},
	}

	assert.True(t, config.matchJob("FT", []string{"FT"}))
	assert.False(t, config.matchJob("Production", []string{"Production"}))
	assert.False(t, config.matchJob("FT", []string{"Firefox"}))
	assert.False(t, config.matchJob("", []string{}))
}

func TestConfigMatchWhenNoEnvironmentOrResources(t *testing.T) {
	config := Config{
		Env:       []string{},
		Resources: []string{},
	}

	assert.True(t, config.matchJob("", []string{}))
}

func TestConfigMatchWhenOnlyNoEnvironment(t *testing.T) {
	config := Config{
		Env:       []string{},
		Resources: []string{"packer", "terraform"},
	}

	assert.True(t, config.matchJob("", []string{"packer"}))
	assert.False(t, config.matchJob("", []string{"docker"}))
}

func TestConfigAgentMatch(t *testing.T) {
	config := Config{
		Env:       []string{"FT", "Staging"},
		Resources: []string{"FT"},
	}

	agentEnv := []string{"FT", "Staging"}
	agentResources := []string{"FT"}

	assert.True(t, config.matchAgent(agentEnv, agentResources))
}

func TestConfigAgentMatchWhenNoEnvironment(t *testing.T) {
	config := Config{
		Env:       []string{},
		Resources: []string{"FT"},
	}

	agentEnv := []string{}
	agentResources := []string{"FT"}

	assert.True(t, config.matchAgent(agentEnv, agentResources))
}

func TestConfigAgentMatchWhenNoEnvironmentOrResource(t *testing.T) {
	config := Config{
		Env:       []string{},
		Resources: []string{},
	}

	agentEnv := []string{}
	agentResources := []string{}

	assert.True(t, config.matchAgent(agentEnv, agentResources))
}

func TestConfigAgentMatchWhenOnlyEnvironment(t *testing.T) {
	config := Config{
		Env:       []string{"FT"},
		Resources: []string{},
	}

	agentEnv := []string{"FT"}
	agentResources := []string{}

	assert.True(t, config.matchAgent(agentEnv, agentResources))

	agentEnv = []string{"FT"}
	agentResources = []string{"Staging"}
	assert.False(t, config.matchAgent(agentEnv, agentResources))
}
