package hello

import "testing"

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Test1",
			want: "Hello",
		},
		//Fail case
		// {
		// 	name: "Test2",
		// 	want: "Hello2",
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Hello()
			if ans != tt.want {
				t.Errorf("name: %s. got %s, want %s", tt.name, ans, tt.want)
			}
		})
	}
}
