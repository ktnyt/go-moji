package moji

import (
	"testing"
)

type testCase struct {
	n string
	s string
	e string
	f Dictionary
	t Dictionary
}

var cases = []testCase{
	// RangeDictionary test cases
	testCase{
		n: "ZE to HE",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "AZaz09.　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		f: ZE,
		t: HE,
	},
	testCase{
		n: "HE to ZE",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "ＡＺａｚ０９．　ＡＺａｚ０９． 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		f: HE,
		t: ZE,
	},
	testCase{
		n: "HG to KK",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "ＡＺａｚ０９．　AZaz09. 「アン゛ガパ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		f: HG,
		t: KK,
	},
	testCase{
		n: "KK to HG",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」あん゜がぱ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		f: KK,
		t: HG,
	},
	testCase{
		n: "ZK to HK",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "ＡＺａｚ０９．　AZaz09. ｢あんﾞがぱ｣ｱﾝﾟｶﾞﾊﾟ｡｢ｱﾝﾞｶﾞﾊﾟ｣",
		f: ZK,
		t: HK,
	},
	testCase{
		n: "HK to ZK",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。「アン゛ガパ」",
		f: HK,
		t: ZK,
	},
	testCase{
		n: "ZS to HS",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "ＡＺａｚ０９． AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		f: ZS,
		t: HS,
	},
	testCase{
		n: "HS to ZS",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "ＡＺａｚ０９．　AZaz09.　「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		f: HS,
		t: ZS,
	},
}

func TestConvert(t *testing.T) {
	for _, c := range cases {
		a := Convert(c.s, c.f, c.t)
		if c.e != a {
			t.Fatalf("convert '%s' failed:\n  expected '%s'\n  but got  '%s'", c.n, c.e, a)
		}
	}
}
