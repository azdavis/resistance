import React from "react";
import { D, Send } from "../types";
import { maxWin } from "../consts";
import Button from "./basic/Button";

type Props = {
  d: D;
  send: Send;
  success: boolean;
  resWin: number;
  spyWin: number;
};

export default ({ d, send, success, resWin, spyWin }: Props) => (
  <div className="MissionResultViewer">
    <h1>Mission result</h1>
    <p>The mission {success ? "succeeded" : "failed"}.</p>
    <p>The resistance has {resWin} points.</p>
    <p>The spies have {spyWin} points.</p>
    {resWin < maxWin && spyWin < maxWin ? (
      <Button value="Ok" onClick={() => d({ t: "AckMissionResult" })} />
    ) : (
      <>
        <p>The {resWin > spyWin ? "resistance has" : "spies have"} won.</p>
        <Button
          value="Return"
          onClick={() => {
            d({ t: "GameLeave" });
            send({ t: "GameLeave" });
          }}
        />
      </>
    )}
  </div>
);