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
	{
		n: "ZE to HE",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "AZaz09.　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		f: ZE,
		t: HE,
	},
	{
		n: "HE to ZE",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "ＡＺａｚ０９．　ＡＺａｚ０９． 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		f: HE,
		t: ZE,
	},
	{
		n: "HG to KK",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "ＡＺａｚ０９．　AZaz09. 「アン゛ガパ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		f: HG,
		t: KK,
	},
	{
		n: "KK to HG",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」あん゜がぱ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		f: KK,
		t: HG,
	},
	{
		n: "ZK to HK",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "ＡＺａｚ０９．　AZaz09. ｢あんﾞがぱ｣ｱﾝﾟｶﾞﾊﾟ｡｢ｱﾝﾞｶﾞﾊﾟ｣",
		f: ZK,
		t: HK,
	},
	{
		n: "HK to ZK",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。「アン゛ガパ」",
		f: HK,
		t: ZK,
	},
	{
		n: "ZS to HS",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "ＡＺａｚ０９． AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		f: ZS,
		t: HS,
	},
	{
		n: "HS to ZS",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		e: "ＡＺａｚ０９．　AZaz09.　「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟ｣",
		f: HS,
		t: ZS,
	},
	{
		n: "HK to UHK",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟｬｭｮｯ｣",
		e: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟﾔﾕﾖﾂ｣",
		f: HK,
		t: UHK,
	},
	{
		n: "ZK to UHK",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパャュョッ。｢ｱﾝﾞｶﾞﾊﾟｬｭｮｯ｣",
		e: "ＡＺａｚ０９．　AZaz09. ｢あんﾞがぱ｣ｱﾝﾟｶﾞﾊﾟﾔﾕﾖﾂ｡｢ｱﾝﾞｶﾞﾊﾟｬｭｮｯ｣",
		f: ZK,
		t: UHK,
	},
	{
		n: "HK to UZK",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。｢ｱﾝﾞｶﾞﾊﾟｬｭｮｯ｣",
		e: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパ。「アン゛ガパヤユヨツ」",
		f: HK,
		t: UZK,
	},
	{
		n: "ZK to UZK",
		s: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパャュョッ。｢ｱﾝﾞｶﾞﾊﾟｬｭｮｯ｣",
		e: "ＡＺａｚ０９．　AZaz09. 「あん゛がぱ」アン゜ガパヤユヨツ。｢ｱﾝﾞｶﾞﾊﾟｬｭｮｯ｣",
		f: ZK,
		t: UZK,
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
