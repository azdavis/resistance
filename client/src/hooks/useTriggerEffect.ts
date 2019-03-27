import { useReducer, useEffect, Reducer, EffectCallback } from "react";

const flip: Reducer<boolean, void> = (s, _) => !s;

export default (effect: EffectCallback, deps?: ReadonlyArray<any>) => {
  const [s, d] = useReducer(flip, true);
  // HACK to get the effect to re-run when we want to.
  useEffect(effect, deps ? deps.concat([s]) : deps);
  // HACK there's no need to explicitly pass `undefined` to `d`, so we can lie
  // to the type system.
  return (d as any) as (() => void);
};
