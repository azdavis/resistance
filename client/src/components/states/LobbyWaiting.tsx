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

export default ({ t, d, send, me, leader, clients }: Props) => {
  const { LobbyWaiting: LW, leave } = t;
  return (
    <div className="LobbyWaiting">
      <h1>{LW.title(clients.length)}</h1>
      {clients.map(({ CID, Name }) => (
        <div className="Truncated" key={CID}>
          {Name}
        </div>
      ))}
      <Button
        value={leave}
        onClick={() => {
          d({ t: "GoLobbies" });
          send({ t: "LobbyLeave" });
        }}
      />
      <Button
        value={LW.start}
        onClick={() => send({ t: "GameStart" })}
        disabled={me !== leader || !okGameSize(clients.length)}
      />
    </div>
  );
};
