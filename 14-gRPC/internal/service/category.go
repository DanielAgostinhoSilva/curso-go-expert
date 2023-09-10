package service

import (
	"context"
	"github.com/DanielAgostinhoSilva/curso-go-expert/14-gRPC/internal/database"
	"github.com/DanielAgostinhoSilva/curso-go-expert/14-gRPC/internal/pb"
	"io"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (props *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := props.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, err
}

func (props *CategoryService) ListCategory(context.Context, *pb.Blank) (*pb.CategoryList, error) {
	categories, err := props.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categoriesPB []*pb.Category
	for _, category := range categories {
		categoriesPB = append(categoriesPB, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryList{
		Categories: categoriesPB,
	}, nil
}

func (props *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := props.CategoryDB.FindByCategoryId(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, err
}

func (props *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	var categories pb.CategoryList
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&categories)
		}
		if err != nil {
			return err
		}
		categoryResult, err := props.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}
		categories.Categories = append(categories.Categories, &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
	}
}

func (props *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		categoryResult, err := props.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}
		err = stream.Send(&pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		})
		if err != nil {
			return err
		}
	}
}
