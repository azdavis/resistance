import React from "react";
import { Lang, D } from "../../types";
import { back } from "../../text";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  d: D;
};

const text = {
  title: {
    en: <h1>Set language</h1>,
  },
  langNames: {
    en: "English",
  },
};

const langs: Array<Lang> = ["en"];

export default ({ lang, d }: Props) => (
  <div className="LangChoosing">
    {text.title[lang]}
    {langs.map(lang => (
      <Button
        key={lang}
        value={text.langNames[lang]}
        onClick={() => d({ t: "SetLang", lang })}
      />
    ))}
    <Button value={back[lang]} onClick={() => d({ t: "GoWelcome" })} />
  </div>
);
