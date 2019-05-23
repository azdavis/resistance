import React from "react";
import t8ns from "../../translations";
import { Lang, D, langs } from "../../etc";
import Button from "../basic/Button";
import Toggle from "../basic/Toggle";

type Props = {
  lang: Lang;
  d: D;
};

export default ({ lang, d }: Props) => {
  const { LangChoosing: t8n, back } = t8ns[lang];
  return (
    <div className="LangChoosing">
      <h1>{t8n.title}</h1>
      {langs.map(x => (
        <Toggle
          key={x}
          value={t8ns[x].langName}
          checked={lang === x}
          onChange={() => d({ t: "SetLang", lang: x })}
        />
      ))}
      <Button value={back} onClick={() => d({ t: "GoWelcome" })} />
    </div>
  );
};
