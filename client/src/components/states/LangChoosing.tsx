import React from "react";
import { D, Translation, Lang, langs } from "../../etc";
import Button from "../basic/Button";
import Toggle from "../basic/Toggle";

type Props = {
  t: Translation;
  d: D;
  setLang: React.Dispatch<React.SetStateAction<Lang>>;
};

export default ({ t, d, setLang }: Props) => (
  <div className="LangChoosing">
    <h1>{t.setLang}</h1>
    <Button value={t.back} onClick={() => d({ t: "GoWelcome" })} />
    {langs.map(([k, v]) => (
      <Toggle
        key={k}
        value={v}
        checked={t.lang === k}
        onChange={() => setLang(k)}
      />
    ))}
  </div>
);
