import React from "react";
import { Translation, D } from "../../etc";
import Button from "../basic/Button";

type Props = {
  t: Translation;
  d: D;
};

export default ({ t, d }: Props) => {
  const HT = t.HowTo;
  return (
    <div className="HowTo">
      <h1>{HT.title}</h1>
      <p>{HT.groupSize}</p>
      <p>{HT.groupNames}</p>
      <p>{HT.decideWinner}</p>
      <p>{HT.captain}</p>
      <p>{HT.occurVote}</p>
      <p>{HT.noOccur}</p>
      <p>{HT.tooManyNoOccur}</p>
      <p>{HT.yesOccur}</p>
      <p>{HT.succeed}</p>
      <p>{HT.fail}</p>
      <Button value={t.back} onClick={() => d({ t: "GoWelcome" })} />
    </div>
  );
};
