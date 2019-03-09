import React from "react";
import { Send, CID, Client } from "../types";
import Button from "./Button";
import FullWidth from "./FullWidth";

type Props = {
  send: Send;
  me: CID;
  leader: CID;
  clients: Array<Client>;
};

const modifiers = (cid: CID, me: CID, leader: CID): string =>
  cid === me && cid === leader
    ? " (you, leader)"
    : cid === me
    ? " (you)"
    : cid === leader
    ? " (leader)"
    : "";

// keep in sync with lobby.go
const minN = 5;

export default ({ send, me, leader, clients }: Props) => (
  <div className="LobbyWaiter">
    <h1>Lobby</h1>
    <h2>Members</h2>
    {clients.map(({ CID, Name }) => (
      <FullWidth key={CID}>
        {Name}
        {modifiers(CID, me, leader)}
      </FullWidth>
    ))}
    <h2>Actions</h2>
    <Button value="Leave" onClick={() => send({ t: "LobbyLeave" })} />
    {me === leader && (
      <Button
        value="Start"
        onClick={() => send({ t: "GameStart" })}
        disabled={clients.length < minN}
      />
    )}
  </div>
);
