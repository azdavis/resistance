import React from "react";
import { D } from "../types";
import Button from "./Button";

type Props = {
  d: D;
};

export default ({ d }: Props) => (
  <div className="Disbanded">
    <h1>Disbanded</h1>
    <p>The game or lobby you were in was disbanded.</p>
    <Button
      value="Return to lobbies"
      onClick={() => d({ t: "AckDisbanded" })}
    />
  </div>
);
