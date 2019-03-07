import React, { useState } from "react";
import Button from "./Button";
import { Send, CID, Client } from "../types";

type Props = {
  send: Send;
  me: CID;
  clients: Array<Client>;
};

export default ({ send, me, clients }: Props): JSX.Element => {
  const [selected, setSelected] = useState([me]);
  return (
    <div className="MissionMemberChooser">
      <h1>New mission</h1>
      <p>Choose the members for the mission.</p>
      {clients.map(({ CID, Name }) => (
        <Button key={CID} value={Name} />
      ))}
      <p>TODO</p>
    </div>
  );
};
