package main

import (
	"fmt"
)

const N = 4

func fillVector(vector [N*N]int) [N*N]int{
	for i:=0;i<N*N;i++{
		vector[i] = i+1
	}
	return vector
}

func vectorToMatrix(vector [N*N]int) [][]int {
	matrix := make([][]int,N)
	for i := 0;i<N;i++{
		matrix[i] = vector[i*N:i*N+N]
	}
	return matrix
}

func mainMatrixFromVector(vector [N*N]int,matrix[][]int) [][]int{
	for i:=0;i<N;i++{
		for j:=0;j<N;j++{
			matrix[i][j] = vector[j+i*N]
		}
	}
	return matrix
}

func bMatrix(vector [N*N]int,matrix [][]int)[][]int{
	for i:=0;i<N;i++{
		for j:=0;j<N;j++{
			matrix[i][j] = vector[N*j+i]
		}
	}
	return matrix 
}

func cMatrix(vector [N*N]int,matrix [][]int) [][]int{
	for i:=0;i<N;i++{
		for j:=0;j<N;j++{
			matrix[i][j] = vector[N*(j+1)-1-i]
		}
	}
	return matrix 
}


func dMatrix(vector [N*N]int,matrix [][]int)[][]int{
	for i:=0;i<N;i++{
		for j:=0;j<N;j++{
			matrix[i][j] = vector[N*(i+1)-1-j]
		}
	}
	return matrix 
}

func eMatrix(vector [N*N]int,matrix [][]int)[][]int{
	for i:=0;i<N;i++{
		for j:=0;j<N;j++{
			matrix[i][j] = vector[(N*N-(j+1))-i*N]
		}
	}
	return matrix 
}

func fMatrix(vector [N*N]int,matrix [][]int)[][]int{
	for i:=0;i<N;i++{
		for j:=0;j<N;j++{
			matrix[i][j] = vector[N*(N-j)-1-i]
		}
	}
	return matrix 
}

func gMatrix(vector [N*N]int,matrix [][]int)[][]int{
	for i:=0;i<N;i++{
		for j:=0;j<N;j++{
			matrix[i][j] = vector[N*(N-j-1)+i]
		}
	}
	return matrix 
}

func hMatrix(vector [N*N]int,matrix [][]int)[][]int{
	for i:=0;i<N;i++{
		for j:=0;j<N;j++{
			matrix[i][j] = vector[N*(N-(i+1))+j]
		}
	}
	return matrix 
}

func main(){
	var vector [N*N]int
	vector = fillVector(vector)
	matrix :=vectorToMatrix(vector)
	matrix = mainMatrixFromVector(vector, matrix)
	matrix = bMatrix(vector, matrix)
	matrix = cMatrix(vector, matrix)
	matrix = dMatrix(vector, matrix)
	matrix = eMatrix(vector, matrix)
	matrix = fMatrix(vector, matrix)
	matrix = gMatrix(vector, matrix)
	matrix = hMatrix(vector, matrix)
	fmt.Print(matrix)
	
}