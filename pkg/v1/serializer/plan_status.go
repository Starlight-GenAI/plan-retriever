package serializer

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

type PlanStatusRequest struct {
	ID string `json:"id"`
}

type PlanStatusResponse struct {
	Status string `json:"status"`
}

func (b PlanStatusRequest) Validate() error {

	if err := v.ValidateStruct(&b, v.Field(&b.ID, v.Required)); err != nil {
		return err
	}

	return nil
}
