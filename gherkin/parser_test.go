package gherkin

import (
	"testing"

	"github.com/gauntlt/gauntlt-go/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestParseSingleScenario(t *testing.T) {
	s := `
# a line comment
@tag1 @tag2
Attack: Basic attack

  # a comment embedded in the description
  A multiline description of the Attack
  that can contain any text like
  Rules:
  - a
  - b

  @tag3 @tag4
  Scenario: The title of the scenario
    Given an attack tool is installed
    When the tool is run
    And some other condition
    Then there should not be vulnerabilities

# End of file comment with some empty lines below

`
	attacks, err := Parse(s)

	assert.NoError(t, err)
	assert.Equal(t, 1, len(attacks))
	assert.Equal(t, "Basic attack", attacks[0].Title)
	assert.Equal(t, "@tag1", attacks[0].Tags[0])
	assert.Equal(t, "@tag2", attacks[0].Tags[1])
	assert.Equal(t, 1, len(attacks[0].Scenarios))
	assert.Equal(t, "A multiline description of the Attack\nthat can contain any text like\nRules:\n- a\n- b", attacks[0].Description)
	assert.Equal(t, "The title of the scenario", attacks[0].Scenarios[0].Title)
	assert.Equal(t, 4, len(attacks[0].Scenarios[0].Steps))
	assert.Equal(t, "@tag3", attacks[0].Scenarios[0].Tags[0])
	assert.Equal(t, "@tag4", attacks[0].Scenarios[0].Tags[1])
	assert.Equal(t, StepType("Given"), attacks[0].Scenarios[0].Steps[0].Type)
	assert.Equal(t, "an attack tool is installed", attacks[0].Scenarios[0].Steps[0].Text)
	assert.Equal(t, StepType("When"), attacks[0].Scenarios[0].Steps[1].Type)
	assert.Equal(t, "the tool is run", attacks[0].Scenarios[0].Steps[1].Text)
	assert.Equal(t, StepType("And"), attacks[0].Scenarios[0].Steps[2].Type)
	assert.Equal(t, "some other condition", attacks[0].Scenarios[0].Steps[2].Text)
	assert.Equal(t, StepType("Then"), attacks[0].Scenarios[0].Steps[3].Type)
	assert.Equal(t, "there should not be vulnerabilities", attacks[0].Scenarios[0].Steps[3].Text)
}

func TestMultipleScenarios(t *testing.T) {
	s := `
Attack: Parsing multiple scenarios
  Scenario: Scenario name here
    Given some precondition
    Then something happens

  Scenario: The second scenario
    Given a condition
    Then something happens
`
	attacks, err := Parse(s)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(attacks[0].Scenarios))
	assert.Equal(t, 2, len(attacks[0].Scenarios[0].Steps))
	assert.Equal(t, 2, len(attacks[0].Scenarios[1].Steps))
}

func TestBackgroundAndScenarios(t *testing.T) {
	s := `
Attack: Parsing multiple scenarios
  @tag_on_background
  Background:
    Given there is some background

	Scenario: Scenario name here
    Given some precondition
    And then something else
    Then something happens

  Scenario: Another scenario
    Given another precondition
    Then something happens
`
	// Scenario: The scenario is a thing
	//   When there is an action
	//   Then something happens

	// Scenario: Another thing
	//   When there is another action
	//   Then something else happens
	attacks, err := Parse(s)
	assert.NoError(t, err)
	// assert.Equal(t, "@tag_on_background", attacks[0].Background.Tags[0])
	// assert.Equal(t, "there is some background", attacks[0].Background.Steps[0].Text)

	if len(attacks[0].Scenarios) != 2 {
		t.Error("Got length of", len(attacks[0].Scenarios))
	}
	assert.Equal(t, 2, len(attacks[0].Scenarios))
}

// func TestMultipleFeatures(t *testing.T) {
// 	s := `
// Feature: Feature 1
//   Scenario: Scenario name here
//     Given some precondition
//     Then something happens

// Feature: Feature 2
//   Scenario: Another scenario
//     Given another precondition
//     Then something happens
// `
// 	f, err := Parse(s)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 2, len(f))
// 	assert.Equal(t, "Feature 1", f[0].Title)
// 	assert.Equal(t, "Feature 2", f[1].Title)
// 	assert.Equal(t, 1, len(f[0].Scenarios))
// 	assert.Equal(t, 1, len(f[1].Scenarios))

// }

// func TestTagParsing(t *testing.T) {
// 	f, err := Parse("@tag1   @tag2@tag3\nFeature: Tag parsing")
// 	assert.NoError(t, err)
// 	assert.Equal(t, 2, len(f[0].Tags))
// 	assert.Equal(t, "@tag1", f[0].Tags[0])
// 	assert.Equal(t, "@tag2@tag3", f[0].Tags[1])
// }

// func TestBacktrackingCommentsAtEnd(t *testing.T) {
// 	s := `
// Feature: Comments at end
//   Scenario: Scenario name here
//     Given some precondition
//     Then something happens
//   # comments here
// `
// 	_, err := Parse(s)

// 	assert.NoError(t, err)
// }

// func TestBacktrackingCommentsDontAffectIndent(t *testing.T) {
// 	s := `
// Feature: Comments at end
//   Scenario: Scenario name here
//     Given some precondition
//     Then something happens
// # comments here
//   Scenario: Another scenario
//     Given a step
// `
// 	f, err := Parse(s)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 2, len(f[0].Scenarios))
// }

