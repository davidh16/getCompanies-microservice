package utills

import (
	"getCompanies-microservice/mapper"
	"getCompanies-microservice/model"
	pb "getCompanies-microservice/proto"
)

func CreateProtoList(list []model.Company) []*pb.Company {
	l := []*pb.Company{}

	for _, company := range list {
		companyProto := mapper.Model2Proto(company)
		l = append(l, companyProto)
	}
	return l
}
