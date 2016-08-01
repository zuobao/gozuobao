package util
import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_IsPublicIP(t *testing.T) {
    assert.True(t, IsPublicIP("223.5.5.5"))
    assert.False(t, IsPublicIP("127.0.0.1"))
    assert.False(t, IsPublicIP("10.0.1.2"))
    assert.True(t, IsPublicIP("172.1.121.2"))
    assert.False(t, IsPublicIP("172.16.121.2"))
    assert.True(t, IsPublicIP("172.33.121.2"))
    assert.False(t, IsPublicIP("192.168.1.1"))
}
