import React from "react";
import { Translation, D } from "../../etc";
import Button from "../basic/Button";
import Scoreboard from "../basic/Scoreboard";

type Props = {
  t: Translation;
  d: D;
  resPts: number;
  spyPts: number;
};

export default ({ d, ...rest }: Props) => (
  <div className="GameEnded">
    <Scoreboard {...rest} />
    <Button value={rest.t.leave} onClick={() => d({ t: "GoLobbies" })} />
  </div>
);
