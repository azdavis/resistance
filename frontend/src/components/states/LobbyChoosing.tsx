import React from "react";
import { Lobby } from "../../shared";
import { Translation, S } from "../../etc";
import Button from "../basic/Button";

type Props = {
  t: Translation;
  send: S;
  lobbies: Array<Lobby>;
};

export default ({ t, send, lobbies }: Props) => (
  <div className="LobbyChoosing">
    <h1>{t.lobbies}</h1>
    <Button value={t.createNew} onClick={() => send({ t: "LobbyCreate" })} />
    <h2>{t.existingLobbies(lobbies.length)}</h2>
    {lobbies.map(({ GID, Leader }) => (
      <Button
        key={GID}
        value={Leader}
        onClick={() => send({ t: "LobbyChoose", GID })}
      />
    ))}
  </div>
);
