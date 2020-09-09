package entity

import "github.com/michaelsidharta/cubic-weight/constant"

type APIResponse struct {
	Objects []Object `json:"objects"`
	Next    string   `json:"next"`
}

func (a APIResponse) FilterObjectByCategory(category string) []Object {
	res := make([]Object, 0, len(a.Objects))
	for _, obj := range a.Objects {
		if obj.Category == category {
			res = append(res, obj)
		}
	}
	return res
}

type Object struct {
	Category string  `json:"category"`
	Title    string  `json:"title"`
	Weight   float64 `json:"weight"`
	Size     Size    `json:"size"`
}

type Size struct {
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

func (s Size) CubicWeight() float64 {
	vol := (s.Height / 100) * (s.Length / 100) * (s.Width / 100)
	return constant.CubicWeightConversionFactor * vol
}
