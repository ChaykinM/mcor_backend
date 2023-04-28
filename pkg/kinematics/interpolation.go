package kinematics

import (
	"log"

	"main.go/pkg/models"
	// "../models"
)

func (k *Kinematics) GetCubicSplines(time []float64, val []float64) []*models.CubicSplineData {

	// val[i] = a[i] + b[i]*(x-x[i]) + c[i]*(x-x[i])^2 + d[i]*(x-x[i])^3
	N := len(val) - 1 // количество участков для интерполирования

	var a []float64 = make([]float64, N+1)
	var b []float64 = make([]float64, N)
	var d []float64 = make([]float64, N)
	var h []float64 = make([]float64, N)
	var alpha []float64 = make([]float64, N)
	var c []float64 = make([]float64, N+1)
	var l []float64 = make([]float64, N+1)
	var mu []float64 = make([]float64, N+1)
	var z []float64 = make([]float64, N+1)

	a = val
	log.Println("Это a = val ", a)
	c[0] = 1
	mu[0] = 0
	z[0] = 0

	for i := 0; i < N; i++ {
		h[i] = time[i+1] - time[i]
	}
	log.Println("Вектор шагов по времени h = ", h)
	for i := 1; i < N; i++ {
		alpha[i] = 3*(a[i+1]-a[i])/h[i] - 3*(a[i]-a[i-1])/h[i-1]
	}
	log.Println("alpha = ", alpha)

	for i := 1; i < N; i++ {
		l[i] = 2*(time[i+1]-time[i-1]) - h[i-1]*mu[i-1]
		mu[i] = h[i] / l[i]
		z[i] = (alpha[i] - h[i-1]*z[i-1]) / l[i]
	}
	// log.Prin
	l[N] = 1
	z[N] = 0
	c[N] = 0

	for j := N - 1; j >= 0; j-- {
		c[j] = z[j] - mu[j]*c[j+1]
		b[j] = (a[j+1]-a[j])/h[j] - h[j]*(c[j+1]+2*c[j])/3
		d[j] = (c[j+1] - c[j]) / 3 / h[j]
	}
	log.Println("c = ", c)
	log.Println("d = ", d)
	log.Println("b = ", b)

	var splines []*models.CubicSplineData
	for i := 0; i < N; i++ {
		var spline models.CubicSplineData
		spline.A = a[i]
		spline.B = b[i]
		spline.C = c[i]
		spline.D = d[i]
		spline.X = time[i]
		// output_set[i].x = x[i]
		log.Println(spline)
		splines = append(splines, &spline)
	}
	return splines
}

// func (k *Kinematics) getCubicSpline(time int, val float64) *CubicSplineData {
// 	/* Кубический сплайн вида: S(t) = a + b*(t - t_prev) + c*(t - t_prev)^2 + d*(t - t_prev)^3 */
// 	var cubicSpline CubicSplineData

// 	return &cubicSpline
// }

// func (k *Kinematics) {

// }
