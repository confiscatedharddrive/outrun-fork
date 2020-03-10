package consts

// in the future, integrity checking will be put in place, and this will be used for it
var PointScores = map[int64][]int64{ // chapterNum: []int{...}
	1:  []int64{1000, 1000, 1000, 1000, 1000, -1}, // last element is the boss fight point, should never encounter
	2:  []int64{50000, 50000, 50000, 50000, 50000, -1},
	3:  []int64{50000, 50000, 50000, 50000, 50000, -1},
	4:  []int64{65000, 65000, 65000, 65000, 65000, -1},
	5:  []int64{65000, 65000, 65000, 65000, 65000, -1}, //This one get glitched sometimes with no apparent reason -F
	6:  []int64{65000, 65000, 65000, 65000, 65000, -1},
	7:  []int64{110000, 110000, 110000, 110000, 110000, -1},
	8:  []int64{110000, 110000, 110000, 110000, 110000, -1},
	9:  []int64{160000, 160000, 160000, 160000, 160000, -1},
	10: []int64{160000, 160000, 160000, 160000, 160000, -1},
	11: []int64{260000, 260000, 260000, 260000, 260000, -1},
	12: []int64{390000, 390000, 390000, 390000, 390000, -1},
	13: []int64{420000, 420000, 420000, 420000, 420000, -1},
	14: []int64{420000, 420000, 420000, 420000, 420000, -1},
	15: []int64{420000, 420000, 420000, 420000, 420000, -1},
	16: []int64{600000, 600000, 600000, 600000, 600000, -1},
	17: []int64{1200000, 1200000, 1200000, 1200000, 1200000, -1},
	18: []int64{1280000, 1280000, 1280000, 1280000, 1280000, -1},
	19: []int64{1600000, 1600000, 1600000, 1600000, 1600000, -1},
	20: []int64{1700000, 1700000, 1700000, 1700000, 1700000, -1},
	21: []int64{2040000, 2040000, 2040000, 2040000, 2040000, -1},
	22: []int64{2160000, 2160000, 2160000, 2160000, 2160000, -1},
	23: []int64{2280000, 2280000, 2280000, 2280000, 2280000, -1},
	24: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	// VALUES BEYOND THIS POINT HAVE NOT BEEN FINALIZED!
	25: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	26: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	27: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	28: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	29: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	30: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	31: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	32: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	33: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	34: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	35: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	36: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	37: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	38: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	39: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	40: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	41: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	42: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	43: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	44: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	45: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	46: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	47: []int64{2800000, 2800000, 2800000, 2800000, 2800000, -1},
	48: []int64{14560000, 14560000, 14560000, 14560000, 14560000, -1},
	49: []int64{16800000, 16800000, 16800000, 16800000, 16800000, -1},
	50: []int64{40000000, 40000000, 40000000, 40000000, 40000000, -1},

	// 9xx series - Test Levels
	901: []int64{1000, 1000, 1000, 1000, 1000, -1},
	902: []int64{1000, 1000, 1000, 1000, 1000, -1},
}
