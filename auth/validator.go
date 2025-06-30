package auth

import "project/common"

func ValidateCreateRequest(request RegisterRequest) error {
	if err := common.ValidateRequest(request); err != nil {
		return err
	}
	return nil
}

