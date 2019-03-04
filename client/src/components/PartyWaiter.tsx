import React from "react";
import { Send, CID, ClientInfo } from "../types";
import Button from "./Button";

type Props = {
  send: Send;
  self: CID;
  leader: CID;
  clients: Array<ClientInfo>;
};

const modifiers = (cid: CID, self: CID, leader: CID): string =>
  cid === self && cid === leader
    ? " (you, leader)"
    : cid === self
    ? " (you)"
    : cid === leader
    ? " (leader)"
    : "";

export default ({ send, self, leader, clients }: Props): JSX.Element => (
  <div className="PartyWaiter">
    <h1>Party</h1>
    <h2>Members</h2>
    {clients.map(({ CID, Name }) => (
      <div key={CID}>
        {Name}
        {modifiers(CID, self, leader)}
      </div>
    ))}
    <h2>Actions</h2>
    <Button value="Leave" onClick={() => send({ T: "PartyLeave" })} />
    {self === leader && (
      <Button value="Start" onClick={() => send({ T: "GameStart" })} />
    )}
  </div>
);
