import React from "react";
import { D, Send, CID, Client } from "../../types";
import { okGameSize } from "../../consts";
import Button from "../Button";
import "./Truncated.css";

type Props = {
  d: D;
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

export default ({ d, send, me, leader, clients }: Props) => (
  <div className="LobbyWaiter">
    <h1>Lobby</h1>
    <h2>Members ({clients.length})</h2>
    {clients.map(({ CID, Name }) => (
      <div className="Truncated" key={CID}>
        {Name}
        {modifiers(CID, me, leader)}
      </div>
    ))}
    <h2>Actions</h2>
    <Button
      value="Leave"
      onClick={() => {
        d({ t: "LobbyLeave" });
        send({ t: "LobbyLeave" });
      }}
    />
    <Button
      value="Start"
      onClick={() => send({ t: "GameStart" })}
      disabled={me !== leader || !okGameSize(clients.length)}
    />
  </div>
);
