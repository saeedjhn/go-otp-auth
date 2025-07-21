package gencode_test

import (
	"testing"

	"github.com/saeedjhn/go-otp-auth/pkg/generator"
)

//go:generate go test -v -race -count=1 -run TestGenUUID

func TestGenUUID_GeneratesValidUUID_Success(t *testing.T) {
	uuid := generator.GenUUID()
	if len(uuid) != 36 {
		t.Error("Expected UUID to be 36 characters long, got", len(uuid))
	}
}

func TestGenUUID_GeneratesUniqueUUIDs_Success(t *testing.T) {
	uuid1 := generator.GenUUID()
	uuid2 := generator.GenUUID()
	if uuid1 == uuid2 {
		t.Error("Expected UUIDs to be unique, got duplicates")
	}
}
