import React from "react";
import Button from "./Button";
import { Send, CID, Client } from "../types";

type Props = {
  send: Send;
  captain: CID;
  me: CID;
  clients: Array<Client>;
};

const modifiers = (cid: CID, me: CID, captain: CID): string =>
  cid === me && cid === captain
    ? " (you, captain)"
    : cid === me
    ? " (you)"
    : cid === captain
    ? " (captain)"
    : "";

export default ({ send, captain, me, clients }: Props): JSX.Element => (
  <div className="MissionMemberChooser">
    <h1>New mission</h1>
    <p>
      {me === captain
        ? "Choose the members for the mission."
        : "The mission captain is choosing the members for the mission."}
    </p>
    {clients.map(({ CID, Name }) => (
      <Button key={CID} value={`${Name}${modifiers(CID, me, captain)}`} />
    ))}
    <p>TODO</p>
  </div>
);
