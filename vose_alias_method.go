package algo

import (
	"container/list"
	"errors"
	"fmt"
	"math/rand"
)

// VoseAlias ...
type VoseAlias struct {
	// 初始化分布数值
	Dist       []float64
	NSample    int64
	NormalDist []float64
	AliasTable []int
	ProbTable  []float64
}

// NewVoseAliasMethod ...
func NewVoseAliasMethod(dist []float64, nsample int64) (voseAlias *VoseAlias, err error) {
	var sum float64
	for _, v := range dist {
		if v < 0 {
			return voseAlias, errors.New("contain negetive number")
		}
		sum += v
	}
	var smallList = list.New()
	var largeList = list.New()
	var normalDist = make([]float64, len(dist))
	var probTable = make([]float64, len(dist))
	var aliasTable = make([]int, len(dist))
	var probSum float64
	for i, v := range dist {
		aliasTable[i] = -1 // default, no any alias
		normalDist[i] = v / sum
		probSum += normalDist[i]
		normalProbValue := float64(len(dist)) * normalDist[i]
		probTable[i] = normalProbValue
		if normalProbValue >= 1.0 {
			largeList.PushBack(i)
		} else {
			smallList.PushBack(i)
		}
	}
	fmt.Println("probSum:", probSum)
	if probSum != 1.0 {
		return voseAlias, errors.New("dist normalize error")
	}

	for largeList.Len() > 0 && smallList.Len() > 0 {
		// fmt.Println("slen:", smallList.Len(), ",llen:", largeList.Len())
		small := smallList.Front()
		small_v := small.Value.(int)
		smallList.Remove(small)
		large := largeList.Front()
		large_v := large.Value.(int)
		largeList.Remove(large)
		fmt.Println("get small:", small_v, ",prob:", probTable[small_v], ",large:", large_v, ",prob:", probTable[large_v])
		aliasTable[small_v] = large_v
		remaining := probTable[large_v] - (1.0 - probTable[small_v])
		probTable[large_v] = remaining
		if remaining >= 1.0 {
			largeList.PushBack(large_v)
		} else {
			smallList.PushBack(large_v)
		}
	}
	fmt.Println("build VoseAlias finished, aliasTable:", aliasTable, ",probTable:", probTable)
	return &VoseAlias{Dist: dist, NSample: nsample, NormalDist: normalDist, AliasTable: aliasTable, ProbTable: probTable}, nil
}

// Sample ...
func (va *VoseAlias) Sample() []int {
	n := len(va.Dist)
	var ret []int
	var i int64
	for i = 0; i < va.NSample; i++ {
		// random in which index
		r := rand.Float64()
		// random prod
		r2 := rand.Float64()
		ridx := int(r * float64(n))
		if va.AliasTable[ridx] == -1 {
			ret = append(ret, ridx)
		} else {
			if va.ProbTable[ridx] > r2 {
				ret = append(ret, ridx)
			} else {
				ret = append(ret, va.AliasTable[ridx])
			}
		}
	}
	return ret
}
