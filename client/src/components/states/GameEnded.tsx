import React from "react";
import { Lang, D } from "../../types";
import { leave } from "../../text";
import Button from "../basic/Button";
import Scoreboard from "../basic/Scoreboard";

type Props = {
  lang: Lang;
  d: D;
  resPts: number;
  spyPts: number;
};

export default ({ lang, d, resPts, spyPts }: Props) => (
  <div className="GameEnded">
    <Scoreboard lang={lang} resPts={resPts} spyPts={spyPts} />
    <Button value={leave[lang]} onClick={() => d({ t: "GoLobbies" })} />
  </div>
);
