import React from "react";
import { Send } from "../types";
import Voter from "./basic/Voter";

type Props = {
  send: Send;
};

const options: Array<[string, boolean]> = [["Succeed", true], ["Fail", false]];

export default ({ send }: Props) => (
  <div className="MissionVoter">
    <h1>Mission vote</h1>
    <Voter
      prompt="Should the mission succeed?"
      options={options}
      onVote={Vote => send({ t: "MissionVote", Vote })}
    />
  </div>
);
