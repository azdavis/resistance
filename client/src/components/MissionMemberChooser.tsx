import React from "react";
import { Send, CID, Client } from "../types";

type Props = {
  send: Send;
  captain: CID;
  me: CID;
  clients: Array<Client>;
};

export default ({  }: Props): JSX.Element => (
  <div className="MissionMemberChooser">
    <h1>NYI</h1>
  </div>
);
