import React from "react";
import { Lang, D } from "../../etc";
import { leave } from "../../text";
import Button from "../basic/Button";
import Scoreboard from "../basic/Scoreboard";

type Props = {
  lang: Lang;
  d: D;
  resPts: number;
  spyPts: number;
};

export default ({ d, ...rest }: Props) => (
  <div className="GameEnded">
    <Scoreboard {...rest} />
    <Button value={leave[rest.lang]} onClick={() => d({ t: "GoLobbies" })} />
  </div>
);
