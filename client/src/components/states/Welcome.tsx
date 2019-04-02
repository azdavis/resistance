import React from "react";
import { Lang, D } from "../../types";
import Button from "../basic/Button";
import ButtonLink from "../basic/ButtonLink";

type Props = {
  lang: Lang;
  d: D;
  loading: boolean;
};

const text = {
  title: {
    en: <h1>Resistance</h1>,
  },
  play: {
    en: "Play",
  },
  learnHow: {
    en: "Learn how to play",
  },
  viewCode: {
    en: "View source code",
  },
};

export default ({ lang, d, loading }: Props) => {
  return (
    <div className="Welcome">
      {text.title[lang]}
      <Button
        value={text.play[lang]}
        onClick={() => d({ t: "GoNameChoose" })}
        disabled={loading}
      />
      <Button value={text.learnHow[lang]} onClick={() => d({ t: "GoHowTo" })} />
      <ButtonLink
        value={text.viewCode[lang]}
        href="https://github.com/azdavis/resistance"
      />
    </div>
  );
};
