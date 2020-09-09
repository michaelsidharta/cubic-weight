package entity

import (
	"reflect"
	"testing"
)

func TestAPIResponse_FilterObjectByCategory(t *testing.T) {
	type fields struct {
		Objects []Object
		Next    string
	}
	type args struct {
		category string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []Object
	}{
		{
			name: "success",
			fields: fields{
				Objects: []Object{
					Object{Category: "A"},
					Object{Category: "A"},
					Object{Category: "B"},
					Object{Category: "C"},
					Object{Category: "D"},
					Object{Category: "A"},
				},
			},
			args: args{category: "A"},
			want: []Object{
				Object{Category: "A"},
				Object{Category: "A"},
				Object{Category: "A"},
			},
		}, {
			name: "success_no_data",
			fields: fields{
				Objects: []Object{
					Object{Category: "Z"},
					Object{Category: "X"},
					Object{Category: "B"},
					Object{Category: "C"},
					Object{Category: "D"},
					Object{Category: "F"},
				},
			},
			args: args{category: "A"},
			want: []Object{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := APIResponse{
				Objects: tt.fields.Objects,
				Next:    tt.fields.Next,
			}
			if got := a.FilterObjectByCategory(tt.args.category); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("APIResponse.FilterObjectByCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSize_CubicVolume(t *testing.T) {
	type fields struct {
		Length float64
		Width  float64
		Height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "success",
			fields: fields{
				Length: 40,
				Width:  30,
				Height: 20,
			},
			want: float64(6.000000000000001),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Size{
				Length: tt.fields.Length,
				Width:  tt.fields.Width,
				Height: tt.fields.Height,
			}
			if got := s.CubicWeight(); got != tt.want {
				t.Errorf("Size.CubicWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
