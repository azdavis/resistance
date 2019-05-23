import React from "react";
import { Translation, D } from "../../etc";
import Button from "../basic/Button";
import ButtonLink from "../basic/ButtonLink";

type Props = {
  t: Translation;
  d: D;
  loading: boolean;
};

export default ({ t, d, loading }: Props) => {
  const { Welcome: W, resName } = t;
  return (
    <div className="Welcome">
      <h1>{resName}</h1>
      <Button
        value={W.play}
        onClick={() => d({ t: "GoNameChoose" })}
        disabled={loading}
      />
      <Button value={W.learnHow} onClick={() => d({ t: "GoHowTo" })} />
      <Button value={W.setLang} onClick={() => d({ t: "GoLangChoose" })} />
      <ButtonLink
        value={W.viewCode}
        href="https://github.com/azdavis/resistance"
      />
    </div>
  );
};
