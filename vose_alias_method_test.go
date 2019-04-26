package algo

import (
	"fmt"
	"testing"
)

func TestVoseAlias(t *testing.T) {
	dist := []float64{3, 4, 5, 2, 1, 10}
	vas, err := NewVoseAliasMethod(dist, 100000000)
	if err != nil {

	} else {
		ret := vas.Sample()
		cmap := make(map[int]int)
		for _, r := range ret {
			cmap[r]++
		}
		for i := 0; i < len(dist); i++ {
			fmt.Println("sample prob:", float64(cmap[i])/float64(vas.NSample), "real prob:", vas.NormalDist[i])
		}
	}
}
