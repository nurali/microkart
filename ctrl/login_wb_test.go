package ctrl

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractPasswd(t *testing.T) {
	t.Run("basic_ok", func(t *testing.T) {
		passwd, err := extractPasswd("Basic abcd1234")
		assert.NoError(t, err)
		assert.Equal(t, "abcd1234", passwd)
	})

	t.Run("blank", func(t *testing.T) {
		passwd, err := extractPasswd("")
		log.Println(err)
		log.Println(passwd)
		assert.Error(t, err)
		assert.Empty(t, passwd)
	})
}
