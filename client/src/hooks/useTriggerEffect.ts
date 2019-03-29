import { useReducer, useEffect, Reducer, EffectCallback as EC } from "react";

const flip: Reducer<boolean, void> = (s, _) => !s;

export default (effect: EC, deps?: ReadonlyArray<any>): (() => void) => {
  const [s, d] = useReducer(flip, true);
  // HACK to get the effect to re-run when we want to. if deps was undefined,
  // then it gets run on every render. calling d will trigger a re-render, so no
  // need to include `s` in the array of deps.
  useEffect(effect, deps ? deps.concat([s]) : deps);
  // HACK there's no need to explicitly pass `undefined` to `d`, so we can lie
  // to the type system.
  return d as any;
};
