package domain

import (
	"testing"
	"time"
)

func Test_isBirthDay(t *testing.T) {
	type args struct {
		data time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Is BirthDay",
			args: args{data: time.Now()},
			want: true,
		},
		{
			name: "Isn't BirthDay",
			args: args{data: time.Now().AddDate(0, 0, -1)},
			want: false,
		},
		{
			name: "When BirthDay is undefined",
			args: args{data: time.Time{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBirthDay(tt.args.data); got != tt.want {
				t.Errorf("isBirthDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBlackFriday(t *testing.T) {
	type args struct {
		data time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Is BlackFriday",
			args: args{data: time.Now()},
			want: true,
		},
		{
			name: "Isn't BlackFriday",
			args: args{data: time.Now().AddDate(0, 0, -1)},
			want: false,
		},
		{
			name: "When BlackFriday is undefined",
			args: args{data: time.Time{}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBlackFriday(tt.args.data); got != tt.want {
				t.Errorf("isBlackFriday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDiscountPercentual(t *testing.T) {
	type args struct {
		blackFriday time.Time
		birthDay    time.Time
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "When Is BlackFriday and Is BirthDay then 10% discount (MAX)",
			args: args{
				blackFriday: time.Now(),
				birthDay:    time.Now()},
			want: 0.1,
		},
		{
			name: "When Is BlackFriday then 10%",
			args: args{
				blackFriday: time.Now(),
				birthDay:    time.Now().AddDate(0, 0, -1)},
			want: 0.1,
		},
		{
			name: "When Is BirthDay then 5%",
			args: args{
				blackFriday: time.Now().AddDate(0, 0, -1),
				birthDay:    time.Now()},
			want: 0.05,
		},
		{
			name: "When Isn't BlackFriday and Isn't BirthDay then no discount",
			args: args{
				blackFriday: time.Now().AddDate(0, 0, -1),
				birthDay:    time.Now().AddDate(0, 0, -1)},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDiscountPercentual(tt.args.blackFriday, tt.args.birthDay); got != tt.want {
				t.Errorf("GetDiscountPercentual() = %v, want %v", got, tt.want)
			}
		})
	}
}
