package main

import "testing"

func Test_convert(t *testing.T) {
	tests := []struct {
		word string
		want string
	}{
		{
			word: "",
			want: "",
		},
		{
			word: "a",
			want: "a",
		},
		{
			word: "word",
			want: "word",
		},
		{
			word: "localization",
			want: "l10n",
		},
		{
			word: "internationalization",
			want: "i18n",
		},
		{
			word: "pneumonoultramicroscopicsilicovolcanoconiosis",
			want: "p43s",
		},
		{
			word: "pneumonoultramicroscopicsilicovolcanoconiosispneumonoultramicroscopicsilicovolcanoconiosis",
			want: "p88s",
		},
		{
			word: "pneumonoultramicroscopicsilicovolcanoconiosispneumonoultramicroscopicsilicovolcanoconiastrarstarosis",
			want: "p98s",
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := convert(tt.word); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
