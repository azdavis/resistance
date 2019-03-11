import React from "react";
import SpyStatus from "./SpyStatus";

type Props = {
  captain: string;
  isSpy: boolean;
};

export default ({ captain, isSpy }: Props) => (
  <div className="MemberWaiter">
    <h1>New mission</h1>
    <SpyStatus isSpy={isSpy} />
    <p>{captain}, the mission captain, is selecting members for the mission.</p>
  </div>
);
