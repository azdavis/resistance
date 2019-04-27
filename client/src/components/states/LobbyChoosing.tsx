import React from "react";
import { Lobby } from "../../shared";
import { Lang, Send } from "../../etc";
import fullWidth from "../../fullWidth";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  send: Send;
  lobbies: Array<Lobby>;
};

const text = {
  title: {
    en: <h1>Lobbies</h1>,
    ja: <h1>ロビー</h1>,
  },
  create: {
    en: "Create new",
    ja: "新たなのを作成する",
  },
  existing: {
    en: (n: number) => <h2>Existing lobbies ({n})</h2>,
    ja: (n: number) => <h2>存在するロビー（{fullWidth(n)}）</h2>,
  },
};

export default ({ lang, send, lobbies }: Props) => (
  <div className="LobbyChoosing">
    {text.title[lang]}
    <Button
      value={text.create[lang]}
      onClick={() => send({ t: "LobbyCreate" })}
    />
    {text.existing[lang](lobbies.length)}
    {lobbies.map(({ GID, Leader }) => (
      <Button
        key={GID}
        value={Leader}
        onClick={() => send({ t: "LobbyChoose", GID })}
      />
    ))}
  </div>
);
