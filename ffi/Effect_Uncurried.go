package ffi

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Effect.Uncurried")

	exports["runEffectFn2"] = func(fn_ Any) Any {
		return func(a Any) Any {
			return func(b Any) Any {
				return func() Any {
					fn := fn_.(Fn2)

					return fn(a, b)
				}
			}
		}
	}
}
