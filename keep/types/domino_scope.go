package types

type ScopeBody struct {
	ApiName            string `json:"apiName"`
	Description        string `json:"description"`
	Icon               string `json:"icon"`
	IconName           string `json:"iconName"`
	IsActive           bool   `json:"isActive"`
	MaximumAccessLevel string `json:"maximumAccessLevel"`
	NSFPath            string `json:"nsfPath"`
	SchemaName         string `json:"schemaName"`
	Server             string `json:"server"`
	DominoBaseScope
}

type DominoBaseScope struct {
	Name      string   `json:"Name"`
	Form      string   `json:"Form"`
	Type      string   `json:"Type"`
	UpdatedBy []string `json:"$UpdatedBy"`
	Revisions string   `json:"$Revisions"`
	DominoBaseDocument
}
