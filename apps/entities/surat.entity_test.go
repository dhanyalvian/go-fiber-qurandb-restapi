package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSurat_CollectionName(t *testing.T) {
	surat := Surat{}
	assert.Equal(t, COLLECTION_SURAT, surat.CollectionName())
}
