import React from "react";
import { Lang, State, Action } from "../../types";

type Props = {
  lang: Lang;
  s: State;
  a: Action;
};

const text = {
  title: {
    en: <h1>Fatal error</h1>,
    ja: <h1>致命的謝り</h1>,
  },
  body: {
    en: <p>An error occurred from which the application cannot recover.</p>,
    ja: <p>アプリが復活できぬ謝りが起きた。</p>,
  },
};

export default ({ lang, s, a }: Props) => (
  <div className="Fatal">
    {text.title[lang]}
    {text.body[lang]}
    <pre>{JSON.stringify({ s, a }, null, 2)}</pre>
  </div>
);
