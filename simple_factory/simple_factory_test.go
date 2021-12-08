package simple_factory

import (
	"reflect"
	"testing"
)

func TestNewConfigFileParser(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name     string
		filename string
		args     args
		want     ConfigFileParser
	}{
		{
			name:     "json",
			filename: "config.json",
			args:     args{t: "json"},
			want:     JsonConfigParser{},
		},
		{
			name:     "yaml",
			filename: "config.ymal",
			args:     args{t: "yaml"},
			want:     YamlConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigFileParser(tt.args.t, tt.filename); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIRuleConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
