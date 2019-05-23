import React from "react";
import t8ns from "../../translations";
import { Lang, D } from "../../etc";
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
    <Button
      value={t8ns[rest.lang].leave}
      onClick={() => d({ t: "GoLobbies" })}
    />
  </div>
);
