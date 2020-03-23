package main

import (
	"container/list"
	"reflect"
)

func (s Specifier) createSlice() *list.List {
	result := list.New()
	if s.theType == "i" {
		switch s.bitSize {
		case 0:
			result.PushBack([]int{})
		case 8:
			result.PushBack([]int8{})
		case 16:
			result.PushBack([]int16{})
		case 32:
			result.PushBack([]int32{})
		case 64:
			result.PushBack([]int64{})
		}
	} else if s.theType == "ui" {
		switch s.bitSize {
		case 0:
			result.PushBack([]uint{})
		case 8:
			result.PushBack([]uint8{})
		case 16:
			result.PushBack([]uint16{})
		case 32:
			result.PushBack([]uint32{})
		case 64:
			result.PushBack([]uint64{})
		}
	} else if s.theType == "r" {
		result.PushBack([]rune{})
	} else if s.theType == "by" {
		result.PushBack([]byte{})
	} else if s.theType == "f" {
		switch s.bitSize {
		case 32:
			result.PushBack([]float32{})
		case 64:
			result.PushBack([]float64{})
		}
	} else if s.theType == "c" {
		switch s.bitSize {
		case 64:
			result.PushBack([]complex64{})
		case 128:
			result.PushBack([]complex128{})
		}
	} else if s.theType == "b" {
		result.PushBack([]bool{})
	} else if s.theType == "s" {
		result.PushBack([]string{})
	}
	return result
}

// convertSlice - function used to convert the @slice variable (which is of reflect.Value type)
// to the type described by the @s specifier.
func (s Specifier) convertSlice(slice reflect.Value, ls *list.List) {
	if s.theType == "i" {
		switch s.bitSize {
		case 0:
			result := slice.Interface().([]int)
			ls.PushBack(result)
		case 8:
			result := slice.Interface().([]int8)
			ls.PushBack(result)
		case 16:
			result := slice.Interface().([]int16)
			ls.PushBack(result)
		case 32:
			result := slice.Interface().([]int32)
			ls.PushBack(result)
		case 64:
			result := slice.Interface().([]int64)
			ls.PushBack(result)
		}
	} else if s.theType == "ui" {
		switch s.bitSize {
		case 0:
			result := slice.Interface().([]uint)
			ls.PushBack(result)
		case 8:
			result := slice.Interface().([]uint8)
			ls.PushBack(result)
		case 16:
			result := slice.Interface().([]uint16)
			ls.PushBack(result)
		case 32:
			result := slice.Interface().([]uint32)
			ls.PushBack(result)
		case 64:
			result := slice.Interface().([]uint64)
			ls.PushBack(result)
		}
	} else if s.theType == "r" {
		result := slice.Interface().([]rune)
		ls.PushBack(result)
	} else if s.theType == "by" {
		result := slice.Interface().([]byte)
		ls.PushBack(result)
	} else if s.theType == "f" {
		switch s.bitSize {
		case 32:
			result := slice.Interface().([]float32)
			ls.PushBack(result)
		case 64:
			result := slice.Interface().([]float64)
			ls.PushBack(result)
		}
	} else if s.theType == "c" {
		switch s.bitSize {
		case 64:
			result := slice.Interface().([]complex64)
			ls.PushBack(result)
		case 128:
			result := slice.Interface().([]complex128)
			ls.PushBack(result)
		}
	} else if s.theType == "b" {
		result := slice.Interface().([]bool)
		ls.PushBack(result)
	} else if s.theType == "s" {
		result := slice.Interface().([]string)
		ls.PushBack(result)
	}
}

func (s Specifier) convertSizedSlice(str []string, ai *int, ls *list.List) string {
	// create the slice.
	slice := reflect.ValueOf(s.createSlice().Front().Value)
	for i := uint(0); i < s.sliceSize; i++ {
		err := s.convert(str, ai, ls)
		if err != "" {
			return err
		}
		result := reflect.ValueOf(ls.Remove(ls.Back()))
		slice = reflect.Append(slice, result)
	}
	s.convertSlice(slice, ls)
	return ""
}
