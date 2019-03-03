import React from "react";
import { Send, CID, ClientInfo } from "../types";
import Button from "./Button";

type Props = {
  send: Send;
  leader: CID;
  clients: Array<ClientInfo>;
};

export default ({ send, leader, clients }: Props): JSX.Element => {
  return (
    <div className="PartyWaiter">
      <h1>Party</h1>
      {clients.map(({ CID, Name }) => (
        <div key={CID}>{Name}</div>
      ))}
      <Button value="Leave" onClick={() => send({ T: "PartyLeave" })} />
    </div>
  );
};
