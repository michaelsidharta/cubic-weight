package service

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/michaelsidharta/cubic-weight/entity"
	"github.com/michaelsidharta/cubic-weight/external"
)

func TestCalculator_GetAverage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		category string
	}
	tests := []struct {
		name    string
		args    args
		init    func() external.IAPI
		want    float64
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				category: "A",
			},
			init: func() external.IAPI {
				m := external.NewMockIAPI(ctrl)
				m.EXPECT().
					Get(gomock.Any(), "/api/products/1").
					Return(entity.APIResponse{
						Objects: []entity.Object{
							entity.Object{
								Category: "A",
								Size: entity.Size{
									Length: 100,
									Width:  200,
									Height: 200,
								},
							},
						},
						Next: "",
					}, nil).
					Times(1)
				return m
			},
			want:    1000.0,
			wantErr: false,
		}, {
			name: "success empty",
			args: args{
				category: "B",
			},
			init: func() external.IAPI {
				m := external.NewMockIAPI(ctrl)
				m.EXPECT().
					Get(gomock.Any(), "/api/products/1").
					Return(entity.APIResponse{
						Objects: []entity.Object{
							entity.Object{
								Category: "A",
								Size: entity.Size{
									Length: 100,
									Width:  200,
									Height: 200,
								},
							},
						},
						Next: "",
					}, nil).
					Times(1)
				return m
			},
			want:    0.0,
			wantErr: false,
		}, {
			name: "error no result",
			args: args{
				category: "A",
			},
			init: func() external.IAPI {
				m := external.NewMockIAPI(ctrl)
				m.EXPECT().
					Get(gomock.Any(), "/api/products/1").
					Return(entity.APIResponse{}, errors.New("BOOM")).
					Times(1)
				return m
			},
			want:    0.0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := InitCalculator(tt.init())
			got, err := c.GetAverage(tt.args.category)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculator.GetAverage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Calculator.GetAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
