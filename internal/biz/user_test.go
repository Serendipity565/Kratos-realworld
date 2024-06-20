package biz

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	s := hashPassword("abc")
	spew.Dump(s)
	//panic(1)
}

func TestVerifyPassword(t *testing.T) {
	a := assert.New(t)

	a.True(verifyPassword("$2a$10$YSMn1Bx3I1.6l.SIw205auu4xZ8lKnUbGeOrjZAomfrOxsEO1QS3i", "abc"))
	a.False(verifyPassword("$2a$10$YSMn1Bx3I1.6l.SIw205auu4xZ8lKnUbGeOrjZAomfrOxsEO1QS3i", "abc1"))
	a.False(verifyPassword("$2a$10$YSMn1Bx3I1.6l.SIw205auu4xZ8lKnUbGeOr", "abc"))
}
