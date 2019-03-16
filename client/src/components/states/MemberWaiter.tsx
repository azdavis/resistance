import React from "react";
import { CID, Client } from "../../types";

type Props = {
  clients: Array<Client>;
  captain: CID;
  numMembers: number;
};

export default ({ captain, clients, numMembers }: Props) => (
  <div className="MemberWaiter">
    <h1>New mission</h1>
    <p>
      The captain, {clients.find(({ CID }) => CID === captain)!.Name}, is
      selecting {numMembers} members for the mission.
    </p>
  </div>
);
