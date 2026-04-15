package cache

import (
	"github.com/Koderbek/link_storage_service/internal/database"
	"github.com/Koderbek/link_storage_service/internal/helper"
	"github.com/Koderbek/link_storage_service/internal/model"
	"sync"
)

const batchSize = 10

type LinkCacheItem struct {
	model.Link
	savedVisits uint
}

type LinkCache struct {
	mu    sync.Mutex
	repo  *database.Repository
	links map[string]*LinkCacheItem
}

func NewLinkCache(repo *database.Repository) *LinkCache {
	return &LinkCache{
		repo:  repo,
		links: make(map[string]*LinkCacheItem),
	}
}

func (c *LinkCache) GetAndIncr(key string) (*LinkCacheItem, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	val, ok := c.links[key]
	id := helper.CodeToId(key)
	if !ok {
		link, err := c.repo.Link(id)
		if err != nil {
			return nil, err
		}

		val = &LinkCacheItem{Link: *link, savedVisits: link.Visits}
		c.links[key] = val
	}

	val.Visits++
	if val.Visits-val.savedVisits >= batchSize {
		//Делаем сохранение каждые 10 посещений
		go func(id, visits uint) {
			c.repo.UpdateVisits(id, visits)
		}(id, val.Visits)
		val.savedVisits = val.Visits
	}

	return val, nil
}

func (c *LinkCache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.links, key)
}
