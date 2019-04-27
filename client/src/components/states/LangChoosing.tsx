import React from "react";
import { Lang, D, langs } from "../../etc";
import { back } from "../../text";
import Button from "../basic/Button";
import Toggle from "../basic/Toggle";

type Props = {
  lang: Lang;
  d: D;
};

const text = {
  title: {
    en: <h1>Set language</h1>,
    ja: <h1>言語の設定</h1>,
  },
  langNames: {
    en: "English",
    ja: "日本語",
  },
};

export default ({ lang, d }: Props) => (
  <div className="LangChoosing">
    {text.title[lang]}
    {langs.map(x => (
      <Toggle
        key={x}
        value={text.langNames[x]}
        checked={lang === x}
        onChange={() => d({ t: "SetLang", lang: x })}
      />
    ))}
    <Button value={back[lang]} onClick={() => d({ t: "GoWelcome" })} />
  </div>
);
