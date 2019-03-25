import React from "react";
import { D, Send } from "../types";
import { maxPts } from "../consts";
import Button from "./basic/Button";

type Props = {
  d: D;
  send: Send;
  success: boolean;
  resPts: number;
  spyPts: number;
};

export default ({ d, send, success, resPts, spyPts }: Props) => (
  <div className="MissionResultViewer">
    <h1>Mission result</h1>
    <p>The mission {success ? "succeeded" : "failed"}.</p>
    <p>The resistance has {resPts} points.</p>
    <p>The spies have {spyPts} points.</p>
    {resPts < maxPts && spyPts < maxPts ? (
      <Button value="Ok" onClick={() => d({ t: "AckMissionResult" })} />
    ) : (
      <>
        <p>The {resPts > spyPts ? "resistance has" : "spies have"} won.</p>
        <Button
          value="Leave"
          onClick={() => {
            d({ t: "GoLobbies" });
            send({ t: "GameLeave" });
          }}
        />
      </>
    )}
  </div>
);
