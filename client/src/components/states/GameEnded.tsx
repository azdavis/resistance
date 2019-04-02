import React from "react";
import Scoreboard from "../basic/Scoreboard";
import Button from "../basic/Button";
import { Lang, D } from "../../types";

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
