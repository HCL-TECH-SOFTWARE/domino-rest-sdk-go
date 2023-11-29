package types

type ListViewBody struct {
	Name             string `json:"name"`
	SelectionFormula string `json:"selectionFormula"`
	Columns          []DesignColumnSimple
	DominoBaseListView
}

type DesignColumnSimple struct {
	Name                   string   `json:"name"`
	Title                  string   `json:"title"`
	Formula                string   `json:"formula"`
	SeparateMultipleValues bool     `json:"separateMultipleValues"`
	Sort                   SortType `json:"sort"`
	Position               int      `json:"position"`
}

type SortType struct {
	Ascending  string `json:"ascending" default:"ascending"`
	Descending string `json:"descending" default:"descending"`
	None       string `json:"none" default:"none"`
}

type DominoBaseListView struct {
	Type     ListType `json:"type"`
	Alias    []string `json:"alias"`
	IsFolder bool     `json:"isFolder"`
	Title    string   `json:"title"`
	UNID     string   `json:"unid"`
	NoteId   string   `json:"noteid"`
}

type ListType struct {
	FOLDER string `json:"folder" default:"folder"`
	VIEW   string `json:"view" default:"view"`
}

type GetListViewEntryOptions struct {
	ListViewEntryOptions
	Subscriber func() (map[string]interface{}, error)
}

type ListViewEntryOptions struct {
	Mode              string                 `json:"mode"`
	Meta              bool                   `json:"meta"`
	StartsWith        string                 `json:"startsWith"`
	PivotColumn       string                 `json:"pivotColumn"`
	MetaAdditional    bool                   `json:"metaAdditional"`
	Category          []string               `json:"category"`
	Column            string                 `json:"column"`
	DistinctDocuments bool                   `json:"distinctDocuments"`
	FTSearchQuery     string                 `json:"ftSearchQUery"`
	Count             int                    `json:"count"`
	UnreadOnly        bool                   `json:"unreadOnly"`
	KeyAllowPartial   bool                   `json:"keyAllowPartial"`
	Documents         bool                   `json:"documents"`
	Key               []string               `json:"key"`
	Direction         SortShort              `json:"direction"`
	Scope             ViewEntryScopes        `json:"scope"`
	RichTextAs        RichTextRepresentation `json:"richTextAs"`
	MarkRead          bool                   `json:"markRead"`
	MarkUnread        bool                   `json:"markUnread"`
	Start             int                    `json:"start"`
}

type SortShort struct {
	ASC  string `json:"asc" default:"asc"`
	DESC string `json:"desc" default:"desc"`
}

type ViewEntryScopes struct {
	ALL        string `json:"all" default:"all"`
	CATEGORIES string `json:"categories" default:"categories"`
	DOCUMENTS  string `json:"documents" default:"documents"`
}
