import React from "react";
import t8ns from "../../translations";
import { Lang, D } from "../../etc";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  d: D;
};

export default ({ lang, d }: Props) => (
  <div className="HowTo">
    {t8ns[lang].HowTo.title}
    {t8ns[lang].HowTo.groupSize}
    {t8ns[lang].HowTo.groupNames}
    {t8ns[lang].HowTo.decideWinner}
    {t8ns[lang].HowTo.captain}
    {t8ns[lang].HowTo.occurVote}
    {t8ns[lang].HowTo.noOccur}
    {t8ns[lang].HowTo.tooManyNoOccur}
    {t8ns[lang].HowTo.yesOccur}
    {t8ns[lang].HowTo.succeed}
    {t8ns[lang].HowTo.fail}
    <Button value={t8ns[lang].back} onClick={() => d({ t: "GoWelcome" })} />
  </div>
);
