package auth

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	tk := GenerateToken("eric")
	spew.Dump(tk)
	panic("tk")
}
