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
      <h1>{t8n.title}</h1>
      <p>{t8n.groupSize}</p>
      <p>{t8n.groupNames}</p>
      <p>{t8n.decideWinner}</p>
      <p>{t8n.captain}</p>
      <p>{t8n.occurVote}</p>
      <p>{t8n.noOccur}</p>
      <p>{t8n.tooManyNoOccur}</p>
      <p>{t8n.yesOccur}</p>
      <p>{t8n.succeed}</p>
      <p>{t8n.fail}</p>
      <Button value={back} onClick={() => d({ t: "GoWelcome" })} />
    </div>
  );
};
