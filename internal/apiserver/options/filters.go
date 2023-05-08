package options

import (
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

// FilteredListOptions contains options for filtering lists of items.
type FilteredListOptions struct {
	metav1.ListOptions

	// Filter allows filtering the list of items based on custom criteria.
	Filter string `json:"filter,omitempty" form:"filter"`
}
