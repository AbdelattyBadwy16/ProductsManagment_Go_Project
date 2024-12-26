package tag

import "project/common"

func ValidateCreateRequest(request CreateRequest) error {
	if err := common.ValidateRequest(request); err != nil {
		return err
	}
	return nil
}

func ValidateDeleteRequest(request DeleteRequest) error {
	if err := common.ValidateRequest(request); err != nil {
		return err
	}
	return nil
}

func ValidateUpdateRequest(request UpdateRequest) error {
	if err := common.ValidateRequest(request); err != nil {
		return err
	}
	return nil
}
