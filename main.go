package main

import (
	"context"
	"errors"
	"fmt"

	shippy "github.com/huarngpa/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
)

// Repository is an interface that defines methods for getting vessel information
type Repository interface {
	FindAvailable(*shippy.Specification) (*shippy.Vessel, error)
}

// VesselRepository is an implementation of Repository
type VesselRepository struct {
	vessels []*shippy.Vessel
}

// FindAvailable finds a vessel that can take consignment based on the specifications
func (repo *VesselRepository) FindAvailable(spec *shippy.Specification) (*shippy.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, req *shippy.Specification, res *shippy.Response) error {
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	res.Vessel = vessel
	return nil
}

func main() {
	vessels := []*shippy.Vessel{
		&shippy.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	}

	repo := &VesselRepository{vessels}

	srv := grpc.NewService(
		micro.Name("shippy.service.vessel"),
	)

	srv.Init()

	shippy.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
