package common

func CalculateAmount(amount int64) (amount_f float64) {
	amount_f = float64(amount) / 100
	return
}

func FloadAmount2Int(amount_f float64) (amount int64) {
	amount = int64(amount_f * 100)
	return
}
