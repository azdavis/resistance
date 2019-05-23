import React from "react";
import { D, Translation, Lang, langs, langNames } from "../../etc";
import Button from "../basic/Button";
import Toggle from "../basic/Toggle";

type Props = {
  t: Translation;
  d: D;
  setLang: React.Dispatch<React.SetStateAction<Lang>>;
};

export default ({ t, d, setLang }: Props) => (
  <div className="LangChoosing">
    <h1>{t.LangChoosing.title}</h1>
    {langs.map(x => (
      <Toggle
        key={x}
        value={langNames[x]}
        checked={t.code === x}
        onChange={() => setLang(x)}
      />
    ))}
    <Button value={t.back} onClick={() => d({ t: "GoWelcome" })} />
  </div>
);
