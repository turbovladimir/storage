package repository

type PhoneChecker struct {
}

type PhoneCheckResult struct {
	ExcludedOfferIds []int `json:"excluded_offer_ids"`
}

func (c *PhoneChecker) Check(phone string) (*PhoneCheckResult, error) {
	var result PhoneCheckResult
	mock := []int{123, 345, 567}
	result.ExcludedOfferIds = append(result.ExcludedOfferIds, mock...)

	return &result, nil
}
