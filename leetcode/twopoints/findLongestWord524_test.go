package twopoints

import "testing"

func Test_findLongestWord(t *testing.T) {
	type args struct {
		s          string
		dictionary []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{s: "abpcplea", dictionary: []string{"ale", "apple", "monkey", "plea"}}, want: "apple"},
		{name: "test", args: args{s: "abpcplea", dictionary: []string{"a", "b", "c"}}, want: "a"},
		{name: "test", args: args{s: "abpcplea", dictionary: []string{"apple", "abpcllllll"}}, want: "apple"},
		{name: "test", args: args{s: "yzzzbfswsnfktmfselqavhbgujptpxijmtobyjbtvvxxlsesddxmioknhxfyftytsscvkvjoyiifuvrmtbpbxfjjhbxtygifrxznfgxzrmiknebavxjzbswrhlmljyemdamsyasopttbmwqthallxbedmwkkbsqgochltnbbtnzcjeyrlgwidmvtxprnvfmuudechshvgzybkqrfllsbtprhzlysxievxtxilaigschxhomjhzmjvnftytnzjtdwkxlypxvgndnjphrvuxcdhpralpihtaembgsmyfxtkrqywhmbprxtiywajcogbdoddbyvfxtaaxzszqiuebxmvxrbycmbqvcylnqsdnrnhuofnvjswaeipptcejipelqlqlwjuqqywjnjzsvqddbpbwedtlczzkkkekrkihbnlvvwotfplfxlvofogmsuadaxvnqnagqshbjhwsbuihgjdxkrsrxjmwznpawvioswcdejhfdjkxzo",
			dictionary: []string{"skgpdpgxysovvvtoxzssdfqbe", "wnkfewhtnfzxcttslziqwjwo"}}, want: "skgpdpgxysovvvtoxzssdfqbe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLongestWord(tt.args.s, tt.args.dictionary); got != tt.want {
				t.Errorf("findLongestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
