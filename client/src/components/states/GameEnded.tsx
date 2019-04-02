import React from "react";
import { Lang, D } from "../../types";
import Button from "../basic/Button";
import Scoreboard from "../basic/Scoreboard";

type Props = {
  lang: Lang;
  d: D;
  resPts: number;
  spyPts: number;
};

const text = {
  leave: {
    en: "Leave",
  },
};

export default ({ lang, d, resPts, spyPts }: Props) => (
  <div className="GameEnded">
    <Scoreboard resPts={resPts} spyPts={spyPts} />
    <Button value={text.leave[lang]} onClick={() => d({ t: "GoLobbies" })} />
  </div>
);
