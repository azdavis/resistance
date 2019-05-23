import React from "react";
import t8ns from "../../translations";
import { Lobby } from "../../shared";
import { Lang, S } from "../../etc";
import Button from "../basic/Button";

type Props = {
  lang: Lang;
  send: S;
  lobbies: Array<Lobby>;
};

export default ({ lang, send, lobbies }: Props) => {
  const t8n = t8ns[lang].LobbyChoosing;
  return (
    <div className="LobbyChoosing">
      <h1>{t8n.title}</h1>
      <Button value={t8n.create} onClick={() => send({ t: "LobbyCreate" })} />
      <h2>{t8n.existing(lobbies.length)}</h2>
      {lobbies.map(({ GID, Leader }) => (
        <Button
          key={GID}
          value={Leader}
          onClick={() => send({ t: "LobbyChoose", GID })}
        />
      ))}
    </div>
  );
};
