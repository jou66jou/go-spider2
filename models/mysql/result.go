package mysql

// 爬蟲結果
type DataResult struct {
	ID       int    `gorm:"column:id;AUTO_INCREMENT"`
	SourceID int    `gorm:"column:lottery_source_id"`           // 來源
	IssueNo  string `gorm:"column:issue_no"`                    // 期號
	PrizeNum string `gorm:"column:prize_num"`                   // 號碼
	IsSelect int    `gorm:"column:is_selected;type:tinyint(1)"` // 1:已經送出redis
}
