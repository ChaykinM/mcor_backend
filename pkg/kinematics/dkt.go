package kinematics

import (
	"log"
	"math"

	// "../models"
	"gonum.org/v1/gonum/mat"
	"main.go/pkg/models"
)

// https://pkg.go.dev/gonum.org/v1/gonum/mat

// "main.go/pkg/models"

// https://github.com/gonum/gonum

func (k *Kinematics) MsomDirectSolver(task *models.MsomDirectTaskRequest) *models.MsomDirectTaskSolution {
	var solution models.MsomDirectTaskSolution

	var l_1, l_2, l_3, h, L, L1, L2, L_CD, L4, L3, P4O float64

	l_1, _ = k.findConfigParam("l_1")
	l_2, _ = k.findConfigParam("l_2")
	l_3, _ = k.findConfigParam("l_3")
	h, _ = k.findConfigParam("h")
	L, _ = k.findConfigParam("L")
	L1, _ = k.findConfigParam("L1")
	L2, _ = k.findConfigParam("L2")
	L_CD, _ = k.findConfigParam("L_CD")
	L4, _ = k.findConfigParam("L4")
	L3, _ = k.findConfigParam("L3")
	P4O, _ = k.findConfigParam("P4O")

	T0_data := []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, l_1,
		0, 0, 0, 1}
	T0 := mat.NewDense(4, 4, T0_data)

	T1_data := []float64{
		1, 0, 0, task.Q_5,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
	T1 := mat.NewDense(4, 4, T1_data)

	T2_data := []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, l_2,
		0, 0, 0, 1}
	T2 := mat.NewDense(4, 4, T2_data)

	T3_data := []float64{
		math.Cos(task.Q_6), -math.Sin(task.Q_6), 0, 0,
		math.Sin(task.Q_6), math.Cos(task.Q_6), 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
	T3 := mat.NewDense(4, 4, T3_data)

	T4_data := []float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, l_3,
		0, 0, 0, 1}
	T4 := mat.NewDense(4, 4, T4_data)

	var T_rot_mech mat.Dense
	var T0_T1 mat.Dense
	T0_T1.Mul(T0, T1)
	var T0_T1_T2 mat.Dense
	T0_T1_T2.Mul(&T0_T1, T2)
	var T0_T1_T2_T3 mat.Dense
	T0_T1_T2_T3.Mul(&T0_T1_T2, T3)
	T_rot_mech.Mul(&T0_T1_T2_T3, T4)

	var x_Pc float64 = 0
	var y_C float64 = -L + L1*math.Cos(task.Q_1) + L2*math.Cos(task.Q_1+task.Q_2)
	var y_D float64 = L + L3*math.Cos(task.Q_3) + L4*math.Cos(task.Q_3+task.Q_4)

	var cosd_betta = (y_D - y_C) / L_CD

	var z_C float64 = h + L1*math.Sin(task.Q_1) + L2*math.Sin(task.Q_1+task.Q_2)
	var z_D float64 = h + L3*math.Sin(task.Q_3) + L4*math.Sin(task.Q_3+task.Q_4)
	var sind_betta float64 = (z_D - z_C) / L_CD

	var y_Pc float64 = (y_C+y_D)/2 + P4O*sind_betta
	var z_Pc float64 = (z_C+z_D)/2 - P4O*cosd_betta

	T_six_mech_data := []float64{
		1, 0, 0, x_Pc,
		0, cosd_betta, -sind_betta, y_Pc,
		0, sind_betta, cosd_betta, z_Pc,
		0, 0, 0, 1}

	T_six_mech := mat.NewDense(4, 4, T_six_mech_data)

	var T_rot_mech_inv mat.Dense
	T_rot_mech_inv.Inverse(&T_rot_mech)

	var T_mech mat.Dense
	T_mech.Mul(&T_rot_mech_inv, T_six_mech)

	solution.Time = task.Time
	solution.X = T_mech.At(0, 3)
	solution.Y = T_mech.At(1, 3)
	solution.Z = T_mech.At(2, 3)
	solution.Alpha = task.Q_6
	solution.Betta = math.Asin(sind_betta)

	log.Println(solution)

	return &solution
}

func (k *Kinematics) MsomDirectValidation() {
	// проверка на границы допустимости в рабочем пространстве
	// проверка на сингулярность
	//

}

func (k *Kinematics) FiveBarDirectSolver() {

}

func (k *Kinematics) FiveBarDirectValidation() {

}
