import React from "react";
import t8ns from "../../translations";
import { Lang, D } from "../../etc";
import Button from "../basic/Button";
import ButtonLink from "../basic/ButtonLink";

type Props = {
  lang: Lang;
  d: D;
  loading: boolean;
};

export default ({ lang, d, loading }: Props) => {
  return (
    <div className="Welcome">
      <h1>{t8ns[lang].resName}</h1>
      <Button
        value={t8ns[lang].Welcome.play}
        onClick={() => d({ t: "GoNameChoose" })}
        disabled={loading}
      />
      <Button
        value={t8ns[lang].Welcome.learnHow}
        onClick={() => d({ t: "GoHowTo" })}
      />
      <Button
        value={t8ns[lang].Welcome.setLang}
        onClick={() => d({ t: "GoLangChoose" })}
      />
      <ButtonLink
        value={t8ns[lang].Welcome.viewCode}
        href="https://github.com/azdavis/resistance"
      />
    </div>
  );
};
