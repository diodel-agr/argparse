package main

import "container/list"

func (s Specifier) createSlice() interface{} {
	if s.theType == "i" {
		switch s.bitSize {
		case 0:
			return []int{}
		case 8:
			return []int8{}
		case 16:
			return []int16{}
		case 32:
			return []int32{}
		case 64:
			return []int64{}
		}
	} else if s.theType == "ui" {
		switch s.bitSize {
		case 0:
			return []uint{}
		case 8:
			return []uint8{}
		case 16:
			return []uint16{}
		case 32:
			return []uint32{}
		case 64:
			return []uint64{}
		}
	} else if s.theType == "r" {
		return []rune{}
	} else if s.theType == "by" {
		return []byte{}
	} else if s.theType == "f" {
		switch s.bitSize {
		case 32:
			return []float32{}
		case 64:
			return []float64{}
		}
	} else if s.theType == "c" {
		switch s.bitSize {
		case 64:
			return []complex64{}
		case 128:
			return []complex128{}
		}
	} else if s.theType == "b" {
		return []bool{}
	} else if s.theType == "s" {
		return []string{}
	}
	return nil
}

// i, ui, r, by
// f
// c
// b, s
func (s Specifier) convertSizedSlice(str []string, ai *int, ls *list.List) string {
	// create the slice.
	slice := s.createSlice()
	for i := uint(0); i < s.sliceSize; i++ {
		result, err := s.convert(str, ai, ls)
		if err != "" {
			return err
		}
		slice = append(slice, result)
	}
	ls.PushBack(slice)
	return ""
}
