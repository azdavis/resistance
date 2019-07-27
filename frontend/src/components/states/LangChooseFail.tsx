import React from "react";

type Props = {
  msg: string;
};

export default ({ msg }: Props) => (
  <div className="LangChooseFail">
    <h1>Error</h1>
    <p>Unable to choose the language.</p>
    <pre>{msg}</pre>
  </div>
);
