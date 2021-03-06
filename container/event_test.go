package container

import (
	"testing"

	"github.com/samalba/dockerclient"
	"github.com/stretchr/testify/assert"
)

func TestIsCreatedByTugbot_True(t *testing.T) {
	attributes := map[string]string{TugbotTest: "true", TugbotCreatedFrom: "aabb"}
	created := IsCreatedByTugbot(
		&dockerclient.Event{Actor: dockerclient.Actor{Attributes: attributes}})
	assert.True(t, created)
}

func TestIsCreatedByTugbot_False(t *testing.T) {
	attributes := map[string]string{TugbotTest: "true"}
	created := IsCreatedByTugbot(
		&dockerclient.Event{Actor: dockerclient.Actor{Attributes: attributes}})
	assert.False(t, created)
}

func TestIsCreatedByTugbot_NoAttributes(t *testing.T) {
	created := IsCreatedByTugbot(&dockerclient.Event{})
	assert.False(t, created)
}

func TestIsSwarmTask_True(t *testing.T) {
	attributes := map[string]string{TugbotTest: "true", TugbotCreatedFrom: "aabb", SwarmTaskID: "a123b"}
	created := IsSwarmTask(
		&dockerclient.Event{Actor: dockerclient.Actor{Attributes: attributes}})
	assert.True(t, created)
}

func TestIsSwarmTask_False(t *testing.T) {
	attributes := map[string]string{TugbotTest: "true", TugbotCreatedFrom: "aabb"}
	created := IsSwarmTask(
		&dockerclient.Event{Actor: dockerclient.Actor{Attributes: attributes}})
	assert.False(t, created)
}

func TestIsSwarmTask_NoAttributes(t *testing.T) {
	created := IsSwarmTask(&dockerclient.Event{})
	assert.False(t, created)
}
