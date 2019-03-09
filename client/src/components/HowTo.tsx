import React from "react";
import { D } from "../types";
import Button from "./Button";

type Props = {
  d: D;
};

export default ({ d }: Props) => (
  <div className="HowTo">
    <h1>How to play</h1>
    <p>TODO</p>
    <Button value="Return" onClick={() => d({ t: "GoNameChoose" })} />
  </div>
);
