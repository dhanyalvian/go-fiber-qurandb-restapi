package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAyat_CollectionName(t *testing.T) {
	ayat := Ayat{}
	assert.Equal(t, COLLECTION_AYAT, ayat.CollectionName())
}
