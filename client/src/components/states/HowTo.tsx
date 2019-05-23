import React from "react";
import t8ns from "../../translations";
import { Lang, D } from "../../etc";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  d: D;
};

export default ({ lang, d }: Props) => {
  const { HowTo: t8n, back } = t8ns[lang];
  return (
    <div className="HowTo">
      {t8n.title}
      {t8n.groupSize}
      {t8n.groupNames}
      {t8n.decideWinner}
      {t8n.captain}
      {t8n.occurVote}
      {t8n.noOccur}
      {t8n.tooManyNoOccur}
      {t8n.yesOccur}
      {t8n.succeed}
      {t8n.fail}
      <Button value={back} onClick={() => d({ t: "GoWelcome" })} />
    </div>
  );
};
