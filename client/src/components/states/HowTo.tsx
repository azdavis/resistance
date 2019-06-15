import React from "react";
import { Translation, D } from "../../etc";
import Button from "../basic/Button";

type Props = {
  t: Translation;
  d: D;
};

// ok to use index as the key, since the array is immutable.
const mkP = (x: string, i: number) => <p key={i}>{x}</p>;

export default ({ t, d }: Props) => (
  <div className="HowTo">
    <h1>{t.howToPlay}</h1>
    <Button value={t.back} onClick={() => d({ t: "GoWelcome" })} />
    {t.howTo.map(mkP)}
  </div>
);
