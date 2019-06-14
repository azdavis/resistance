import React from "react";
import { Translation, D } from "../../etc";
import Button from "../basic/Button";

type Props = {
  t: Translation;
  d: D;
};

export default ({ t, d }: Props) => (
  <div className="HowTo">
    <h1>{t.howToPlay}</h1>
    <p>{t.groupSize}</p>
    <p>{t.groupNames}</p>
    <p>{t.decideWinner}</p>
    <p>{t.rounds}</p>
    <p>{t.occurVote}</p>
    <p>{t.noOccur}</p>
    <p>{t.tooManyNoOccur}</p>
    <p>{t.yesOccur}</p>
    <p>{t.succeedPt}</p>
    <p>{t.failPt}</p>
    <Button value={t.back} onClick={() => d({ t: "GoWelcome" })} />
  </div>
);
