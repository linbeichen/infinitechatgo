package dto

type PageRequest struct {
	PageNum  int32 `json:"pageNum" form:"pageNum"`
	PageSize int32 `json:"pageSize" form:"pageSize"`
}

type PageResponse[T any] struct {
	List        []T   `json:"list"`
	Total       int64 `json:"total"`
	PageSize    int64 `json:"pageSize"`
	PageNum     int64 `json:"pageNum"`
	Pages       int64 `json:"pages"`
	HasNext     bool  `json:"hasNext"`
	HasPrevious bool  `json:"hasPrevious"`
}

func NewPageResponseFromAll[T any](allData []T, pageNum, pageSize int) *PageResponse[T] {
	if len(allData) == 0 {
		return &PageResponse[T]{
			List:        []T{},
			Total:       0,
			PageSize:    int64(pageSize),
			PageNum:     int64(pageNum),
			Pages:       0,
			HasNext:     false,
			HasPrevious: false,
		}
	}

	total := int64(len(allData))
	pages := (total + int64(pageSize) - 1) / int64(pageSize)

	fromIndex := (pageNum - 1) * pageSize
	toIndex := min(fromIndex+pageSize, len(allData))

	var list []T
	if fromIndex < len(allData) {
		list = allData[fromIndex:toIndex]
	} else {
		list = []T{}
	}

	return &PageResponse[T]{
		List:        list,
		Total:       total,
		PageSize:    int64(pageSize),
		PageNum:     int64(pageNum),
		Pages:       pages,
		HasNext:     toIndex < len(allData),
		HasPrevious: fromIndex > 0,
	}
}

// 数据库分页用（已经在 SQL 里做了分页，直接传 total 和 当页使用）
func NewPageResponse[T any](list []T, total int64, pageNum, pageSize int) *PageResponse[T] {
	pages := (total + int64(pageSize) - 1) / int64(pageSize)

	return &PageResponse[T]{
		List:        list,
		Total:       total,
		PageSize:    int64(pageSize),
		PageNum:     int64(pageNum),
		Pages:       pages,
		HasNext:     pageNum < int(pages),
		HasPrevious: pageNum > 1,
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
