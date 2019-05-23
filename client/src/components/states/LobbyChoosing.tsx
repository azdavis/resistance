import React from "react";
import { Lobby } from "../../shared";
import { Translation, S } from "../../etc";
import Button from "../basic/Button";

type Props = {
  t: Translation;
  send: S;
  lobbies: Array<Lobby>;
};

export default ({ t, send, lobbies }: Props) => {
  const LC = t.LobbyChoosing;
  return (
    <div className="LobbyChoosing">
      <h1>{LC.title}</h1>
      <Button value={LC.create} onClick={() => send({ t: "LobbyCreate" })} />
      <h2>{LC.existing(lobbies.length)}</h2>
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
