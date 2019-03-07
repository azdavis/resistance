import React from "react";
import { Send, Lobby } from "../types";
import Button from "./Button";

type Props = {
  send: Send;
  lobbies: Array<Lobby>;
};

export default ({ send, lobbies }: Props): JSX.Element => (
  <div className="LobbyChooser">
    <h1>Lobbies</h1>
    <Button
      value="Create new lobby"
      onClick={() => send({ T: "LobbyCreate" })}
    />
    <h2>Existing lobbies ({lobbies.length})</h2>
    {lobbies.map(({ GID, Leader }) => (
      <Button
        key={GID}
        value={Leader}
        onClick={() => send({ T: "LobbyChoose", GID })}
      />
    ))}
  </div>
);
