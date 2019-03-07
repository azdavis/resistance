import React from "react";

type Props = {
  captain: string;
};

export default ({ captain }: Props): JSX.Element => (
  <div className="MissionMemberWaiter">
    <h1>New mission</h1>
    <p>{captain}, the mission captain, is selecting members for the mission.</p>
  </div>
);
