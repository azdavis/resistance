import React from "react";
import { CID, Client, okGameSize } from "../../shared";
import { Translation, D, S } from "../../etc";
import Button from "../basic/Button";
import "../basic/Truncated.css";

type Props = {
  t: Translation;
  d: D;
  send: S;
  me: CID;
  leader: CID;
  clients: Array<Client>;
};

export default ({ t, d, send, me, leader, clients }: Props) => (
  <div className="LobbyWaiting">
    <h1>{t.lobbyWaiting(clients.length)}</h1>
    {clients.map(({ CID, Name }) => (
      <div className="Truncated" key={CID}>
        {Name}
      </div>
    ))}
    <Button
      value={t.leave}
      onClick={() => {
        d({ t: "GoLobbies" });
        send({ t: "LobbyLeave" });
      }}
    />
    <Button
      value={t.start}
      onClick={() => send({ t: "GameStart" })}
      disabled={me !== leader || !okGameSize(clients.length)}
    />
  </div>
);
