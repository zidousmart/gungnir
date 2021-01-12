package http

import (
	"github.com/gin-gonic/gin/binding"
)

type GHttpServerOption interface {
	apply(*GHttpEngine)
}

type optionFunc func(*GHttpEngine)

func (f optionFunc) apply(he *GHttpEngine) {
	f(he)
}

// EnableDecoderUseNumber is used to call the UseNumber method on the JSON
// Decoder instance. UseNumber causes the Decoder to unmarshal a number into an
// interface{} as a Number instead of as a float64.
func EnableDecoderUseNumber() GHttpServerOption {
	return optionFunc(func(he *GHttpEngine) {
		binding.EnableDecoderUseNumber = true
	})
}

// EnableDecoderDisallowUnknownFields is used to call the DisallowUnknownFields method
// on the JSON Decoder instance. DisallowUnknownFields causes the Decoder to
// return an error when the destination is a struct and the input contains object
// keys which do not match any non-ignored, exported fields in the destination.
func EnableDecoderDisallowUnknownFields() GHttpServerOption {
	return optionFunc(func(he *GHttpEngine) {
		binding.EnableDecoderDisallowUnknownFields = true
	})
}
