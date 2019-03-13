import React from "react";
import { Send } from "../types";
import Button from "./Button";

type Props = {
  send: Send;
};

export default ({ send }: Props) => (
  <div className="MissionVoter">
    <h1>Mission vote</h1>
    <p>Should the mission succeed?</p>
    <Button
      value="The mission should succeed"
      onClick={() => send({ t: "MissionVote", Vote: true })}
    />
    <Button
      value="The mission should fail"
      onClick={() => send({ t: "MissionVote", Vote: false })}
    />
  </div>
);
