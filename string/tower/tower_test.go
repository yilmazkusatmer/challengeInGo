package tower

import (
	"reflect"
	"testing"
)

type args struct {
	height int
}

var tests = []struct {
	name     string
	height   int
	expected []string
}{
	{"", 2, []string{"   |", "  *|*", " **|**", "-------"}},
	{"", 3, []string{"    |", "   *|*", "  **|**", " ***|***", "---------"}},
}

func TestBuildTower(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := BuildTower(tt.height); !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BuildTower() produce %v, but expected %v", result, tt.expected)
			}
		})
	}
}
