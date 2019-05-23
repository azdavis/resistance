import React from "react";
import { Translation, State, Action } from "../../etc";

type Props = {
  t: Translation;
  s: State;
  a: Action;
};

export default ({ t, s, a }: Props) => (
  <div className="Invalid">
    <h1>{t.Invalid.title}</h1>
    <p>{t.Invalid.body}</p>
    <pre>{JSON.stringify({ s, a }, null, 2)}</pre>
  </div>
);
