package gojieba

/*
#include "jieba.h"
*/
import "C"

func isDirExists(path string) bool { _ = "STUB: not implemented"; return false }

func cstrings(x **C.char) []string { _ = "STUB: not implemented"; return nil }

func convertWords(s string, words *C.Word) []Word { _ = "STUB: not implemented"; return nil }

//func cwordweights(x unsafe.Pointer) []WordWeight {
//	var s []WordWeight
//	eltSize := 16
//	for (*(*C.char))(x) != nil {
//		ww := WordWeight{
//			C.GoString((*C.char))(x)),
//			(*x).weight,
//		}
//		s = append(s, ww)
//		x = (*C.struct_CWordWeight)(unsafe.Pointer(uintptr(unsafe.Pointer(x)) + eltSize))
//	}
//	return s
//}
