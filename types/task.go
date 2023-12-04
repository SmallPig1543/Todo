package types

type TaskCreateReq struct {
	Content string `json:"content" form:"content"`
	Title   string `json:"title" form:"title"`
}

type TaskUpdateReq struct {
	Id      uint   `json:"id" form:"id"`
	Status  int    `json:"status" form:"status"`
	Content string `json:"content" form:"content"`
	Title   string `json:"title" form:"title"`
}

type TaskShowReq struct {
	Id uint `json:"id" form:"id"`
}

type TaskListReq struct {
	Limit int `json:"limit" form:"limit"`
	Start int `json:"start" form:"start"`
}

type TaskSearchReq struct {
	Info   string `json:"info" form:"info"`
	Status int    `json:"status" form:"status"`
	Start  int    `json:"start" form:"start"`
}

type TaskDeleteReq struct {
	Id     uint `json:"id" form:"id"`
	Status int  `json:"status" form:"status"`
}

type TaskInfoResp struct {
	Id        uint   `json:"id,omitempty"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	View      int    `json:"view,omitempty"`
	Status    int    `json:"status"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}