// func TestScenarioOutlines(t *testing.T) {
// 	s := `
// Feature: Scenario outlines
//   Scenario Outline: Scenario 1
//     Given some value <foo>
//     Then some result <bar>

//     Examples:
//     | foo | bar |
//     | 1   | 2   |
//     | 3   | 4   |

//   Scenario: Scenario 2
//     Given some other scenario
// `
// 	f, err := Parse(s)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 2, len(f[0].Scenarios))
// 	assert.Equal(t, StringData("| foo | bar |\n| 1   | 2   |\n| 3   | 4   |"), f[0].Scenarios[0].Examples)
// }

// func TestStepArguments(t *testing.T) {
// 	s := `
// Feature: Step arguments
//   Scenario: Scenario 1
//     Given some data
//                           | 1   | 2   |
//                           | 3   | 4   |
//     And some docstring
//     """
//      hello
//      world
//     """
//     And some table
//     | 1 | 2 |
//     Then other text

//   Scenario: Scenario 2
//     Given some other scenario
// `
// 	f, err := Parse(s)
// 	assert.NoError(t, err)
// 	assert.Equal(t, 2, len(f[0].Scenarios))
// 	assert.Equal(t, StringData("| 1   | 2   |\n| 3   | 4   |"), f[0].Scenarios[0].Steps[0].Argument)
// 	assert.Equal(t, StringData(" hello\n world"), f[0].Scenarios[0].Steps[1].Argument)
// 	assert.Equal(t, StringData("| 1 | 2 |"), f[0].Scenarios[0].Steps[2].Argument)
// 	assert.Equal(t, "other text", f[0].Scenarios[0].Steps[3].Text)
// }

// func TestFailureNoFeature(t *testing.T) {
// 	_, err := Parse("")
// 	assert.EqualError(t, err, `parse error (<unknown>.feature:1): no features parsed.`)
// }

// func TestTagWithoutFeature(t *testing.T) {
// 	_, err := Parse("@tag")
// 	assert.EqualError(t, err, `parse error (<unknown>.feature:1): tags not applied to feature.`)
// }

// func TestFailureExpectingFeature(t *testing.T) {
// 	_, err := Parse("@tag\n@tag")
// 	assert.EqualError(t, err, `parse error (<unknown>.feature:2): expected "Feature:", found "@tag".`)
// }

// func TestFailureInvalidTag(t *testing.T) {
// 	_, err := Parse("@tag tag")
// 	assert.EqualError(t, err, `parse error (<unknown>.feature:1): invalid tag "tag".`)
// }

// func TestFailureDescriptionAfterTags(t *testing.T) {
// 	s := `
// Feature: Descriptions after tags
//   @tag1
//   Descriptions should not be allowed after tags

//   Scenario: Scenario name here
//     Given some precondition
//     Then something happens
// `
// 	_, err := Parse(s)
// 	assert.EqualError(t, err, `parse error (<unknown>.feature:4): illegal description text after tags.`)
// }

// func TestFailureDescriptionAfterScenario(t *testing.T) {
// 	s := `
// Feature: Descriptions after scenario
//   Scenario: Scenario name here
//     Given some precondition
//     Then something happens

//   Descriptions should not be allowed after scenario

//   Scenario: Another scenario
//     Given some precondition
//     Then something happens
// `
// 	_, err := Parse(s)
// 	assert.EqualError(t, err, `parse error (<unknown>.feature:7): illegal description text after scenario.`)
// }

// func TestFailureMultipleBackgrounds(t *testing.T) {
// 	s := `
// Feature: Multiple backgrounds
//   Background:
//     Given one

//   Background:
//     Given two
// `
// 	_, err := Parse(s)
// 	assert.EqualError(t, err, `parse error (<unknown>.feature:6): multiple backgrounds not allowed.`)
// }

// func TestFailureBackgroundAfterScenario(t *testing.T) {
// 	s := `
// Feature: Background after scenario
//   Scenario: Scenario name here
//     Given some precondition
//     Then something happens

//   Background:
//     Given it's after a scenario
// `
// 	_, err := Parse(s)
// 	assert.EqualError(t, err, `parse error (<unknown>.feature:7): illegal background after scenario.`)
// }

// func TestFailureInvalidStep(t *testing.T) {
// 	s := `
// Feature: Invalid steps
//   Scenario: Scenario name here
//     Invalid step
// `
// 	_, err := Parse(s)
// 	assert.EqualError(t, err, `parse error (<unknown>.feature:4): illegal step prefix "Invalid".`)
// }

// func TestFailureNoStepText(t *testing.T) {
// 	s := `
// Feature: No step text
//   Scenario: Scenario name here
//     Given
// `
// 	_, err := Parse(s)
// 	assert.EqualError(t, err, `parse error (<unknown>.feature:4): expected step text after "Given".`)
// }

// func TestFailureInvalidTagOnScenario(t *testing.T) {
// 	s := `
// Feature: Invalid tag on scenario
//   @invalid tags
//   Scenario:
//     Given a scenario
// `
// 	_, err := Parse(s)
// 	assert.EqualError(t, err, `parse error (<unknown>.feature:3): invalid tag "tags".`)

// }

// func TestFailureInvalidBackground(t *testing.T) {
// 	s := `
// Feature: Invalid background
//   Background:
//     Invalid step
//   Scenario: A scenario
//     Given a scenario
// `
// 	_, err := Parse(s)
// 	assert.EqualError(t, err, `parse error (<unknown>.feature:4): illegal step prefix "Invalid".`)
// }
