package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/jose-manaloto/joseapp/graph/generated"
	"github.com/jose-manaloto/joseapp/graph/model"
)

func (r *mutationResolver) AddHouse(ctx context.Context, input model.HouseInput) (*model.House, error) {
	var house model.House
	house.RealtorName = input.RealtorName
	house.Address = input.Address
	err := r.DB.Create(&house).Error
	if err != nil {
		return nil, err
	}
	return &house, nil
}

func (r *mutationResolver) UpdateHouse(ctx context.Context, houseID int, input model.HouseInput) (*model.House, error) {
	var updatedHouse model.House
	updatedHouse.RealtorName = input.RealtorName
	updatedHouse.Address = input.Address
	updatedHouse.Issues = mapIssuesFromInput(input.Issues)
	updatedHouse.ID = houseID
	r.DB.Save(&updatedHouse)
	return &updatedHouse, nil
}

func (r *mutationResolver) RemoveHouse(ctx context.Context, houseID int) (bool, error) {
	r.DB.Where("house_id = ?", houseID).Delete(&model.Issue{})
	r.DB.Where("house_id = ?", houseID).Delete(&model.House{})
	return true, nil
}

func (r *queryResolver) Houses(ctx context.Context) ([]*model.House, error) {
	var orders []*model.House
	r.DB.Preload("Issues").Find(&orders)

	return orders, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func mapIssuesFromInput(issueInput []model.IssueInput) []model.Issue {
	var issues []model.Issue
	for _, issueInput := range issueInput {
		issues = append(issues, model.Issue{
			IssueTitle:       issueInput.IssueTitle,
			IssueDescription: issueInput.IssueDescription,
			Urgent:           issueInput.Urgent,
		})
	}
	return issues
}
