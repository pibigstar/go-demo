package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-ego/gse/hmm/idf"

	"github.com/go-ego/gse"
)

var (
	test1 = `["小梅十字绣主播买十字绣，绣完可以回收，我们就买了但是现在找不到人了","不是商品出现问题，是商家出现问题了","商家失联，拿高回收诱惑我们买图，现在找不到商家，我要求退款"]`
)
var (
	// 只要中文和英文
	reg = regexp.MustCompile("[\u4e00-\u9fa5a-zA-Z]+")
)

func main() {
	var seg1 gse.Segmenter

	test1 = strings.ReplaceAll(test1, "\"", "")
	if err := seg1.LoadDict(); err != nil {
		panic(err)
	}
	if err := seg1.LoadStop(); err != nil {
		panic(err)
	}
	//seg1.LoadDict("/Users/bytedance/work/my/tools/fenci/gse/data/t_1.txt")
	//seg1.LoadDict("zh_s")
	//seg1.LoadDictEmbed("zh_s")

	text := reg.FindAllString(test1, -1)
	var s string
	for _, v := range text {
		s += v
	}

	s1 := seg1.Cut(s, true)
	s2 := seg1.Cut(s, false)
	s3 := seg1.CutAll(s)

	fmt.Printf("seg Cut1: %s \n", strings.Join(s1, "/"))
	fmt.Printf("seg Cut2: %s \n", strings.Join(s2, "/"))
	fmt.Printf("seg Cut3: %s \n", strings.Join(s3, "/"))

	extAndRank(seg1, test1)
}

func extAndRank(segs gse.Segmenter, text string) {
	var te idf.TagExtracter
	te.WithGse(segs)
	err := te.LoadIdf()
	fmt.Println("load idf: ", err)

	segments := te.ExtractTags(text, 20)
	fmt.Println("segments: ", len(segments), segments)
	// segments:  5 [{科幻片 1.6002581704125} {全片 1.449761569875} {摄影机 1.2764747747375} {拍摄 0.9690261695075} {制作 0.8246043033375}]

	var tr idf.TextRanker
	tr.WithGse(segs)

	results := tr.TextRank(text, 5)
	fmt.Println("results: ", results)
	// results:  [{机 1} {全片 0.9931964427972227} {摄影 0.984870660504368} {使用 0.9769826633059524} {是 0.8489363954683677}]
}
