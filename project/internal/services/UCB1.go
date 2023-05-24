package services

import (
	"math"
)

type BannerForCalc struct {
	BannerId int
	Shown    int
	Clicks   int
}

func MaxUCB1(bannersForCalc []*BannerForCalc) int {
	var (
		banner     BannerForCalc
		totalShown float64
		rating     float64 = -1
	)

	for _, b := range bannersForCalc {
		shownItem := b.Shown
		if shownItem == 0 {
			shownItem = 1
		}
		totalShown += float64(shownItem)
	}
	for _, b := range bannersForCalc {
		if r := calculate(float64(b.Clicks), float64(b.Shown), totalShown); r > rating {
			banner = *b
			rating = r
		}
	}
	return banner.BannerId
}

func calculate(clicks, shown, totalShown float64) float64 {
	if shown == 0 {
		shown = 1
	}
	return clicks/shown + math.Sqrt(2*math.Log(totalShown)/shown)
}
