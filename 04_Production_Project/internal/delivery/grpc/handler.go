package grpc

import (
	"context"

	"github.com/Atharva0506/trading_bot/internal/service"
	signalpb "github.com/Atharva0506/trading_bot/proto/signal/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// SignalGRPCHandler implements the generated SignalServiceServer interface.
type SignalGRPCHandler struct {
	signalpb.UnimplementedSignalServiceServer
	service *service.SignalService
}

// NewSignalGRPCHandler returns a new gRPC handler backed by the SignalService.
func NewSignalGRPCHandler(service *service.SignalService) *SignalGRPCHandler {
	return &SignalGRPCHandler{
		service: service,
	}
}

// CreateSignal handles the gRPC CreateSignal RPC.
func (h *SignalGRPCHandler) CreateSignal(ctx context.Context, req *signalpb.CreateSignalRequest) (*signalpb.CreateSignalResponse, error) {
	signal, err := h.service.CreateSignal(ctx, req.GetSymbol(), req.GetAction(), req.GetPrice())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create signal: %v", err)
	}

	return &signalpb.CreateSignalResponse{
		Signal: &signalpb.Signal{
			Id:        signal.ID.String(),
			Symbol:    string(signal.Symbol),
			Action:    string(signal.Action),
			Price:     signal.Price,
			Timestamp: timestamppb.New(signal.Timestamp),
		},
	}, nil
}

// GetSignals handles the gRPC GetSignals RPC.
func (h *SignalGRPCHandler) GetSignals(ctx context.Context, req *signalpb.GetSignalsRequest) (*signalpb.GetSignalsResponse, error) {
	signals, err := h.service.GetAllSignals(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch signals: %v", err)
	}

	pbSignals := make([]*signalpb.Signal, len(signals))
	for i, s := range signals {
		pbSignals[i] = &signalpb.Signal{
			Id:        s.ID.String(),
			Symbol:    string(s.Symbol),
			Action:    string(s.Action),
			Price:     s.Price,
			Timestamp: timestamppb.New(s.Timestamp),
		}
	}
	return &signalpb.GetSignalsResponse{Signals: pbSignals}, nil
}

// GetSignalsBySymbol handles the gRPC GetSignalsBySymbol RPC.
func (h *SignalGRPCHandler) GetSignalsBySymbol(ctx context.Context, req *signalpb.GetSignalsBySymbolRequest) (*signalpb.GetSignalsResponse, error) {
	signals, err := h.service.GetSignalsBySymbol(ctx, req.GetSymbol())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch signals: %v", err)
	}

	pbSignals := make([]*signalpb.Signal, len(signals))
	for i, s := range signals {
		pbSignals[i] = &signalpb.Signal{
			Id:        s.ID.String(),
			Symbol:    string(s.Symbol),
			Action:    string(s.Action),
			Price:     s.Price,
			Timestamp: timestamppb.New(s.Timestamp),
		}
	}
	return &signalpb.GetSignalsResponse{Signals: pbSignals}, nil
}
