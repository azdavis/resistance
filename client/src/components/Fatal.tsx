import React from "react";
import { State, Action } from "../types";

type Props = {
  s: State;
  a: Action;
};

export default ({ s, a }: Props) => (
  <div className="Fatal">
    <h1>Fatal error</h1>
    <p>An unforseen occurrence occurred.</p>
    <pre>{JSON.stringify({ s, a }, null, 2)}</pre>
  </div>
);
