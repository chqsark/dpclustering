package main

import (
	"github.com/chqsark/dpclustering"
	matrix "github.com/skelterjohn/go.matrix"
	"fmt"
	"math/rand"
	/*
	"code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/plotutil"
	*/
)	

func norm(min, max float64) float64 {
  return rand.Float64() * (max - min) + min
}

func main() {
	
	rows, cols := 30, 2
	var i, j int
	elem := make(map[int]float64)
    
	for i = 0; i < rows*cols; i++ {
		if i % 2 == 0 {
			elem[i] = norm(0, 5)	
		} else {
			elem[i] = norm(0, 10)	
		}	
	}

	m := matrix.MakeSparseMatrix(elem, rows, cols)
	
	var instances []dpclustering.Instance
	instances = dpclustering.Initialization(m)
	dpclustering.GetRho(instances, 0.1)
	dpclustering.GetDelta(instances) 
	dpclustering.GetClusters(instances, 3)
	fmt.Println("params:")
	for i = 0; i < rows; i++ {	
		fmt.Printf("%v", i)
		for j = 0; j < cols; j++ {
			fmt.Printf(" %v", instances[i].Data.Get(i,j))
		}
		fmt.Printf(" %v %v %v\n",instances[i].Rho, instances[i].Delta, instances[i].Cluster)
		
	}

}

