package driver

import (
	"github.com/MattLaidlaw/godb-go-driver/pkg/assert"
	"testing"
)

func TestClient(t *testing.T) {
	client, err := NewClient("localhost:6342")
	assert.ExpectEq(err, nil, t)

	insertedCount, err := client.Set("key", "val")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(insertedCount, 1, t)

	value, err := client.Get("key")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(value, "val", t)

	deletedCount, err := client.Del("key")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(deletedCount, 1, t)

	value, err = client.Get("key")
	assert.ExpectEq(err, nil, t)
	assert.ExpectEq(value, "", t)
}
