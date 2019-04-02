import React, { useState } from "react";
import Toggle from "./Toggle";

type Props<T> = {
  prompt: string;
  options: Array<[string, T]>;
  onVote: (x: T) => void;
};

export default function<T>({ prompt, options, onVote }: Props<T>) {
  const [vote, setVote] = useState<string | null>(null);
  const didVote = vote !== null;
  return (
    <div className="Voter">
      <p>{prompt}</p>
      {options.map(([k, v]) => (
        <Toggle
          key={k}
          value={k}
          checked={vote === k}
          onChange={() => {
            setVote(k);
            onVote(v);
          }}
          disabled={didVote}
        />
      ))}
    </div>
  );
}
