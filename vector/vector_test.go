package vector

import (
	"math"
	"math/rand"
	"testing"
)

func TestNormalize(t *testing.T) {
	tests := []struct {
		input, output Vector2
	}{
		{Vector2{0.0, 5.2}, Vector2{0.0, 1.0}},
		{Vector2{3.0, 4.0}, Vector2{0.6, 0.8}},
		{Vector2{0.3, 0.4}, Vector2{0.6, 0.8}},
		{Vector2{999, 0.0}, Vector2{1.0, 0.0}},
	}

	// Test for cases above
	for _, test := range tests {
		if actual := test.input.Normalize(); actual != test.output {
			t.Errorf("expected %v but got %v", test.output, actual)
		}
	}

	// Test for NaN case
	if actual := (Vector2{0.0, 0.0}).Normalize(); !math.IsNaN(actual.X) || !math.IsNaN(actual.Y) {
		t.Errorf("expected %v but got %v", Vector2{math.NaN(), math.NaN()}, actual)
	}
}

const jinyang = "JIIIIINNNN YAAAAAAAAAAANNNNNGGGGGG" // Do you like octopus?

func TestString(t *testing.T) {
	in := Vector2{145.152, 146.561}
	out := "(145.152, 146.561)"

	if actual := in.String(); actual != out {
		t.Errorf("expected %s but got %s", out, actual)
	}
}

func TestStringerOverride(t *testing.T) {
	o := StringerOverride
	StringerOverride = func(Vector2) string {
		return jinyang
	}

	if actual := (Vector2{}).String(); actual != jinyang {
		t.Errorf(
			"StringerOverride fail, expected %s but got %s",
			jinyang,
			actual,
		)
	}
	StringerOverride = o
}

func BenchmarkNormalize(b *testing.B) {
	vectors := make([]Vector2, 500, 500)
	for i, _ := range vectors {
		vectors[i] = randomVector()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = vectors[i%len(vectors)].Normalize()
	}
}

func randomVector() Vector2 {
	return Vector2{
		(rand.Float64() - 0.5) * 2000000,
		(rand.Float64() - 0.5) * 2000000,
	}
}
