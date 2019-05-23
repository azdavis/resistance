import React from "react";
import t8ns from "../../translations";
import { Lang, State, Action } from "../../etc";

type Props = {
  lang: Lang;
  s: State;
  a: Action;
};

export default ({ lang, s, a }: Props) => (
  <div className="Fatal">
    {t8ns[lang].Fatal.title}
    {t8ns[lang].Fatal.body}
    <pre>{JSON.stringify({ s, a }, null, 2)}</pre>
  </div>
);
