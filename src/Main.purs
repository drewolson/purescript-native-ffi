module Main where

import Prelude

import Effect (Effect)
import Effect.Class.Console as Console
import Effect.Uncurried (EffectFn1, EffectFn2, runEffectFn1, runEffectFn2)

foreign import _stuff :: EffectFn1 Int Int

foreign import _other :: EffectFn2 Int Int Int

foreign import myRunEffectFn2 :: forall a b c. EffectFn2 a b c -> a -> b -> Effect c

main :: Effect Unit
main = do
  result <- runEffectFn1 _stuff 1

  Console.logShow result

  result2 <- myRunEffectFn2 _other 1 2

  Console.logShow result2

  result3 <- runEffectFn2 _other 1 2

  Console.logShow result3
