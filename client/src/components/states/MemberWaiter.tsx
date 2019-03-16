import React from "react";

type Props = {
  captain: string;
  numMembers: number;
};

export default ({ captain, numMembers }: Props) => (
  <div className="MemberWaiter">
    <h1>New mission</h1>
    <p>
      {captain}, the captain, is selecting {numMembers} members for the mission.
    </p>
  </div>
);
