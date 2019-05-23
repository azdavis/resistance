import React from "react";
import t8ns from "../../translations";
import { Lang, State, Action } from "../../etc";

type Props = {
  lang: Lang;
  s: State;
  a: Action;
};

export default ({ lang, s, a }: Props) => {
  const t8n = t8ns[lang].Fatal;
  return (
    <div className="Fatal">
      <h1>{t8n.title}</h1>
      <p>{t8n.body}</p>
      <pre>{JSON.stringify({ s, a }, null, 2)}</pre>
    </div>
  );
};
