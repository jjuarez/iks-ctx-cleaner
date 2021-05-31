package codes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodesString(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Codes[0].String(), "1001", "should be 1001")
	assert.Equal(Codes[1].String(), "1002", "should be 1002")
	assert.Equal(Codes[2].String(), "1003", "should be 1003")
}

func TestCodesMessage(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(Codes[0].Message(), "READERROR", "should be READERROR")
	assert.Equal(Codes[1].Message(), "UNMARSHALERROR", "should be UNMARSHALERROR")
	assert.Equal(Codes[2].Message(), "MARSHALERROR", "should be MARSHALERROR")
}
