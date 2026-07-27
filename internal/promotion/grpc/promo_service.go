package grpc

import (
	"context"

	"github.com/eavehh/marketpl.microserv/internal/promotion/application/commands"
	"github.com/eavehh/marketpl.microserv/internal/promotion/application/queries"
	"github.com/eavehh/marketpl.microserv/internal/promotion/grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PromotionService struct {
	pb.UnimplementedPromotionServiceServer
	GetByCatalogItem *queries.Get_by_catalog_item_handler
	CreateHandler    *commands.Create_promo_handler
}

func NewPromotionService(
	queryHandler *queries.Get_by_catalog_item_handler,
	commandHandler *commands.Create_promo_handler,
) *PromotionService {
	return &PromotionService{
		GetByCatalogItem: queryHandler,
		CreateHandler:    commandHandler,
	}
}

func (s PromotionService) GetPromoByCatalogItem(ctx context.Context, req *pb.GetPromoByCatalogItemRequest) (*pb.GetPromoByCatalogItemResponse, error) {
	query := &queries.Get_by_catalog_item_query{
		Catalog_item_id: req.CatalogItemId,
	}

	p, err := s.GetByCatalogItem.Handle(ctx, *query)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}
	if p == nil {
		return nil, status.Errorf(codes.NotFound, "found nothiong for: %s", req.CatalogItemId)
	}
	return &pb.GetPromoByCatalogItemResponse{
		Promotion: &pb.Promotion{
			Id:            p.Id,
			CatalogItemId: p.Catalog_item_id,
			Title:         p.Title,
			Value:         p.Value,
		},
	}, nil
}

func (s PromotionService) CreatePromo(ctx context.Context, req *pb.CreatePromoRequest) (*pb.CreatePromoResponse, error) {
	// if req == nil {
	// 	return nil, status.Errorf(codes.Canceled, "Internal error: %v")
	// }

	cmd := &commands.Create_promo_command{
		Catalog_item_id: req.CatalogItemId,
		Title:           req.Title,
		Value:           req.Value,
	}

	result, err := s.CreateHandler.Handle(ctx, *cmd)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal error: %v", err)
	}

	return &pb.CreatePromoResponse{
		Id:          result.Id,
		Success:     result.Success,
		Description: result.Description,
	}, nil
}
