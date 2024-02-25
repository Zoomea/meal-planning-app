package fsdatabase

import (
	"testing"

	"github.com/Zoomea/meal-planning-app/testutil"
)

// Test the logic for setting the ID field if it is present and has type int64.
func TestSetID(t *testing.T) {
	hasID := struct {
		ID int64
	}{
		ID: 1,
	}

	setIDIfExists(&hasID, 2)
	testutil.Assert(t, hasID.ID, 2)

	// Mainly just testing nothing panics
	noID := struct{}{}
	setIDIfExists(&noID, 1)
}
