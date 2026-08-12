package commands

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/eavehh/marketpl.microserv/internal/basket/application/interfaces"
	"github.com/eavehh/marketpl.microserv/internal/basket/domain"
	"github.com/eavehh/marketpl.microserv/internal/promotion/grpc/pb"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Save_cart_command struct {
	Cart domain.Shopping_cart `json:"cart" validate:"required" `
}

type Save_cart_handler struct {
	repo         interfaces.Cart_repository
	promo_client pb.PromotionServiceClient // интерфейс который требует определения всех rpc методов
}

func New_save_cart_handler(repo interfaces.Cart_repository, promo_client pb.PromotionServiceClient) *Save_cart_handler {
	return &Save_cart_handler{repo: repo, promo_client: promo_client}
}

func (h *Save_cart_handler) Handle(ctx context.Context, cmd Save_cart_command) (string, error) {
	if err := validate.Struct(&cmd); err != nil {
		return "", err
	}
	if (h.promo_client !=nil){
		_ = h.apply_promotions(ctx, &cmd.Cart)
	}
	_, err := h.repo.Save(ctx, &cmd.Cart)
	if err != nil {
		return "", err
	}
	return cmd.Cart.Account_name, nil
}

func (h *Save_cart_handler) get_discount_for_item(ctx context.Context, item *domain.Shopping_cart_item,
) (*pb.GetPromoByCatalogItemResponse, error) {
	req := &pb.GetPromoByCatalogItemRequest{
		CatalogItemId: item.Item_id.String(),
	}

	return h.promo_client.GetPromoByCatalogItem(ctx, req)
}

func (h *Save_cart_handler) apply_promotions(ctx context.Context, cart *domain.Shopping_cart) error {
	for i := range cart.Items {
		discount, err := h.get_discount_for_item(ctx, &cart.Items[i])
		if err != nil {
			return err
		}

		val, _ := parse_discount_value(discount)
		cart.Items[i].Unit_price -= val
	}
	return nil
}

func parse_discount_value(d *pb.GetPromoByCatalogItemResponse) (float64, error) {
	value := strings.TrimSpace(d.Promotion.Value)

	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("bad promo value: %q: %w", value, err)
	}
	return f, nil
}
