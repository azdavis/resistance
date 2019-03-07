import React from "react";
import { State, Action } from "../types";

type Props = {
  s: State;
  a: Action;
};

const showJSON = (x: any): string => JSON.stringify(x, null, 2);

export default ({ s, a }: Props): JSX.Element => (
  <div className="Invalid">
    <h1>Fatal error</h1>
    <p>This is probably the developer's fault.</p>
    <p>Previous state:</p>
    <pre>{showJSON(s)}</pre>
    <p>Action:</p>
    <pre>{showJSON(a)}</pre>
  </div>
);
