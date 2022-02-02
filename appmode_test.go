package appmode_test

import (
	"testing"

	"github.com/Kichiyaki/appmode/v2"

	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	assert.True(t, appmode.Equal(appmode.DevelopmentMode))

	appmode.Set(appmode.ProductionMode)
	assert.True(t, appmode.Equal(appmode.ProductionMode))

	appmode.Set(appmode.DevelopmentMode)
	assert.True(t, appmode.Equal(appmode.DevelopmentMode))

	appmode.Set(appmode.TestMode)
	assert.True(t, appmode.Equal(appmode.TestMode))

	appmode.Set("")
	assert.True(t, appmode.Equal(appmode.DevelopmentMode))

	assert.Panics(t, func() {
		appmode.Set("anyother")
	})
}
