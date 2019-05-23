import React from "react";
import t8ns from "../../translations";
import { CID, Client, okGameSize } from "../../shared";
import { Lang, D, S } from "../../etc";
import Button from "../basic/Button";
import "../basic/Truncated.css";

type Props = {
  lang: Lang;
  d: D;
  send: S;
  me: CID;
  leader: CID;
  clients: Array<Client>;
};

export default ({ lang, d, send, me, leader, clients }: Props) => (
  <div className="LobbyWaiting">
    {t8ns[lang].LobbyWaiting.title(clients.length)}
    {clients.map(({ CID, Name }) => (
      <div className="Truncated" key={CID}>
        {Name}
      </div>
    ))}
    <Button
      value={t8ns[lang].leave}
      onClick={() => {
        d({ t: "GoLobbies" });
        send({ t: "LobbyLeave" });
      }}
    />
    <Button
      value={t8ns[lang].LobbyWaiting.start}
      onClick={() => send({ t: "GameStart" })}
      disabled={me !== leader || !okGameSize(clients.length)}
    />
  </div>
);
