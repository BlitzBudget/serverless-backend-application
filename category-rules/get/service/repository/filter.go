package repository

import "get-category-rule/service/models"

// Filter the category ids if they are present
func FilterByCategoryId(responseItems []*models.ResponseItem, av *models.QueryParameter) []*models.ResponseItem {
	var filteredItems []*models.ResponseItem

	if av.CategoryId == nil {
		return responseItems
	}

	for _, categoryRule := range responseItems {
		if categoryRule.CategoryId == av.CategoryId {
			filteredItems = append(filteredItems, categoryRule)
		}
	}
	return filteredItems
}
