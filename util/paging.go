package util

type PagingParams struct {
	LowerId   int64 `form:"lowerId"`
	UpperId   int64 `form:"upperId"`
	StartId   int64 `form:"startId"`
	PageSize  int64 `form:"pageSize"`
	PageIndex int64 `form:"pageIndex"`
	Reverse   bool  `form:"reverse"`
//	Count     int64
}

func (pp *PagingParams) Process(minStartId int64, defaultPageSize int64, maxPageSize int64) {
	if pp.StartId < minStartId {
		pp.StartId = minStartId
	}

	if pp.PageSize <= 0 {
		pp.PageSize = defaultPageSize
	} else if pp.PageSize > maxPageSize && maxPageSize > 0 {
		pp.PageSize = maxPageSize
	}
}

func (pp *PagingParams) Adjust(count int64, defaultPageSize int64) {

	if pp.PageSize <= 0 {
		pp.PageSize = defaultPageSize
	}

	if count == 0 {
		pp.PageIndex = 1
		return
	}

	if pp.PageIndex <= 0 {
		pp.PageIndex = 1
	}

	if (pp.PageIndex-1)*pp.PageSize >= count {
		pp.PageIndex = count / pp.PageSize
		if pp.PageIndex*pp.PageSize < count {
			pp.PageIndex += 1
		}
	}
}
