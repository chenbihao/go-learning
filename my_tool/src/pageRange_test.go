package tools

import (
	"reflect"
	"testing"
)

func TestParsePageRange(t *testing.T) {
	type args struct {
		pageRange string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "4",
			args:    args{pageRange: "4"},
			want:    []int{4},
			wantErr: false,
		}, {
			name:    "nil",
			args:    args{pageRange: ""},
			want:    nil,
			wantErr: true,
		}, {
			name:    "1,3",
			args:    args{pageRange: "1,3"},
			want:    []int{1, 3},
			wantErr: false,
		}, {
			name:    "1-2,4-5",
			args:    args{pageRange: "1-2,4-5"},
			want:    []int{1, 2, 4, 5},
			wantErr: false,
		}, {
			name:    "1-2,4-5",
			args:    args{pageRange: "1-2,4-5"},
			want:    []int{1, 2, 4, 5},
			wantErr: false,
		}, {
			name:    "1-5,7,10-15",
			args:    args{pageRange: "1-5,7,10-12"},
			want:    []int{1, 2, 3, 4, 5, 7, 10, 11, 12},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePageRange(tt.args.pageRange)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePageRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePageRange() got = %v, want %v", got, tt.want)
			}
		})
	}
}
