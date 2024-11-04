package main

import (
	"fmt"

	"github.com/lei006/gomath"
)

func main() {
	// 创建两个 2x2 的矩阵，使用 float64 类型
	A := gomath.Mat{1.0, 2.0, 3.0, 4.0}
	B := gomath.Mat{5.0, 6.0, 7.0, 8.0}

	// 矩阵加法
	C := gomath.Add(A, B)
	fmt.Println("A + B =", C)

	// 矩阵乘法
	D := gomath.Mul(A, B)
	fmt.Println("A * B =", D)

	// 矩阵转置
	AT := A.T()
	fmt.Println("A^T =", AT)

	// 计算逆矩阵
	AInv, err := A.Inv()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("A^-1 =", AInv)
	}

	// 计算行列式
	det := A.Det()
	fmt.Println("det(A) =", det)
}
