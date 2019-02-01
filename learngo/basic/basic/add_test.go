package basic

import "testing"

func TestAdd(t *testing.T){
	tests :=[]struct{a,b,c int64}{
		{1,2,3},
		{2,2,4},
		{3,4,5},
		{4,-2,2},
	}

	for _,tt :=range tests{
		if actual:= Add(tt.a,tt.b);actual!=tt.c{
			t.Errorf("Add (%d ，%d) got %d ,except %d",tt.a,tt.b,actual,tt.c)
		}
	}
}
