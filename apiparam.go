package types

type Pagination struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type ListParams struct {
	Since int64 `json:"since"`
	Size  int   `json:"size"`
	Desc  bool  `json:"desc"`
}
