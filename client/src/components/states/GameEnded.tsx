import React from "react";
import Button from "../basic/Button";
import { D } from "../../types";

type Props = {
  d: D;
  resPts: number;
  spyPts: number;
};

export default ({ d, resPts, spyPts }: Props) => (
  <div className="GameEnded">
    <h1>Game over</h1>
    <p>Resistance points: {resPts}</p>
    <p>Spy points: {spyPts}</p>
    <p>Winner: {resPts > spyPts ? "Resistance" : "Spies"}</p>
    <Button
      value="Leave"
      onClick={() => {
        d({ t: "GoLobbies" });
      }}
    />
  </div>
);
