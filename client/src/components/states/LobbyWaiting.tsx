import React from "react";
import { Lang, D, Send, CID, Client } from "../../types";
import { okGameSize } from "../../shared";
import { leave } from "../../text";
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
    en: <h1>Lobby</h1>,
  },
  members: {
    en: (n: number) => <h2>Members ({n})</h2>,
  },
  actions: {
    en: <h2>Actions</h2>,
  },
  start: {
    en: "Start",
  },
};

export default ({ lang, d, send, me, leader, clients }: Props) => (
  <div className="LobbyWaiting">
    {text.title[lang]}
    {text.members[lang](clients.length)}
    {clients.map(({ CID, Name }) => (
      <div className="Truncated" key={CID}>
        {Name}
      </div>
    ))}
    {text.actions[lang]}
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
