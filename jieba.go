package gojieba

/*
#cgo CXXFLAGS: -I./deps/cppjieba/include -I./deps/cppjieba/deps/limonp/include -DLOGGING_LEVEL=LL_WARNING -O3 -Wno-deprecated -Wno-unused-variable -std=c++11
#include <stdlib.h>
#include "jieba.h"
*/
import "C"

import (

	// These blank imports ensure `go mod vendor` copies the C++ header files
	// from deps/cppjieba and deps/limonp into the vendor directory, so that
	// CGo builds work correctly when using `go mod vendor`. They also ensure
	// the dictionary data files are included in the vendor directory for the
	// default dictionary paths to work at runtime.
	_ "github.com/yanyiwu/gojieba/deps/cppjieba/deps/limonp/include/limonp"
	_ "github.com/yanyiwu/gojieba/deps/cppjieba/dict"
	_ "github.com/yanyiwu/gojieba/deps/cppjieba/dict/pos_dict"
	_ "github.com/yanyiwu/gojieba/deps/cppjieba/include/cppjieba"
)

type TokenizeMode int

const (
	DefaultMode TokenizeMode = iota
	SearchMode
)

type Word struct {
	Str   string
	Start int
	End   int
}

type Jieba struct {
	jieba C.Jieba
	freed int32
}

func NewJieba(paths ...string) *Jieba { _ = "STUB: not implemented"; return nil }

// check if the dictionary files exist

// set finalizer to free the memory when the object is garbage collected

func (x *Jieba) Free() { _ = "STUB: not implemented"; return }

// only free once

// Deprecated: Use Free() instead. Free() now calls Trim() automatically.
func (x *Jieba) FreeWithTrim() {
	_ = "STUB: not implemented"

	// Deprecated: WithTrim is no longer necessary; Free() now calls Trim()
	// automatically on Linux. Calling this method is a no-op.
	return
}

func (x *Jieba) WithTrim() *Jieba { _ = "STUB: not implemented"; return nil }

func (x *Jieba) Cut(s string, hmm bool) []string { _ = "STUB: not implemented"; return nil }

// can directly use free now...

func (x *Jieba) CutAll(s string) []string { _ = "STUB: not implemented"; return nil }

func (x *Jieba) CutForSearch(s string, hmm bool) []string { _ = "STUB: not implemented"; return nil }

func (x *Jieba) Tag(s string) []string { _ = "STUB: not implemented"; return nil }

func (x *Jieba) AddWord(s string) { _ = "STUB: not implemented"; return }

func (x *Jieba) AddWordEx(s string, freq int, tag string) { _ = "STUB: not implemented"; return }

func (x *Jieba) RemoveWord(s string) { _ = "STUB: not implemented"; return }

func (x *Jieba) Tokenize(s string, mode TokenizeMode, hmm bool) []Word {
	_ = "STUB: not implemented"
	return nil
}

type WordWeight struct {
	Word   string
	Weight float64
}

func (x *Jieba) Extract(s string, topk int) []string { _ = "STUB: not implemented"; return nil }

func (x *Jieba) ExtractWithWeight(s string, topk int) []WordWeight {
	_ = "STUB: not implemented"
	return nil
}

func cwordweights(x *C.struct_CWordWeight) []WordWeight { _ = "STUB: not implemented"; return nil }

// convertCWordToSlice convert *C.Word to []string (zero-copy)
func convertCWordToSlice(s string, x *C.Word) []string { _ = "STUB: not implemented"; return nil }

// 假设 C++ 返回以 {0,0} 结尾的哨兵

// convertCWordToStructs convert *C.Word to []Word (Go Struct)
func convertCWordToStructs(s string, x *C.Word) []Word { _ = "STUB: not implemented"; return nil }

func Trim() { _ = "STUB: not implemented"; return }
