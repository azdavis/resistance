import React from "react";
import SpyStatus from "./SpyStatus";

type Props = {
  captain: string;
  isSpy: boolean;
  numMembers: number;
};

export default ({ captain, isSpy, numMembers }: Props) => (
  <div className="MemberWaiter">
    <h1>New mission</h1>
    <SpyStatus isSpy={isSpy} />
    <p>
      {captain}, the captain, is selecting {numMembers} members for the mission.
    </p>
  </div>
);
