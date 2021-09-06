package driver

import (
	"github.com/MattLaidlaw/go-assert"
	"testing"
)

func TestClient(t *testing.T) {
	client, err := NewClient(":6342")
	assert.ExpectEq(err, nil, t)

	value, err := client.Get("key")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(value, "", t)

	insertedCount, err := client.Set("key", "val")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(insertedCount, float64(1), t)

	value, err = client.Get("key")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(value, "val", t)

	deletedCount, err := client.Del("key")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(deletedCount, float64(1), t)

	value, err = client.Get("key")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(value, "", t)

	deletedCount, err = client.Del("key")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(deletedCount, float64(0), t)
}
