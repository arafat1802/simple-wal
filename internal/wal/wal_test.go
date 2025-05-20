package wal

import(
	"testing",
	"simple-wal/internal/wal"
)

func TestAdd(t *testing.T){
	tests := []struct {
		name      string
		a,b       int
		expected  int
	}(
			{name: "positive numbers", a: 2, b: 3, expected: 5}
			{name: "negative numbers", a: -2, b: -3, expected: -5}
			{name: "zero and positive", a: 0, b: 3, expected: 3}
			{name: "zero and zero", a: 0, b: 0, expected: 0}
			{name: "positive and negative", a: 2, b: -3, expected: -1}
	)

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T)  {
			result := wal.Add(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("Add(%d %d): " got %d, want %d, tt.a, tt.b, result, tt.expected)
			}
		})
	}

}