package models

type (
	PaginateReq struct {
		Start string `query:"start" validate:"max=64"`
		Limit int    `query:"limit" validate:"required,min=2,max=10"`
	}

	PaginateRes struct {
		Date  any           `json:"date"`
		Limit int           `json:"limit"`
		Total int64         `json:"total"`
		First FirstPaginate `json:"first"`
		Next  NextPaginate  `json:"next"`
	}

	FirstPaginate struct {
		Href  string `json:"href"`
		Start string `jsin:"start"`
	}

	NextPaginate struct {
		Href string `json:"href"`
	}

	KafkaOffset struct {
		Offset int64 `json:"offset" bson:"offset"`
	}
)
