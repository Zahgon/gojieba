package gojieba

import (
	"path"
)

var (
	DICT_DIR        string
	DICT_PATH       string
	HMM_PATH        string
	USER_DICT_PATH  string
	IDF_PATH        string
	STOP_WORDS_PATH string
)

func init() {
	DICT_DIR = path.Join(path.Dir(getCurrentFilePath()), "deps/cppjieba/dict")
	DICT_PATH = path.Join(DICT_DIR, "jieba.dict.utf8")
	HMM_PATH = path.Join(DICT_DIR, "hmm_model.utf8")
	USER_DICT_PATH = path.Join(DICT_DIR, "user.dict.utf8")
	IDF_PATH = path.Join(DICT_DIR, "idf.utf8")
	STOP_WORDS_PATH = path.Join(DICT_DIR, "stop_words.utf8")
}

const TOTAL_DICT_PATH_NUMBER = 5

func getDictPaths(args ...string) [TOTAL_DICT_PATH_NUMBER]string {
	_ = "STUB: not implemented"
	return nil
}

func getCurrentFilePath() string { _ = "STUB: not implemented"; return "" }
