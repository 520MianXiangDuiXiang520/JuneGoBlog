package dao

import (
	"testing"
)

func TestQueryArticleTotalByTagIDFromCache(t *testing.T) {

}

func TestQueryTagByIDFromCache(t *testing.T) {
	_, _ = queryTagByIDFromCache(13)
}
