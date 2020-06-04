// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: weather.ext.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for WeatherService service

func NewWeatherServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for WeatherService service

type WeatherService interface {
	//获取实况的天气数据
	Now(ctx context.Context, in *DataReq, opts ...client.CallOption) (*NowData, error)
	//获取天气预报数据
	Forecast(ctx context.Context, in *DataReq, opts ...client.CallOption) (*ForecastData, error)
}

type weatherService struct {
	c    client.Client
	name string
}

func NewWeatherService(name string, c client.Client) WeatherService {
	return &weatherService{
		c:    c,
		name: name,
	}
}

func (c *weatherService) Now(ctx context.Context, in *DataReq, opts ...client.CallOption) (*NowData, error) {
	req := c.c.NewRequest(c.name, "WeatherService.Now", in)
	out := new(NowData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *weatherService) Forecast(ctx context.Context, in *DataReq, opts ...client.CallOption) (*ForecastData, error) {
	req := c.c.NewRequest(c.name, "WeatherService.Forecast", in)
	out := new(ForecastData)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WeatherService service

type WeatherServiceHandler interface {
	//获取实况的天气数据
	Now(context.Context, *DataReq, *NowData) error
	//获取天气预报数据
	Forecast(context.Context, *DataReq, *ForecastData) error
}

func RegisterWeatherServiceHandler(s server.Server, hdlr WeatherServiceHandler, opts ...server.HandlerOption) error {
	type weatherService interface {
		Now(ctx context.Context, in *DataReq, out *NowData) error
		Forecast(ctx context.Context, in *DataReq, out *ForecastData) error
	}
	type WeatherService struct {
		weatherService
	}
	h := &weatherServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&WeatherService{h}, opts...))
}

type weatherServiceHandler struct {
	WeatherServiceHandler
}

func (h *weatherServiceHandler) Now(ctx context.Context, in *DataReq, out *NowData) error {
	return h.WeatherServiceHandler.Now(ctx, in, out)
}

func (h *weatherServiceHandler) Forecast(ctx context.Context, in *DataReq, out *ForecastData) error {
	return h.WeatherServiceHandler.Forecast(ctx, in, out)
}