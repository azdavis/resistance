import React from "react";
import { Lang, D } from "../../types";
import { resName } from "../../text";
import Button from "../basic/Button";
import ButtonLink from "../basic/ButtonLink";

type Props = {
  lang: Lang;
  d: D;
  loading: boolean;
};

const text = {
  play: {
    en: "Play",
    ja: "遊ぶ",
  },
  learnHow: {
    en: "Learn how to play",
    ja: "遊び方を知る",
  },
  setLang: {
    en: "Set language",
    ja: "言語を設定する",
  },
  viewCode: {
    en: "View source code",
    ja: "コードを見る",
  },
};

export default ({ lang, d, loading }: Props) => {
  return (
    <div className="Welcome">
      <h1>{resName[lang]}</h1>
      <Button
        value={text.play[lang]}
        onClick={() => d({ t: "GoNameChoose" })}
        disabled={loading}
      />
      <Button value={text.learnHow[lang]} onClick={() => d({ t: "GoHowTo" })} />
      <Button
        value={text.setLang[lang]}
        onClick={() => d({ t: "GoLangChoose" })}
      />
      <ButtonLink
        value={text.viewCode[lang]}
        href="https://github.com/azdavis/resistance"
      />
    </div>
  );
};
