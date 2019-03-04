import React from "react";
import { Send, PartyInfo } from "../types";
import Button from "./Button";

type Props = {
  send: Send;
  parties: Array<PartyInfo>;
};

export default ({ send, parties }: Props): JSX.Element => (
  <div className="PartyChooser">
    <h1>Parties</h1>
    <Button
      value="Create new party"
      onClick={() => send({ T: "PartyCreate" })}
    />
    <h2>Existing parties ({parties.length})</h2>
    {parties.map(({ PID, Leader }) => (
      <Button
        key={PID}
        value={Leader}
        onClick={() => send({ T: "PartyChoose", PID })}
      />
    ))}
  </div>
);
