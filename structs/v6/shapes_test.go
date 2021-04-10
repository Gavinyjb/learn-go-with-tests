package structs

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}

}
func BenchmarkPerimeter(b *testing.B) {
	rectangle := Rectangle{12.0, 6.0}
	for i := 0; i < b.N; i++ {
		Perimeter(rectangle)
	}
}
func BenchmarkRectangle_Area(b *testing.B) {
	rectangle := Rectangle{12.0, 6.0}
	for i := 0; i < b.N; i++ {
		rectangle.Area()
	}
}
