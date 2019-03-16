import React, { useState } from "react";
import Button from "./Button";

type Props<T> = {
  prompt: string;
  options: Array<[string, T]>;
  onVote: (x: T) => void;
};

export default function<T>({ prompt, options, onVote }: Props<T>) {
  const [vote, setVote] = useState<string | null>(null);
  return (
    <div className="Voter">
      <p>{vote === null ? prompt : `You voted: ${vote}`}</p>
      {options.map(([k, v]) => (
        <Button
          key={k}
          value={k}
          onClick={() => {
            setVote(k);
            onVote(v);
          }}
          disabled={vote !== null}
        />
      ))}
    </div>
  );
}
