import React from "react";
import { CID, Client } from "../types";
import { getCaptain } from "../consts";

type Props = {
  clients: Array<Client>;
  captain: CID;
  members: number;
};

export default ({ captain, clients, members }: Props) => (
  <div className="MemberWaiter">
    <h1>New mission</h1>
    <p>
      {getCaptain(clients, captain)} is selecting {members} members for the
      mission.
    </p>
  </div>
);
