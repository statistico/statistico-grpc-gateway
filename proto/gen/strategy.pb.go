// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.6.1
// source: strategy.proto

package statistico

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type StrategyTradeSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Market             string               `protobuf:"bytes,1,opt,name=market,proto3" json:"market,omitempty"`
	MarketRunnerFilter *RunnerFilter        `protobuf:"bytes,2,opt,name=market_runner_filter,json=marketRunnerFilter,proto3" json:"market_runner_filter,omitempty"`
	CompetitionIds     []uint64             `protobuf:"varint,3,rep,packed,name=competition_ids,json=competitionIds,proto3" json:"competition_ids,omitempty"`
	SeasonIds          []uint64             `protobuf:"varint,4,rep,packed,name=season_ids,json=seasonIds,proto3" json:"season_ids,omitempty"`
	DateFrom           *timestamp.Timestamp `protobuf:"bytes,5,opt,name=dateFrom,proto3" json:"dateFrom,omitempty"`
	DateTo             *timestamp.Timestamp `protobuf:"bytes,6,opt,name=dateTo,proto3" json:"dateTo,omitempty"`
	ResultFilters      []*ResultFilter      `protobuf:"bytes,7,rep,name=result_filters,json=resultFilters,proto3" json:"result_filters,omitempty"`
	StatFilters        []*StatFilter        `protobuf:"bytes,8,rep,name=stat_filters,json=statFilters,proto3" json:"stat_filters,omitempty"`
}

func (x *StrategyTradeSearchRequest) Reset() {
	*x = StrategyTradeSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyTradeSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyTradeSearchRequest) ProtoMessage() {}

