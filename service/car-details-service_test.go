package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDetails(t *testing.T) {
	carDetailsService := NewCarDetailsService()
	resp := carDetailsService.GetDetails()
	assert.NotNil(t, resp)
	assert.Equal(t, 1, resp.ID)
	assert.Equal(t, "Mitsubishi", resp.Brand)
	assert.Equal(t, "Montero", resp.Model)
	assert.Equal(t, 2002, resp.Year)
	assert.Equal(t, "Alyosha", resp.FirstName)
	assert.Equal(t, "Caldero", resp.LastName)
	assert.Equal(t, "acaldero0@behance.net", resp.Email)
}