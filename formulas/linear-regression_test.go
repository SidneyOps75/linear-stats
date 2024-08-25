package formula

import "testing"

func TestLinearRegression(t *testing.T) {
	type args struct {
		x []float64
		y []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 float64
	}{
		{
			name: "Basic linear relationship",
			args: args{
				x: []float64{1, 2, 3, 4, 5},
				y: []float64{2, 4, 6, 8, 10},
			},
			want:  2, // Slope (m)
			want1: 0, // Intercept (b)
		},
		{
			name: "Positive slope and non-zero intercept",
			args: args{
				x: []float64{1, 2, 3, 4, 5},
				y: []float64{3, 5, 7, 9, 11},
			},
			want:  2, // Slope (m)
			want1: 1, // Intercept (b)
		},
		{
			name: "Negative slope",
			args: args{
				x: []float64{1, 2, 3, 4, 5},
				y: []float64{10, 8, 6, 4, 2},
			},
			want:  -2, // Slope (m)
			want1: 12, // Intercept (b)
		},
		{
			name: "Zero slope",
			args: args{
				x: []float64{1, 2, 3, 4, 5},
				y: []float64{5, 5, 5, 5, 5},
			},
			want:  0, // Slope (m)
			want1: 5, // Intercept (b)
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LinearRegression(tt.args.x, tt.args.y)
			if got != tt.want {
				t.Errorf("LinearRegression() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LinearRegression() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
