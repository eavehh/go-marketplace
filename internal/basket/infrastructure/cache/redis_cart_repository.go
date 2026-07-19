package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/eavehh/marketpl.microserv/internal/basket/application/interfaces"
	"github.com/eavehh/marketpl.microserv/internal/basket/domain"
	"github.com/redis/go-redis/v9"
)

// TODO добавить логирования для ошибок, Особенно для delete
const cacheTTL = 30 * time.Second

type Redis_cart_repository struct {
	repo   interfaces.Cart_repository
	client *redis.Client
}

func New_redis_cart_repository(repo interfaces.Cart_repository, client *redis.Client) *Redis_cart_repository {
	return &Redis_cart_repository{repo: repo,
		client: client}
}

func (r *Redis_cart_repository) Save(ctx context.Context, cart *domain.Shopping_cart) (*domain.Shopping_cart, error) {
	result, err := r.repo.Save(ctx, cart)
	// сразу кладем в pq так как кэщ вторичен
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(result)
	if err == nil {
		_ = r.client.Set(ctx, result.Account_name, data, cacheTTL).Err()
	}

	return result, nil
}

func (r *Redis_cart_repository) Get(ctx context.Context, account_name string) (*domain.
	Shopping_cart, error) {
	cached, err := r.client.Get(ctx, account_name).Result()
	if err == nil && cached != "" {
		var cart domain.Shopping_cart
		err = json.Unmarshal([]byte(cached), &cart)
		if err == nil {
			return &cart, nil
		}
	}

	cart, err := r.repo.Get(ctx, account_name)
	if err != nil {
		return nil, err
	}

	if data, err := json.Marshal(cart); err == nil {
		_ = r.client.Set(ctx, cart.Account_name, data, cacheTTL).Err()
	}

	return cart, nil

}

func (r *Redis_cart_repository) Delete(ctx context.Context, account_name string) error {
	err := r.repo.Delete(ctx, account_name)
	if err != nil {
		return err
	}

	_ = r.client.Del(ctx, account_name).Err()

	return nil
}
