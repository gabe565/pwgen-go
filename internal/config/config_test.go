package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProfileRef_MarshalText(t *testing.T) {
	type fields struct {
		Name  string
		Param int
	}
	tests := []struct {
		fields  fields
		want    string
		wantErr require.ErrorAssertionFunc
	}{
		{fields{Name: "diceware"}, "diceware", require.NoError},
		{fields{Name: "diceware", Param: 4}, "diceware:4", require.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			p := &ProfileRef{Name: tt.fields.Name, Param: tt.fields.Param}
			got, err := p.MarshalText()
			tt.wantErr(t, err)
			assert.EqualValues(t, tt.want, got)
		})
	}
}

func TestProfileRef_UnmarshalText(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		args    args
		want    ProfileRef
		wantErr require.ErrorAssertionFunc
	}{
		{args{"diceware"}, ProfileRef{Name: "diceware"}, require.NoError},
		{args{"diceware:4"}, ProfileRef{Name: "diceware", Param: 4}, require.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.args.text, func(t *testing.T) {
			var p ProfileRef
			tt.wantErr(t, p.UnmarshalText([]byte(tt.args.text)))
			assert.Equal(t, tt.want, p)
		})
	}
}
