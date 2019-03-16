import React from "react";
import { CID, Client } from "../../types";
import { getCaptain } from "../../consts";

type Props = {
  clients: Array<Client>;
  captain: CID;
  numMembers: number;
};

export default ({ captain, clients, numMembers }: Props) => (
  <div className="MemberWaiter">
    <h1>New mission</h1>
    <p>
      {getCaptain(clients, captain)} is selecting {numMembers} members for the
      mission.
    </p>
  </div>
);
