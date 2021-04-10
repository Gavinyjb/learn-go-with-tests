package structs

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "矩形", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "圆形", shape: Circle{10}, want: 314.1592653589793},
		{name: "三角形", shape: Triangle{12, 6}, want: 36.0},
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
