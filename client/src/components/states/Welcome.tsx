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
  const { Welcome: t8n, resName } = t8ns[lang];
  return (
    <div className="Welcome">
      <h1>{resName}</h1>
      <Button
        value={t8n.play}
        onClick={() => d({ t: "GoNameChoose" })}
        disabled={loading}
      />
      <Button value={t8n.learnHow} onClick={() => d({ t: "GoHowTo" })} />
      <Button value={t8n.setLang} onClick={() => d({ t: "GoLangChoose" })} />
      <ButtonLink
        value={t8n.viewCode}
        href="https://github.com/azdavis/resistance"
      />
    </div>
  );
};
