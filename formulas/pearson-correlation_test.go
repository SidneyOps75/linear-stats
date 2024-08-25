package formula

import "testing"

func TestPearsonCorrelation(t *testing.T) {
	type args struct {
		x []float64
		y []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Perfect positive correlation",
			args: args{
				x: []float64{1, 2, 3, 4, 5},
				y: []float64{2, 4, 6, 8, 10},
			},
			want: 1, // Pearson correlation of 1 (perfect positive correlation)
		},
		{
			name: "Perfect negative correlation",
			args: args{
				x: []float64{1, 2, 3, 4, 5},
				y: []float64{10, 8, 6, 4, 2},
			},
			want: -1, // Pearson correlation of -1 (perfect negative correlation)
		},
		{
			name: "No correlation",
			args: args{
				x: []float64{1, 2, 3, 4, 5},
				y: []float64{5, 5, 5, 5, 5},
			},
			want: 0, // No correlation since y is constant
		},
		{
			name: "Moderate positive correlation",
			args: args{
				x: []float64{1, 2, 3, 4, 5},
				y: []float64{1, 2, 4, 3, 5},
			},
			want: 0.9, // Approximate correlation of 0.8 (moderate positive correlation)
		},
		{
			name: "Single data point",
			args: args{
				x: []float64{1},
				y: []float64{1},
			},
			want: 0, // Single data point results in undefined correlation, treated as 0
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PearsonCorrelation(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("PearsonCorrelation() = %v, want %v", got, tt.want)
			}
		})
	}
}
