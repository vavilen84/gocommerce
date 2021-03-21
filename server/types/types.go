package types

type ValidationErrors map[Field]error

type Scenario string

type ValidationMap map[Scenario]ValidationRules

type Rules string

type Field string

type ValidationRules map[Field]Rules
