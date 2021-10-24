package read

import (
	"reflect"
	"testing"
)

func TestFeed_ReadFeedByFile(t *testing.T) {
	type fields struct {
		Name string
		URI  string
		Type string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*Feed
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Feed{
				Name: tt.fields.Name,
				URI:  tt.fields.URI,
				Type: tt.fields.Type,
			}
			got, err := f.ReadFeedByFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFeedByFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFeedByFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
