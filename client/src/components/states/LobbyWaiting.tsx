import React from "react";
import { Lang, D, Send, CID, Client } from "../../shared";
import { okGameSize } from "../../shared";
import { leave } from "../../text";
import fullWidth from "../../fullWidth";
import Button from "../basic/Button";
import "../basic/Truncated.css";

type Props = {
  lang: Lang;
  d: D;
  send: Send;
  me: CID;
  leader: CID;
  clients: Array<Client>;
};

const text = {
  title: {
    en: (n: number) => <h1>Lobby ({n})</h1>,
    ja: (n: number) => <h1>ロビー（{fullWidth(n)}）</h1>,
  },
  start: {
    en: "Start",
    ja: "始める",
  },
};

export default ({ lang, d, send, me, leader, clients }: Props) => (
  <div className="LobbyWaiting">
    {text.title[lang](clients.length)}
    {clients.map(({ CID, Name }) => (
      <div className="Truncated" key={CID}>
        {Name}
      </div>
    ))}
    <Button
      value={leave[lang]}
      onClick={() => {
        d({ t: "GoLobbies" });
        send({ t: "LobbyLeave" });
      }}
    />
    <Button
      value={text.start[lang]}
      onClick={() => send({ t: "GameStart" })}
      disabled={me !== leader || !okGameSize(clients.length)}
    />
  </div>
);
