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
  return (
    <div className="LangChoosing">
      {t8ns[lang].LangChoosing.title}
      {langs.map(x => (
        <Toggle
          key={x}
          value={t8ns[x].langName}
          checked={lang === x}
          onChange={() => d({ t: "SetLang", lang: x })}
        />
      ))}
      <Button value={t8ns[lang].back} onClick={() => d({ t: "GoWelcome" })} />
    </div>
  );
};
