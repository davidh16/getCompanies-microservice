package server

import (
	"context"
	"getCompanies-microservice/initializers"
	"getCompanies-microservice/model"
	pb "getCompanies-microservice/proto"
	"getCompanies-microservice/utills"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math"
	"net"
)

func init() {
	initializers.Connect2DB()
}

type ListServer struct {
	pb.UnimplementedListOfCompaniesServer
}

func (l *ListServer) GetCompaniesList(ctx context.Context, req *pb.CompanyListReq) (*pb.CompaniesList, error) {
	list := []model.Company{}
	var total int64
	initializers.DB.Model(&model.Company{}).Count(&total)

	limit := req.GetLimit()
	if &limit == nil {
		limit = 10
	}

	page := req.GetPage()
	if &page == nil {
		page = 1
	}

	offset := (page - 1) * limit

	lastPage := math.Ceil(float64(total) / float64(limit))

	result := initializers.DB.Limit(int(limit)).Offset(int(offset)).Find(&list)
	if result.Error != nil {
		log.Printf("failed to fetch data: %v", result.Error)
		return nil, result.Error
	}
	response := pb.CompaniesList{}
	response.List = utills.CreateProtoList(list)
	response.Page = page
	response.Limit = limit
	response.LastPage = int64(lastPage)
	response.Total = total

	return &response, nil
}

func RunListServer() {
	lis, err := net.Listen("tcp", "localhost:5050")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	reflection.Register(s)

	pb.RegisterListOfCompaniesServer(s, &ListServer{})

	log.Println("Server running on port: 5050")

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}
