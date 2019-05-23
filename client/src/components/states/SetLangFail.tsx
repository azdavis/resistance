import React from "react";

type Props = {
  msg: string;
};

export default ({ msg }: Props) => (
  <div className="SetLangFail">
    <h1>Fail</h1>
    <p>Unable to set the language.</p>
    <pre>{msg}</pre>
  </div>
);
