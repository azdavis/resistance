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

export default ({ lang, send, lobbies }: Props) => (
  <div className="LobbyChoosing">
    {t8ns[lang].LobbyChoosing.title}
    <Button
      value={t8ns[lang].LobbyChoosing.create}
      onClick={() => send({ t: "LobbyCreate" })}
    />
    {t8ns[lang].LobbyChoosing.existing(lobbies.length)}
    {lobbies.map(({ GID, Leader }) => (
      <Button
        key={GID}
        value={Leader}
        onClick={() => send({ t: "LobbyChoose", GID })}
      />
    ))}
  </div>
);
