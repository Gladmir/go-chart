package chart

import (
	"testing"

	assert "github.com/blendlabs/go-assert"
)

func TestConcatSeries(t *testing.T) {
	assert := assert.New(t)

	s1 := ContinuousSeries{
		XValues: Seq(1.0, 10.0),
		YValues: Seq(1.0, 10.0),
	}

	s2 := ContinuousSeries{
		XValues: Seq(11, 20.0),
		YValues: Seq(10.0, 1.0),
	}

	s3 := ContinuousSeries{
		XValues: Seq(21, 30.0),
		YValues: Seq(1.0, 10.0),
	}

	cs := ConcatSeries([]Series{s1, s2, s3})
	assert.Equal(30, cs.Len())

	x0, y0 := cs.GetValue(0)
	assert.Equal(1.0, x0)
	assert.Equal(1.0, y0)

	xm, ym := cs.GetValue(19)
	assert.Equal(20.0, xm)
	assert.Equal(1.0, ym)

	xn, yn := cs.GetValue(29)
	assert.Equal(30.0, xn)
	assert.Equal(10.0, yn)
}