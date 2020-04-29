package structmethods

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 5.0}
	got := rectangle.Perimeter()
	want := 30.0

	if got != want {
		t.Errorf("got %f instead  %f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rectangle", shape: Rectangle{10.0, 5.0}, want: 50.0},
		{name: "circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "triangle", shape: Triangle{2.0, 4.0}, want: 4.0},
	}

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			checkArea(t, test.shape, test.want)
		})
	}

}
