import React from "react";
import { Send } from "../../types";
import Voter from "../Voter";

type Props = {
  send: Send;
};

export default ({ send }: Props) => (
  <div className="MissionVoter">
    <h1>Mission vote</h1>
    <Voter
      prompt="Should the mission succeed?"
      options={[["Succeed", true], ["Fail", false]]}
      onVote={Vote => send({ t: "MissionVote", Vote })}
    />
  </div>
);
