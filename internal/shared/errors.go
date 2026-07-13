package shared

import "fmt"

type Not_fount_error struct {
	Resource string
	Key      any
}

func New_not_fount_error(resource string, key any) *Not_fount_error {
	return &Not_fount_error{Resource: resource, Key: key}
}

func (e *Not_fount_error) Error() string {
	return fmt.Sprintf("object %q with ket %v does not found", e.Resource, e.Key)
}
