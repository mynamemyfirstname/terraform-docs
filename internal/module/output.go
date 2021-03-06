package module

import (
	"github.com/segmentio/terraform-docs/pkg/tfconf"
)

// TerraformOutput is used for unmarshalling `terraform outputs --json` into
type TerraformOutput struct {
	Sensitive bool        `json:"sensitive"`
	Type      interface{} `json:"type"`
	Value     interface{} `json:"value"`
}

type outputsSortedByName []*tfconf.Output

func (a outputsSortedByName) Len() int      { return len(a) }
func (a outputsSortedByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a outputsSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

type outputsSortedByPosition []*tfconf.Output

func (a outputsSortedByPosition) Len() int      { return len(a) }
func (a outputsSortedByPosition) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a outputsSortedByPosition) Less(i, j int) bool {
	return a[i].Position.Filename < a[j].Position.Filename || a[i].Position.Line < a[j].Position.Line
}
