package mapper

import (
	"getCompanies-microservice/model"
	pb "getCompanies-microservice/proto"
)

func Model2Proto(company model.Company) *pb.Company {
	proto := pb.Company{
		Name:    company.Name,
		Address: company.Address,
		Oib:     company.OIB,
		Id:      company.ID,
	}
	return &proto
}
