package injection

import "fmt"

// CharactersContainer is the root structure for the YAML array.
type CharactersContainer struct {
	Characters []Character `yaml:"characters"`
}
// Character represents each character in the YAML.
type Character struct {
	UUID            int               `yaml:"uuid"`
	Name            string            `yaml:"name"`
	NPC             bool              `yaml:"npc"`
	Age             int               `yaml:"age"`
	TemporaryStatus map[string]int    `yaml:"temporary_status"`
	CharacterStatus map[string]int    `yaml:"character_status"`
	Equipment       []int             `yaml:"equipment"`
	Inventory       []map[string]int  `yaml:"inventory"`
}
// Processor defines an interface for processing a character.
type Processor interface {
	Process(c Character)
}
// ConsoleProcessor implements the Processor interface.
type ConsoleProcessor struct{}

// Process prints the character to the console.
func (cp *ConsoleProcessor) Process(c Character) {
	fmt.Printf("Processing Character: %+v\n", c)
}