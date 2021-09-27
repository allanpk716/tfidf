package seg

import "github.com/go-ego/gse"

type JiebaTokenizer struct {
	x *gse.Segmenter
}

func NewJieba() *JiebaTokenizer {

	seg := gse.Segmenter{}
	seg.LoadDict()
	return &JiebaTokenizer{
		x: &seg,
	}
}

func (j *JiebaTokenizer) Seg(text string) []string {
	// x := gojieba.NewJieba()
	// defer x.Free()
	// fmt.Println(x.ExtractWithWeight(text, 5))
	// return x.Cut(text, true)
	return j.x.Cut(text, true)

}

func (j *JiebaTokenizer) Free() {
	if j.x != nil {
		//j.x.Free()
	}
}
