import React from "react";
import Button from "../basic/Button";
import { Lang, D } from "../../types";

type Props = {
  lang: Lang;
  d: D;
  resPts: number;
  spyPts: number;
};

const text = {
  title: {
    en: <h1>Game over</h1>,
  },
  resPts: {
    en: (n: number) => <p>Resistance points: {n}</p>,
  },
  spyPts: {
    en: (n: number) => <p>Spy points: {n}</p>,
  },
  winner: {
    en: (isRes: boolean) => <p>Winner: {isRes ? "Resistance" : "Spies"}</p>,
  },
  leave: {
    en: "Leave",
  },
};

export default ({ lang, d, resPts, spyPts }: Props) => (
  <div className="GameEnded">
    {text.title[lang]}
    {text.resPts[lang](resPts)}
    {text.spyPts[lang](spyPts)}
    {text.winner[lang](resPts > spyPts)}
    <Button value={text.leave[lang]} onClick={() => d({ t: "GoLobbies" })} />
  </div>
);
