package bitset
// author : zicheng.lang @ blockin.com
// date   : 09/28 2018
type bitsMap []byte

type BitSet struct {
	size uint64
	data *bitsMap
}

const (
	OneBitSize = 8
	OneBitSizeBinary = 3 //0000 1000
)

// [0000 0000] [0000 0000] [0000 0000] [0000 0000]
//  0       7   8       15

func New(size uint64) *BitSet {
	realIndex := size - 1
	IndexNum := (realIndex >> OneBitSizeBinary) + 1

	bitsMap := make(bitsMap,IndexNum, IndexNum)
	b := BitSet{size, &bitsMap}
	return &b
}

func getPosIndex(pos uint64) uint64 {
	return pos >> OneBitSizeBinary
}

func (b *BitSet) SetTrue (pos uint64) {
	index := getPosIndex(pos)
	(*b.data)[index]  |= byte(0x01) << (pos % OneBitSize)
}

func (b *BitSet) SetFalse(pos uint64){
	index := getPosIndex(pos)
	(*b.data)[index]  &= ^(byte(0x01) << (pos % OneBitSize))
}

func (b *BitSet) IsTrue (pos uint64) bool {
	index := getPosIndex(pos)
	isTrue := ((*b.data)[index] & (byte(0x01) << (pos % OneBitSize) )) != 0
	return isTrue
}

func (b *BitSet) SearchFalsePos() ([]uint64) {
	falseList := make([]uint64,0)
	for pos := uint64(0); pos < b.size  ; pos++ {
		if !b.IsTrue(pos) {
			falseList = append(falseList, pos)
		}
	}
	return falseList
}

func (b *BitSet) SearchTruePos() ([]uint64) {
	trueList := make([]uint64,0)
	for pos := uint64(0); pos < b.size  ; pos++ {
		if b.IsTrue(pos) {
			trueList = append(trueList, pos)
		}
	}
	return trueList
}





