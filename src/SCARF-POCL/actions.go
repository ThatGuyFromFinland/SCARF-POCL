package main

import (
	"encoding/csv"
	"os"
	"strings"
	//"fmt"
)

type Action struct {
	name            string
	transformations string
	preconditions   string
}

func (this *Action) isResolverOf(trait string) bool {
	for _, transformation := range strings.Split(this.transformations, ";") {
		transformationParts := strings.Split(transformation, ":")
		transformationTrait := strings.Join(transformationParts[1:], ":")
		if trait == transformationTrait {
			return true
		}
	}
	return false
}

/* Obsolete?
func (this *Action) hasConsent(character Character) bool {
	for _, precondition := range strings.Split(this.preconditions, ";") {
		condition := strings.Split(precondition, ":")[1]
		if character.hasState(condition) {
			return true
		}
	}
	return false
}
*/

func NewActionList(path string) []Action {
	var actions []Action

	f, err := os.Open(path)
	if err != nil {
		return nil
	}

	defer f.Close() // this needs to be after the err check

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil
	}

	var new_action Action
	for key, value := range lines {
		if key == 0 {
			continue
		}

		new_action = Action{name: value[0], transformations: value[1], preconditions: value[2]}
		actions = append(actions, new_action)

	}

	//fmt.Printf("%v\n", actions)
	return actions
}
