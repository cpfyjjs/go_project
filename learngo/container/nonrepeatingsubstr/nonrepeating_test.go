package nonrepeatingsubstr

import "testing"

func TestSubstr(t *testing.T){
	tests :=[]struct{
		s string
		ans int
	}{
		// Normal cases
		{"abcabcbb",3},
		{"pwwkew",3},

		// Edge cases
		{"",0},
		{"b",1},
		{"bbbbbbbb",1},
		{"abcabcabcd",4},
	}


	for _,tt := range(tests){
		if acctual := LengthOfNonRepeatingSubstr(tt.s);acctual!= tt.ans{
			t.Errorf("LengthOfNonRepatingSubstr(%s) got %d,except %d",tt.s,acctual,tt.ans)
		}
	}
}


func BenchmarkLengthOfNonRepeatingSubstr(b *testing.B) {
	s,ans := "化肥会挥发 黑化肥发灰,灰化肥发黑 黑化肥发灰会挥发;灰化肥发挥会发黑,黑化肥挥发发灰会挥发;灰化肥挥发发黑会发挥 黑灰化肥会挥发发",8
	b.ResetTimer()

	for i:=0;i<b.N;i++{
		if acctual := LengthOfNonRepeatingSubstr(s);acctual!= ans{
			b.Errorf("LengthOfNonRepatingSubstr(%s) got %d,except %d",s,acctual,ans)
		}
	}
}





