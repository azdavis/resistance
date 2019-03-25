import React from "react";
import { Send, Lobby } from "../../types";
import Button from "../basic/Button";

type Props = {
  send: Send;
  lobbies: Array<Lobby>;
};

export default ({ send, lobbies }: Props) => (
  <div className="LobbyChoosing">
    <h1>Lobbies</h1>
    <Button value="Create new" onClick={() => send({ t: "LobbyCreate" })} />
    <h2>Existing lobbies ({lobbies.length})</h2>
    {lobbies.map(({ GID, Leader }) => (
      <Button
        key={GID}
        value={Leader}
        onClick={() => send({ t: "LobbyChoose", GID })}
      />
    ))}
  </div>
);
