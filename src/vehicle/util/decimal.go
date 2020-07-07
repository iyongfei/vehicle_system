package util

import "github.com/shopspring/decimal"

// 加法
func DecimalAdd(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {

	return d1.Add(d2)
}

// 减法
func DecimalSub(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Sub(d2)
}

// 乘法
func DecimalMul(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Mul(d2)
}

// 除法
func DecimalDiv(d1 decimal.Decimal, d2 decimal.Decimal) decimal.Decimal {
	return d1.Div(d2)
}

// int
func DecimalInt(d decimal.Decimal) int64 {
	return d.IntPart()
}

// float
func DecimalFloat(d decimal.Decimal) float64 {
	f, exact := d.Float64()
	if !exact {
		return f
	}
	return 0
}
