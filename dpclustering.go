package dpclustering

import (
	matrix "github.com/skelterjohn/go.matrix"
	"container/heap"
	"math"
	"sort"
)

type Instance struct {
	Rho float64
	Delta float64
	Data *matrix.SparseMatrix
	Index int
	Cluster int
}

type PriorityQ []float64

func (pq PriorityQ) Len() int { return len(pq) }
func (pq PriorityQ) Less(i, j int) bool { return pq[i] > pq[j] }
func (pq PriorityQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQ) Push(x interface{}) {
	*pq = append(*pq, x.(float64))
}

func (pq *PriorityQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func Initialization(M *matrix.SparseMatrix) []Instance {
	rows, _ := M.GetSize()
	result := make([]Instance, rows)
	for i := 0; i < rows; i++ {
		result[i].Data = M.GetRowVector(i)
		result[i].Index = i
	}
	return result
}

func GetRho(instances []Instance, ratio float32) {
//rho is the mean distance to the nearest K neighbors
//details for this: http://rseghers.com/machine-learning/rodriguez-laoi-clustering/
	if ratio > 1 {
		panic("ratio should be (0, 1]")
	}
	rows := len(instances)
	K := int(float32(rows) * ratio)

	for i := 0; i < rows; i++ {
		nn := 0
		sum := 0.0
		tmp := make(PriorityQ, 0)
		heap.Init(&tmp)			
		for j := 0; j < rows; j++ {
			if i == j { continue }
			d := getDistance(instances[i].Data, instances[j].Data)
			if nn < K {
				heap.Push(&tmp, d)
				nn++
				sum += d
			} else {	
				if d < tmp[0] {
					sum = sum - tmp[0] + d
					tmp[0] = d
					heap.Fix(&tmp, 0)
				}
			}
		}
		instances[i].Rho = float64(K) / sum 
	}
}

func getDistance(v1 *matrix.SparseMatrix, v2 *matrix.SparseMatrix) float64 {
	_, cols := v1.GetSize()
	d := 0.0
	
	for i := 0; i < cols; i++ {
		d += math.Pow(v1.Get(0, i) - v2.Get(0, i), 2) 
	}
	return math.Sqrt(d)

}

func GetDelta(instances []Instance)  {
	rows := len(instances)
	var tmp float64
		
	for i := 0; i < rows; i++ {
		min, max := math.Inf(1), 0.0
		for j := 0; j < rows; j++ {
			if instances[i].Rho < instances[j].Rho {
				tmp = getDistance(instances[i].Data, instances[j].Data)
				if min > tmp {
					min = tmp
				}
			}
		}
		if min != math.Inf(1) {
			instances[i].Delta = min
		} else {
			for j := 0; j < rows; j++ {
				tmp = getDistance(instances[i].Data, instances[j].Data)
				max = math.Max(tmp, max)
			}
			instances[i].Delta = max
		}	
	}
}		

type ByRhoDelta []Instance

func (a ByRhoDelta) Len() int { return len(a) }
func (a ByRhoDelta) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRhoDelta) Less(i, j int) bool { return a[i].Rho * a[i].Delta > a[j].Rho * a[j].Delta }


func GetClusters(instances []Instance, N int) {
	rows := len(instances)
	var i, j, idx int
	
	sort.Sort(ByRhoDelta(instances))
	//"After the cluster centers have been found, each remaining point is assigned to the same 
	//cluster as its nearest neighbor of higher density"
	
	//propagating from the centers	

	for i = 0; i < rows; i++ {
		if i < N {
			instances[i].Cluster = i
		} else {
			min := math.Inf(1)
			for j = 0; j < i; j++ {
				if instances[j].Rho > instances[i].Rho {
					d := getDistance(instances[i].Data, instances[j].Data)
					if min > d {
						min = d
						idx = j
					}
				}
			}
			instances[i].Cluster = instances[idx].Cluster
		}
	}

}


	
