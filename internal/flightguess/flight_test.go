package flightguess

import (
	"reflect"
	"testing"
)

func TestFlight_positions(t *testing.T) {
	type fields struct {
		head Position
		body Body
	}
	tests := []struct {
		name   string
		fields fields
		want   []Position
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Flight{
				head: tt.fields.head,
				body: tt.fields.body,
			}
			if got := f.positions(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("positions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_isBodyOf(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		f *Flight
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := p.isBodyOf(tt.args.f); got != tt.want {
				t.Errorf("isBodyOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_isHeadOf(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		f *Flight
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := p.isHeadOf(tt.args.f); got != tt.want {
				t.Errorf("isHeadOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPosition_isPartOf(t *testing.T) {
	type fields struct {
		x int
		y int
	}
	type args struct {
		f *Flight
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "测试A",
			fields: fields{
				x: 3,
				y: 5,
			},
			args: args{
				f: &Flight{
					head: Position{3, 5},
				},
			},
			want: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Position{
				x: tt.fields.x,
				y: tt.fields.y,
			}
			if got := p.isPartOf(tt.args.f); got != tt.want {
				t.Errorf("isPartOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filter(t *testing.T) {
	type args struct {
		flights   FlightGroup
		predicate func(flight *Flight) bool
	}
	tests := []struct {
		name string
		args args
		want []*Flight
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.flights.filter(tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
