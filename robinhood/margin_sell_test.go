package robinhood

import (
	"reflect"
	"testing"
)

func Test_margin_call(t *testing.T) {
	type args struct {
		trades_to_parse [][]string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			"Test #1",
			args{[][]string{
				{"1", "AAPL", "B", "10", "10"},
				{"3", "GOOG", "B", "20", "5"},
				{"10", "AAPL", "S", "5", "15"},
			}},
			[][]string{
				{"CASH", "875"},
				{"AAPL", "5"},
				{"GOOG", "20"},
			},
		},
		{
			"Test #2",
			args{[][]string{
				{"1", "GOOG", "B", "20", "5"},
				{"3", "AAPL", "B", "10", "10"},
				{"10", "AAPL", "S", "5", "15"},
			}},
			[][]string{
				{"CASH", "875"},
				{"AAPL", "5"},
				{"GOOG", "20"},
			},
		},
		{
			"Test #3",
			args{[][]string{
				{"1", "AAPL", "B", "10", "10"},
				{"3", "GOOG", "B", "20", "5"},
				{"10", "AAPL", "S", "10", "15"},
			}},
			[][]string{
				{"CASH", "950"},
				{"GOOG", "20"},
			},
		},
		{
			"Test #4",
			args{[][]string{
				{"2", "FB", "B", "50", "13"},
				{"13", "NFLX", "B", "12", "13"},
				{"14", "FB", "S", "20", "20"},
				{"15", "TSLA", "B", "25", "11"},
				{"16", "FB", "S", "2", "10"},
				{"17", "TSLA", "B", "33", "9"},
				{"19", "TSLA", "S", "6", "14"},
				{"20", "TSLA", "S", "22", "16"},
			}},
			[][]string{
				{"CASH", "478"},
				{"FB", "28"},
				{"NFLX", "12"},
				{"TSLA", "30"},
			},
		},
		{
			"Test #5",
			args{[][]string{
				{"1", "AAPL", "B", "20", "10"},
				{"2", "GOOG", "B", "105", "9"},
			}},
			[][]string{
				{"CASH", "5"},
				{"AAPL", "5"},
				{"GOOG", "105"},
			},
		},
		{
			"Test #6",
			args{[][]string{
				{"23", "NFLX", "B", "49", "10"},
				{"24", "NFLX", "S", "1", "11"},
				{"26", "FB", "B", "48", "15"},
				{"31", "FB", "S", "28", "12"},
				{"34", "TSLA", "B", "37", "7"},
			}},
			[][]string{
				{"CASH", "88"},
				{"FB", "6"},
				{"NFLX", "48"},
				{"TSLA", "37"},
			},
		},
		{
			"Test #7",
			args{[][]string{
				{"2", "GOOG", "B", "46", "20"},
				{"6", "GOOG", "S", "37", "5"},
				{"10", "GOOG", "S", "7", "1"},
				{"13", "GOOG", "S", "1", "4"},
				{"15", "GOOG", "S", "1", "4"},
				{"17", "AAPL", "B", "18", "18"},
				{"22", "AAPL", "S", "9", "1"},
				{"23", "AAPL", "S", "4", "3"},
				{"25", "AAPL", "S", "2", "3"},
				{"28", "GOOG", "B", "36", "15"},
				{"31", "AAPL", "B", "49", "20"},
				{"32", "GOOG", "S", "2", "5"},
				{"43", "NFLX", "B", "41", "17"},
				{"46", "NFLX", "B", "21", "19"},
				{"49", "NFLX", "B", "27", "17"},
				{"50", "NFLX", "S", "1", "2"},
				{"54", "AAPL", "B", "46", "19"},
			}},
			[][]string{
				{"CASH", "2"},
			},
		},
		{
			"Test #8",
			args{[][]string{
				{"1", "AAPL", "B", "5", "100"},
				{"2", "GOOG", "B", "5", "75"},
				{"3", "AAPLO", "B", "5", "50"},
			}},
			[][]string{
				{"CASH", "25"},
				{"AAPL", "5"},
				{"AAPLO", "5"},
				{"GOOG", "3"},
			},
		},
		{
			"Test #10",
			args{[][]string{
				{"3", "AAPL", "B", "23", "16"},
				{"8", "AAPL", "S", "8", "3"},
				{"16", "AAPL", "S", "2", "5"},
				{"20", "AAPL", "B", "46", "18"},
				{"25", "AAPL", "S", "14", "5"},
				{"30", "AAPL", "S", "22", "4"},
				{"32", "AAPL", "S", "11", "4"},
				{"35", "AAPL", "S", "1", "3"},
				{"36", "AAPL", "S", "1", "1"},
				{"41", "AAPL", "S", "1", "5"},
				{"45", "FB", "B", "20", "19"},
				{"53", "FB", "S", "7", "4"},
				{"55", "GOOG", "B", "19", "15"},
				{"59", "FB", "S", "2", "3"},
				{"61", "FB", "S", "2", "1"},
				{"66", "GOOG", "S", "1", "1"},
			}},
			[][]string{
				{"CASH", "9"},
				{"GOOG", "1"},
			},
		},
		{
			"Test #9",
			args{
				[][]string{
					{"4", "AAPL", "B", "3", "10"},
					{"10", "GOOG", "B", "38", "12"},
					{"14", "AAPL", "S", "1", "17"},
					{"17", "AAPL", "S", "1", "10"},
					{"19", "GOOG", "S", "31", "12"},
					{"23", "FB", "B", "19", "18"},
					{"25", "GOOGO", "B", "4", "20"},
					{"30", "FB", "B", "20", "19"},
					{"35", "GOOG", "S", "1", "9"},
					{"40", "GOOGO", "B", "2", "5"},
					{"46", "AAPL", "B", "26", "13"},
					{"50", "GOOGO", "S", "6", "7"},
					{"55", "GOOGO", "B", "5", "11"},
					{"59", "GOOGO", "S", "2", "10"},
					{"61", "FBO", "B", "21", "20"},
				},
			},
			[][]string{
				{"CASH", "6"},
				{"AAPL", "27"},
				{"FB", "26"},
				{"FBO", "1"},
				{"GOOG", "6"},
				{"GOOGO", "3"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := margin_call(tt.args.trades_to_parse); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("margin_call() = %v, want %v", got, tt.want)
			}
		})
	}
}
