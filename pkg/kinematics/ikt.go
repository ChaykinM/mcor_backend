package kinematics

import (
	"log"
	"math"

	"main.go/pkg/models"
	// "../models"
)

func (k *Kinematics) MsomInverseSolver(task *models.MsomInverseTaskRequest) *models.MsomInverseTaskSolution {
	var solution models.MsomInverseTaskSolution

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

	var q5 float64 = task.Betta
	var q6 float64 = task.Alpha

	var Q_5 float64 = -task.X*math.Cos(q6) + task.Y*math.Sin(q6)
	var Q_6 = q6

	var d_2 = l_1 + l_2 + l_3

	var Y_P0 float64 = task.Y*math.Cos(q6) + task.X*math.Sin(q6)
	var Z_P0 float64 = task.Z + d_2

	var Y_P2 float64 = -P4O*math.Sin(q5) + Y_P0
	var Z_P2 float64 = P4O*math.Cos(q5) + Z_P0

	var Y_C float64 = -(L_CD/2)*math.Cos(q5) + Y_P2
	var Z_C float64 = -(L_CD/2)*math.Sin(q5) + Z_P2
	var Y_D float64 = (L_CD/2)*math.Cos(q5) + Y_P2
	var Z_D float64 = (L_CD/2)*math.Sin(q5) + Z_P2

	var L_AC float64 = math.Sqrt(((Y_C + L) * (Y_C + L)) + ((Z_C - h) * (Z_C - h)))
	var L_FD float64 = math.Sqrt(((Y_D - L) * (Y_D - L)) + ((Z_D - h) * (Z_D - h)))

	var q1 float64 = math.Acos((L1*L1 + L_AC*L_AC - L2*L2) / (2 * L1 * L_AC))
	var q2 float64 = math.Atan((Z_C - h) / (Y_C + L))
	var q3 float64 = math.Acos((L3*L3 + L_FD*L_FD - L4*L4) / (2 * L3 * L_FD))
	var q4 float64 = math.Atan((Z_D-h)/(Y_D-L)) + math.Pi

	var Q_1 float64 = (q2 + q1) //qreal
	var Q_2 float64 = -(math.Acos((L_AC*L_AC - L1*L1 - L2*L2) / (2 * L1 * L2)))
	var Q_3 float64 = (q4 - q3) //qreal
	var Q_4 float64 = (math.Acos((L_FD*L_FD - L4*L4 - L3*L3) / (2 * L4 * L3)))

	solution.Time = task.Time
	solution.Q_1 = Q_1
	solution.Q_2 = Q_2
	solution.Q_3 = Q_3
	solution.Q_4 = Q_4
	solution.Q_5 = Q_5
	solution.Q_6 = Q_6

	log.Println(solution)

	return &solution
}

func (k *Kinematics) FiveBarIverseSolver() {

}
