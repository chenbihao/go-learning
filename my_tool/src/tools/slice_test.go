package tools

import "testing"

func TestInSliceInt(t *testing.T) {
	type args struct {
		items []int
		item  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{{
		name: "test1",
		args: args{
			items: []int{1, 2, 3, 4, 5},
			item:  1,
		},
		want: true,
	}, {
		name: "test2",
		args: args{
			items: []int{1, 2, 3, 4, 5},
			item:  6,
		},
		want: false,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InSliceInt(tt.args.items, tt.args.item); got != tt.want {
				t.Errorf("InSliceInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInSliceStr(t *testing.T) {
	type args struct {
		items []string
		item  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				items: []string{"123", "asd", "cc1", "1234"},
				item:  "cc1",
			},
			want: true,
		}, {
			name: "test2",
			args: args{
				items: []string{"123", "asd", "cc1", "1234"},
				item:  "cc12",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InSliceStr(tt.args.items, tt.args.item); got != tt.want {
				t.Errorf("InSliceStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
