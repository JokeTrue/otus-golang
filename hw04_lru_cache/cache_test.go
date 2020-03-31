package hw04_lru_cache //nolint:golint,stylecheck

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("clear", func(t *testing.T) {
		c := NewCache(5)

		for i := 0; i < 5; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}

		c.Clear()
		_, ok := c.Get("1")
		require.False(t, ok)
	})

	t.Run("complex", func(t *testing.T) {
		c := NewCache(30)

		for i := 0; i < 31; i++ {
			c.Set(Key('A'+i), i)
		}

		val, ok := c.Get("A")
		require.False(t, ok)
		require.Nil(t, val)

		wasInCache := c.Set("B", 1)
		require.True(t, wasInCache)
	})

	t.Run("key movements", func(t *testing.T) {
		c := NewCache(3)

		c.Set("A", 1)
		c.Set("B", 2)
		c.Set("C", 3)

		c.Get("A")
		c.Get("B")
		c.Get("A")
		c.Get("B")
		c.Set("D", 4)

		val, ok := c.Get("C")
		require.False(t, ok)
		require.Nil(t, val)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
