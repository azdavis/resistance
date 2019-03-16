import React from "react";
import { D, Send } from "../../types";
import { maxWin } from "../../consts";
import Button from "../Button";

type Props = {
  send: Send;
  d: D;
  success: boolean;
  resWin: number;
  spyWin: number;
};

export default ({ send, d, success, resWin, spyWin }: Props) => (
  <div className="MissionResultViewer">
    <h1>Mission result</h1>
    <p>The mission {success ? "succeeded" : "failed"}.</p>
    <p>The resistance has {resWin} points.</p>
    <p>The spies have {spyWin} points.</p>
    {resWin < maxWin && spyWin < maxWin ? (
      <Button value="Ok" onClick={() => d({ t: "AckMissionResult" })} />
    ) : (
      <Button
        value="Return"
        onClick={() => {
          d({ t: "GameLeave" });
          send({ t: "GameLeave" });
        }}
      />
    )}
  </div>
);