func (x *StrategyTradeSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyTradeSearchRequest.ProtoReflect.Descriptor instead.
func (*StrategyTradeSearchRequest) Descriptor() ([]byte, []int) {
	return file_strategy_proto_rawDescGZIP(), []int{0}
}

func (x *StrategyTradeSearchRequest) GetMarket() string {
	if x != nil {
		return x.Market
	}
	return ""
}

func (x *StrategyTradeSearchRequest) GetMarketRunnerFilter() *RunnerFilter {
	if x != nil {
		return x.MarketRunnerFilter
	}
	return nil
}

func (x *StrategyTradeSearchRequest) GetCompetitionIds() []uint64 {
	if x != nil {
		return x.CompetitionIds
	}
	return nil
}

func (x *StrategyTradeSearchRequest) GetSeasonIds() []uint64 {
	if x != nil {
		return x.SeasonIds
	}
	return nil
}

func (x *StrategyTradeSearchRequest) GetDateFrom() *timestamp.Timestamp {
	if x != nil {
		return x.DateFrom
	}
	return nil
}

func (x *StrategyTradeSearchRequest) GetDateTo() *timestamp.Timestamp {
	if x != nil {
		return x.DateTo
	}
	return nil
}

func (x *StrategyTradeSearchRequest) GetResultFilters() []*ResultFilter {
	if x != nil {
		return x.ResultFilters
	}
	return nil
}

func (x *StrategyTradeSearchRequest) GetStatFilters() []*StatFilter {
	if x != nil {
		return x.StatFilters
	}
	return nil
}

type StrategyTrade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MarketName    string               `protobuf:"bytes,1,opt,name=market_name,json=marketName,proto3" json:"market_name,omitempty"`
	RunnerName    string               `protobuf:"bytes,2,opt,name=runner_name,json=runnerName,proto3" json:"runner_name,omitempty"`
	RunnerPrice   float32              `protobuf:"fixed32,3,opt,name=runner_price,json=runnerPrice,proto3" json:"runner_price,omitempty"`
	EventId       uint64               `protobuf:"varint,4,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	CompetitionId uint64               `protobuf:"varint,5,opt,name=competition_id,json=competitionId,proto3" json:"competition_id,omitempty"`
	SeasonId      uint64               `protobuf:"varint,6,opt,name=season_id,json=seasonId,proto3" json:"season_id,omitempty"`
	EventDate     *timestamp.Timestamp `protobuf:"bytes,7,opt,name=event_date,json=eventDate,proto3" json:"event_date,omitempty"`
	Result        TradeResultEnum      `protobuf:"varint,8,opt,name=result,proto3,enum=statistico.TradeResultEnum" json:"result,omitempty"`
}

func (x *StrategyTrade) Reset() {
	*x = StrategyTrade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StrategyTrade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StrategyTrade) ProtoMessage() {}

func (x *StrategyTrade) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StrategyTrade.ProtoReflect.Descriptor instead.
func (*StrategyTrade) Descriptor() ([]byte, []int) {
	return file_strategy_proto_rawDescGZIP(), []int{1}
}

func (x *StrategyTrade) GetMarketName() string {
	if x != nil {
		return x.MarketName
	}
	return ""
}

func (x *StrategyTrade) GetRunnerName() string {
	if x != nil {
		return x.RunnerName
	}
	return ""
}

func (x *StrategyTrade) GetRunnerPrice() float32 {
	if x != nil {
		return x.RunnerPrice
	}
	return 0
}

func (x *StrategyTrade) GetEventId() uint64 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *StrategyTrade) GetCompetitionId() uint64 {
	if x != nil {
		return x.CompetitionId
	}
	return 0
}

func (x *StrategyTrade) GetSeasonId() uint64 {
	if x != nil {
		return x.SeasonId
	}
	return 0
}

func (x *StrategyTrade) GetEventDate() *timestamp.Timestamp {
	if x != nil {
		return x.EventDate
	}
	return nil
}

func (x *StrategyTrade) GetResult() TradeResultEnum {
	if x != nil {
		return x.Result
	}
	return TradeResultEnum_SUCCESS
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_strategy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_strategy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_strategy_proto_rawDescGZIP(), []int{2}
}

func (x *ErrorResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_strategy_proto protoreflect.FileDescriptor

var file_strategy_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x1a, 0x0a, 0x65, 0x6e,
	0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65,
	0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x03, 0x0a, 0x1a, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x12, 0x4a, 0x0a, 0x14,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x12, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x52, 0x75, 0x6e, 0x6e,
	0x65, 0x72, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70,
	0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x04, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x73,
	0x12, 0x36, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x32, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x12, 0x3f, 0x0a, 0x0e,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x6f, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0d,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x12, 0x39, 0x0a,
	0x0c, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22, 0xc3, 0x02, 0x0a, 0x0d, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x54, 0x72, 0x61, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72,
	0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f,
	0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xa2,
	0x01, 0x0a, 0x0d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x40, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x2c,
	0x92, 0x41, 0x29, 0x2a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x20, 0x63, 0x6f, 0x64, 0x65, 0x8a, 0x01, 0x07, 0x5e, 0x5b, 0x30, 0x2d, 0x39,
	0x5d, 0x24, 0xa2, 0x02, 0x07, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x4f, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x35, 0x92, 0x41, 0x32, 0x2a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x20, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x8a, 0x01, 0x14, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d,
	0x39, 0x5d, 0x7b, 0x31, 0x2c, 0x20, 0x33, 0x32, 0x7d, 0x24, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xd5, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc1, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x26, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x2e, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x54, 0x72, 0x61, 0x64, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73,
	0x74, 0x69, 0x63, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x54, 0x72, 0x61,
	0x64, 0x65, 0x22, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x19, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2f, 0x74, 0x72, 0x61, 0x64, 0x65, 0x2d, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0x3e, 0x12, 0x15, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x20, 0x54, 0x72, 0x61, 0x64, 0x65, 0x20, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x1a, 0x25, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x61, 0x20, 0x74, 0x72,
	0x61, 0x64, 0x65, 0x20, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61,
	0x20, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x30, 0x01, 0x42, 0xab, 0x03, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f,
	0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x3b, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x92, 0x41, 0xed, 0x02, 0x12, 0x1e, 0x0a, 0x17,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x20, 0x53, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x20, 0x41, 0x50, 0x49, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x02, 0x32,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x52, 0x42, 0x0a, 0x03, 0x34, 0x32, 0x32, 0x12, 0x3b, 0x0a, 0x1a, 0x49, 0x6e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x65,
	0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x1b, 0x1a, 0x19, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x3d, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x36,
	0x0a, 0x15, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x1b, 0x1a, 0x19, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x39, 0x0a, 0x03, 0x35, 0x30, 0x32, 0x12, 0x32, 0x0a,
	0x11, 0x42, 0x61, 0x64, 0x20, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x20, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x1b, 0x1a, 0x19, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74,
	0x69, 0x63, 0x6f, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x72, 0x66, 0x0a, 0x2d, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f,
	0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x35, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x6f, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x6f, 0x2d, 0x67, 0x72, 0x70,
	0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_strategy_proto_rawDescOnce sync.Once
	file_strategy_proto_rawDescData = file_strategy_proto_rawDesc
)

func file_strategy_proto_rawDescGZIP() []byte {
	file_strategy_proto_rawDescOnce.Do(func() {
		file_strategy_proto_rawDescData = protoimpl.X.CompressGZIP(file_strategy_proto_rawDescData)
	})
	return file_strategy_proto_rawDescData
}

var file_strategy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_strategy_proto_goTypes = []interface{}{
	(*StrategyTradeSearchRequest)(nil), // 0: statistico.StrategyTradeSearchRequest
	(*StrategyTrade)(nil),              // 1: statistico.StrategyTrade
	(*ErrorResponse)(nil),              // 2: statistico.ErrorResponse
	(*RunnerFilter)(nil),               // 3: statistico.RunnerFilter
	(*timestamp.Timestamp)(nil),        // 4: google.protobuf.Timestamp
	(*ResultFilter)(nil),               // 5: statistico.ResultFilter
	(*StatFilter)(nil),                 // 6: statistico.StatFilter
	(TradeResultEnum)(0),               // 7: statistico.TradeResultEnum
}
var file_strategy_proto_depIdxs = []int32{
	3, // 0: statistico.StrategyTradeSearchRequest.market_runner_filter:type_name -> statistico.RunnerFilter
	4, // 1: statistico.StrategyTradeSearchRequest.dateFrom:type_name -> google.protobuf.Timestamp
	4, // 2: statistico.StrategyTradeSearchRequest.dateTo:type_name -> google.protobuf.Timestamp
	5, // 3: statistico.StrategyTradeSearchRequest.result_filters:type_name -> statistico.ResultFilter
	6, // 4: statistico.StrategyTradeSearchRequest.stat_filters:type_name -> statistico.StatFilter
	4, // 5: statistico.StrategyTrade.event_date:type_name -> google.protobuf.Timestamp
	7, // 6: statistico.StrategyTrade.result:type_name -> statistico.TradeResultEnum
	0, // 7: statistico.StrategyService.StrategyTradeSearch:input_type -> statistico.StrategyTradeSearchRequest
	1, // 8: statistico.StrategyService.StrategyTradeSearch:output_type -> statistico.StrategyTrade
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_strategy_proto_init() }
func file_strategy_proto_init() {
	if File_strategy_proto != nil {
		return
	}
	file_enum_proto_init()
	file_filter_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_strategy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyTradeSearchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_strategy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StrategyTrade); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_strategy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_strategy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strategy_proto_goTypes,
		DependencyIndexes: file_strategy_proto_depIdxs,
		MessageInfos:      file_strategy_proto_msgTypes,
	}.Build()
	File_strategy_proto = out.File
	file_strategy_proto_rawDesc = nil
	file_strategy_proto_goTypes = nil
	file_strategy_proto_depIdxs = nil
}