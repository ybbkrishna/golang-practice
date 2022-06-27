package stack

import (
	"reflect"
	"testing"
)

func TestStack_Push(t *testing.T) {
	type fields struct {
		head     *node
		length   int
		Capacity int
	}
	type args struct {
		elem interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"[success] Null push",

		},
		{

		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				head:     tt.fields.head,
				length:   tt.fields.length,
				Capacity: tt.fields.Capacity,
			}
			got, err := s.Push(tt.args.elem)
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Push() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Stack.Push() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	type fields struct {
		head     *node
		length   int
		Capacity int
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stack{
				head:     tt.fields.head,
				length:   tt.fields.length,
				Capacity: tt.fields.Capacity,
			}
			got, err := s.Pop()
			if (err != nil) != tt.wantErr {
				t.Errorf("Stack.Pop() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Stack.Pop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_Len(t *testing.T) {
	type fields struct {
		head     *node
		length   int
		Capacity int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack{
				head:     tt.fields.head,
				length:   tt.fields.length,
				Capacity: tt.fields.Capacity,
			}
			if got := s.Len(); got != tt.want {
				t.Errorf("Stack.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStack_String(t *testing.T) {
	type fields struct {
		head     *node
		length   int
		Capacity int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Stack{
				head:     tt.fields.head,
				length:   tt.fields.length,
				Capacity: tt.fields.Capacity,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("Stack.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
