import React from "react";
import { Send, Lobby } from "../types";
import Button from "./Button";

type Props = {
  send: Send;
  lobbies: Array<Lobby>;
};

export default ({ send, lobbies }: Props) => (
  <div className="LobbyChooser">
    <h1>Lobbies</h1>
    <Button
      value="Create new lobby"
      onClick={() => send({ t: "LobbyCreate" })}
    />
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
