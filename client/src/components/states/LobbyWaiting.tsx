import React from "react";
import { D, Send, CID, Client } from "../../types";
import { okGameSize } from "../../consts";
import Button from "../basic/Button";
import "../basic/Truncated.css";

type Props = {
  d: D;
  send: Send;
  me: CID;
  leader: CID;
  clients: Array<Client>;
};

export default ({ d, send, me, leader, clients }: Props) => (
  <div className="LobbyWaiting">
    <h1>Lobby</h1>
    <h2>Members ({clients.length})</h2>
    {clients.map(({ CID, Name }) => (
      <div className="Truncated" key={CID}>
        {Name}
      </div>
    ))}
    <h2>Actions</h2>
    <Button
      value="Leave"
      onClick={() => {
        d({ t: "GoLobbies" });
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
