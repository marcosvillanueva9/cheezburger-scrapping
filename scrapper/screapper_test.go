package scrapper_test

import (
	"testing"

	"github.com/marcosvillanueva9/cheezburger-scrapping/scrapper"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	tests := []int{1,10,20,50}

	for testNum := range tests {
		Exec(t, testNum)
	}
}

func Exec(t *testing.T, linkNums int) {
	
	err, links := scrapper.Run(linkNums)

	assert.Nil(t, err)
	assert.Equal(t, linkNums, len(links))
}
