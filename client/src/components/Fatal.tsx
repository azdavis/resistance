import React from "react";
import { State, Action } from "../types";

type Props = {
  s: State;
  a: Action;
};

export default ({ s, a }: Props): JSX.Element => (
  <div className="Fatal">
    <h1>Fatal error</h1>
    <p>
      An occurrence, that the developer of this app did not foresee occurring,
      occurred. Said occurrence is shown below.
    </p>
    <pre>{JSON.stringify({ s, a }, null, 2)}</pre>
  </div>
);
