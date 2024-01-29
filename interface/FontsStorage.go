package _interface

import "github.com/golang/freetype/truetype"


type FontsStorage interface {
	LoadFontByName(name string) *truetype.Font
	LoadFontByPath(path string) []*truetype.Font
	LoadFontsByNames(assetFontNames []string) []*truetype.Font
}

