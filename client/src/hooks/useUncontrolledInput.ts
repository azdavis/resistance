import { useRef, MutableRefObject } from "react";

type Input = {
  ref: MutableRefObject<HTMLInputElement | null>;
  get: () => string;
};

const useUncontrolledInput = (): Input => {
  const ref = useRef<HTMLInputElement | null>(null);
  return {
    ref,
    get: () => ref.current!.value
  };
};

export default useUncontrolledInput;
