package weekday

import (
	"fmt"
	"time"
)

// Find は指定した年月日から入力日付と曜日を返します
func Find(year, month, day int) (string, string) {
	// 指定された日付を作成
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	// 入力日付（文字列）を作成
	inputDate := fmt.Sprintf("%04d/%02d/%02d", year, month, day)

	// 曜日を取得して文字列に変換
	weekday := date.Weekday().String()

	return inputDate, weekday
}