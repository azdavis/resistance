import React from "react";
import { Send, CID, Client } from "../types";

type Props = {
  send: Send;
  captain: CID;
  self: CID;
  leader: CID;
  clients: Array<Client>;
  isSpy: boolean;
};

export default ({  }: Props): JSX.Element => (
  <div className="MissionMemberChooser">
    <h1>NYI</h1>
  </div>
);
