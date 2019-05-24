import React from "react";
import { Translation, D } from "../../etc";
import Button from "../basic/Button";
import ButtonLink from "../basic/ButtonLink";

type Props = {
  t: Translation;
  d: D;
  loading: boolean;
};

export default ({ t, d, loading }: Props) => (
  <div className="Welcome">
    <h1>{t.resName}</h1>
    <Button
      value={t.play}
      onClick={() => d({ t: "GoNameChoose" })}
      disabled={loading}
    />
    <Button value={t.learnHow} onClick={() => d({ t: "GoHowTo" })} />
    <Button value={t.setLang} onClick={() => d({ t: "GoLangChoose" })} />
    <ButtonLink
      value={t.viewCode}
      href="https://github.com/azdavis/resistance"
    />
  </div>
);
