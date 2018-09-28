package bitset

import (
	"fmt"
	"log"
	"testing"
	"time"
)

const TestBitsetSize = 65

const (
	TestPos0 = iota
	TestPos1
	TestPos2
	TestPos3
	TestPos4
	TestPos5
	TestPos6
	TestPos7
	TestPosTotal
)

const BenchMarkLoopTime = 15
func TestNew(t *testing.T) {
	bb := New(TestBitsetSize)
	if bb == nil {
		t.Error("new bitset error.")
	}

	if bb.size != TestBitsetSize {
		t.Error("illegal bitset size.")
	}
}

func TestBitSet_SetTrue(t *testing.T) {
	bb := New(TestBitsetSize)
	samble := uint64(0x01)
	for i:= uint64(0); i< TestBitsetSize; i++ {
		bb.SetTrue(i)
		p := (samble << i)
		p |= 0x01
		samble = p
		log.Println(bb.data,i,p)
	}
}

func TestBitSet_IsTrue(t *testing.T) {
	bb := New(TestBitsetSize)
	for i := uint64(0); i< TestBitsetSize - 1 ; i ++ {
		bb.SetTrue(i)
		if ! bb.IsTrue(i) {
			t.Errorf("error in pos %d, should be true",i)
		}
	}
}

func TestBitSet_SearchFalsePos(t *testing.T) {
	// full
	bb := New(TestBitsetSize)
	for i := uint64(0); i< TestBitsetSize  ; i = i + 1 {
		bb.SetTrue(i)
		if ! bb.IsTrue(i) {
			t.Errorf("error in pos %d, should be true",i)
		}
	}
	list := bb.SearchFalsePos()
	if len(list) != 0 {
		t.Errorf("bitset should be full, but is not. %v",list)
	}

	// half

	bb2 := New(TestBitsetSize)
	for i := uint64(0); i< TestBitsetSize - 1 ; i = i + 2 {
		bb2.SetTrue(i)
		if ! bb2.IsTrue(i) {
			t.Errorf("error in pos %d, should be true",i)
		}
	}
	list2 := bb2.SearchFalsePos()

	if len(list2) == 0 {
		t.Errorf("bitset should not be full, but is not. %v",list2)
	}


	bb3 := New(TestPosTotal)

	bb3.SetTrue(TestPos0)
	list3 := bb3.SearchFalsePos()
	for _,x := range list3 {
		if x == TestPos0 {
			t.Errorf("This pos :%d should not be false",x)
		}
	}

	bb3.SetTrue(TestPos1)
	list3 = bb3.SearchFalsePos()
	for _,x := range list3 {
		if x == TestPos1 || x == TestPos0{
			t.Errorf("This pos :%d should not be false",x)
		}
	}


	bb3.SetTrue(TestPos2)
	list3 = bb3.SearchFalsePos()
	for _,x := range list3 {
		if x == TestPos0 || x == TestPos1 || x == TestPos2{
			t.Errorf("This pos :%d should not be false",x)
		}
	}
}

func TestBitSet_SetFalse(t *testing.T) {

	bb := New(TestBitsetSize)
	for i := uint64(0); i< TestBitsetSize  ; i = i + 1 {
		bb.SetTrue(i)
		if ! bb.IsTrue(i) {
			t.Errorf("error in pos %d, should be true",i)
		}
	}

	for i := uint64(0); i< TestBitsetSize  ; i = i + 1 {
		bb.SetFalse(i)
		if bb.IsTrue(i) {
			t.Errorf("error in pos %d, should be false",i)
		}
	}

	falseList := bb.SearchFalsePos()
	if uint64(len(falseList) )!= bb.size {
		t.Errorf("all pos should be false, but not. falselist is : %v", falseList)
	}



	bb3 := New(TestPosTotal)
	for i := uint64(0); i< TestPosTotal  ; i = i + 1 {
		bb3.SetTrue(i)
	}

	if len(bb3.SearchFalsePos()) != 0 {
		t.Errorf("false list should all empty,but not %v", bb3.SearchFalsePos())
	}

	bb3.SetFalse(TestPos0)
	list3 := bb3.SearchFalsePos()

	testPass := false
	for _,x := range list3 {
		if x == TestPos0 {
			testPass = true
		}
	}
	if !testPass {
		t.Error("This pos :")
	}

}

func BenchmarkBitSet_SearchFalsePos(b *testing.B) {

	for Size := uint64(0); Size < uint64(BenchMarkLoopTime); Size ++ {
		bs := New(0x8000 << Size)
		for pos := uint64(0); pos < bs.size; pos += 2 {
			bs.SetTrue(pos)
		}

		b.Run(fmt.Sprintf("%d",0x8000 << Size), func(b *testing.B) {
			//b.ResetTimer()
			falseList := bs.SearchFalsePos()
			if uint64(len(falseList)) < (bs.size >> 1) {
				b.Errorf("false list len error : %d", len(falseList))
			}
			//b.StopTimer()
		})

		time.Sleep(time.Second *1)
	}
}

func BenchmarkBitSet_SetTrue(b *testing.B) {
	for Size := uint64(0); Size < uint64(BenchMarkLoopTime); Size ++ {
		bs := New(0x8000 << Size)

		b.Run(fmt.Sprintf("size : %d",0x8000 << Size), func(b *testing.B) {
			for pos := uint64(0); pos < bs.size; pos += 1 {
				bs.SetTrue(pos)
			}
		})

		time.Sleep(time.Second *1)
	}
}

func BenchmarkBitSet_IsTrue(b *testing.B) {
	for Size := uint64(0); Size < uint64(BenchMarkLoopTime); Size ++ {
		bs := New(0x8000 << Size)

		b.Run(fmt.Sprintf("size : %d",0x8000 << Size), func(b *testing.B) {
			for pos := uint64(0); pos < bs.size; pos += 1 {
				bs.IsTrue(pos)
			}
		})

		time.Sleep(time.Second *1)
	}
}
