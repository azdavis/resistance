import { useEffect, useReducer, EffectCallback } from "react";

const reducer = (s: boolean, _: void): boolean => !s;

export default (
  effect: EffectCallback,
  deps?: ReadonlyArray<any>,
): (() => void) => {
  const [s, d] = useReducer(reducer, true);
  useEffect(effect, [s].concat(deps || []));
  return (d as any) as (() => void);
};
