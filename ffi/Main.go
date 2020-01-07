package ffi

import (
	. "github.com/purescript-native/go-runtime"
)

func init() {
	exports := Foreign("Main")

	exports["_stuff"] = func(i_ Any) Any {
		i := i_.(int)

		return i + 1
	}

	exports["_other"] = func(i_, j_ Any) Any {
		i := i_.(int)
		j := j_.(int)

		return i + j
	}

	exports["myRunEffectFn2"] = func(fn_ Any) Any {
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
