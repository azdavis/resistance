import React from "react";
import { Translation, State, Action } from "../../etc";

type Props = {
  t: Translation;
  s: State;
  a: Action;
};

export default ({ t, s, a }: Props) => {
  const F = t.Fatal;
  return (
    <div className="Fatal">
      <h1>{F.title}</h1>
      <p>{F.body}</p>
      <pre>{JSON.stringify({ s, a }, null, 2)}</pre>
    </div>
  );
};
