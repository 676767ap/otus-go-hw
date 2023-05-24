package services

import "testing"

func TestMaxUCB1(t *testing.T) {
	tests := []struct {
		name string
		args []*BannerForCalc
		want int
	}{
		{
			name: "all banners have no show events, pick first banner",
			args: []*BannerForCalc{
				&BannerForCalc{BannerId: 1, Shown: 0, Clicks: 0},
				&BannerForCalc{BannerId: 2, Shown: 0, Clicks: 0},
				&BannerForCalc{BannerId: 3, Shown: 0, Clicks: 0},
			},
			want: 1,
		},
		{
			name: "one banner has show events, but not clicks, pick second banner",
			args: []*BannerForCalc{
				&BannerForCalc{BannerId: 1, Shown: 2, Clicks: 0},
				&BannerForCalc{BannerId: 2, Shown: 0, Clicks: 0},
				&BannerForCalc{BannerId: 3, Shown: 0, Clicks: 0},
			},
			want: 2,
		},
		{
			name: "two banners have show events, but not clicks, pick third banner",
			args: []*BannerForCalc{
				&BannerForCalc{BannerId: 1, Shown: 2, Clicks: 0},
				&BannerForCalc{BannerId: 2, Shown: 2, Clicks: 0},
				&BannerForCalc{BannerId: 3, Shown: 0, Clicks: 0},
			},
			want: 3,
		},
		{
			name: "all banners have different amount of show events, but not clicks, pick third banner",
			args: []*BannerForCalc{
				&BannerForCalc{BannerId: 1, Shown: 3, Clicks: 0},
				&BannerForCalc{BannerId: 2, Shown: 4, Clicks: 0},
				&BannerForCalc{BannerId: 3, Shown: 2, Clicks: 0},
			},
			want: 3,
		},
		{
			name: "one banner has clicks, pick first banner",
			args: []*BannerForCalc{
				&BannerForCalc{BannerId: 1, Shown: 5, Clicks: 1},
				&BannerForCalc{BannerId: 2, Shown: 4, Clicks: 0},
				&BannerForCalc{BannerId: 3, Shown: 4, Clicks: 0},
			},
			want: 1,
		},
		{
			name: "one banner has clicks, but too many show events, pick second banner",
			args: []*BannerForCalc{
				&BannerForCalc{BannerId: 1, Shown: 6, Clicks: 1},
				&BannerForCalc{BannerId: 2, Shown: 4, Clicks: 0},
				&BannerForCalc{BannerId: 3, Shown: 4, Clicks: 0},
			},
			want: 2,
		},
		{
			name: "all banners have clicks, pick second banner",
			args: []*BannerForCalc{
				&BannerForCalc{BannerId: 1, Shown: 7, Clicks: 2},
				&BannerForCalc{BannerId: 2, Shown: 4, Clicks: 1},
				&BannerForCalc{BannerId: 3, Shown: 4, Clicks: 0},
			},
			want: 2,
		},
		{
			name: "all banners have clicks, pick third banner",
			args: []*BannerForCalc{
				&BannerForCalc{BannerId: 1, Shown: 7, Clicks: 2},
				&BannerForCalc{BannerId: 2, Shown: 5, Clicks: 1},
				&BannerForCalc{BannerId: 3, Shown: 4, Clicks: 1},
			},
			want: 3,
		},
		{
			name: "one banner has many show events and clicks",
			args: []*BannerForCalc{
				&BannerForCalc{BannerId: 1, Shown: 16000, Clicks: 799},
				&BannerForCalc{BannerId: 2, Shown: 9000, Clicks: 59},
				&BannerForCalc{BannerId: 3, Shown: 3000, Clicks: 9},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxUCB1(tt.args); got != tt.want {
				t.Errorf("MaxUCB1() = %v, want %v", got, tt.want)
			}
		})
	}
}
