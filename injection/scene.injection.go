package injection

import (
	"context"
	"sync"
)

type SceneContainer struct {
	mu    sync.Mutex // Protect concurrent access to the scene
	Scene Scene
}

type Scene struct {
	Position            string           `yaml:"position"`
	Coordinate          string           `yaml:"coordinate"`
	Weather             string           `yaml:"weather"`
	PlayerPartyMember   []string         `yaml:"player_party_member"`
	OpponentPartyMember []string         `yaml:"opponent_party_member"`
	ItemNearby          []map[string]int `yaml:"item_nearby"`
}

type SceneManipulator interface {
	Change(ctx context.Context, d Scene)
}

// Change updates the Scene state in a thread-safe manner.
func (sp *SceneProcessor) Change(ctx context.Context, d Scene) {
	sp.container.mu.Lock()
	defer sp.container.mu.Unlock()

	// Update the scene with the new data.
	sp.container.Scene = d
}

type SceneProcessor struct{
	container *SceneContainer
}

func NewSceneProcessor(container *SceneContainer) *SceneProcessor {
	return &SceneProcessor{container: container}
}
