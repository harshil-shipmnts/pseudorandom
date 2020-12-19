// https://en.wikipedia.org/wiki/Barnsley_fern

package main

type rule struct {
	a      float64
	b      float64
	c      float64
	d      float64
	tx     float64
	ty     float64
	weight float64
	color  string
}

var rules = []rule{
	rule{
		a:      0.85,
		b:      0.04,
		c:      -0.04,
		d:      0.85,
		tx:     0,
		ty:     1.6,
		weight: 0.85,
		color:  "#2ECC71",
	},
	rule{
		a:      -0.15,
		b:      0.28,
		c:      0.26,
		d:      0.24,
		tx:     0,
		ty:     0.44,
		weight: 0.07,
		color:  "#2E86C1",
	},
	rule{
		a:      0.2,
		b:      -0.26,
		c:      0.23,
		d:      0.22,
		tx:     0,
		ty:     1.6,
		weight: 0.07,
		color:  "#F7DC6F",
	},
	rule{
		a:      0,
		b:      0,
		c:      0,
		d:      0.16,
		tx:     0,
		ty:     0,
		weight: 0.01,
		color:  "#6E2C00",
	},
}

var x = nextRandFloat()
var y = nextRandFloat()

var iterateCount = 0

func iterate() {
	for i := 0; i < 100; i++ {
		var rule = getRule()
		x1 := x*rule.a + y*rule.b + rule.tx
		y1 := x*rule.c + y*rule.d + rule.ty
		x = x1
		y = y1
		plot(x, y, rule.color)
	}
	iterateCount++
	// if iterateCount%1000 == 0 {
	// 	cv.Save()
	// 	os.Exit(0)
	// }
}

func plot(x float64, y float64, color string) {
	cv.SetFillStyle(color)
	cv.FillRect(x*100, y*100, 1, 1)
}

func getRule() rule {
	var rand = nextRandFloat()
	for i := 0; i < len(rules); i++ {
		var rule = rules[i]
		if rand < rule.weight {
			return rule
		}
		rand = rand - rule.weight
	}
	return rules[0]
}
