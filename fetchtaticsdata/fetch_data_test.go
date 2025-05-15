package fetchtaticsdata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchData(t *testing.T) {
	err := fetchData()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	assert.Equal(t, err, nil, "err should be nil")
}
