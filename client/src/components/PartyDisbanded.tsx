import React, { Dispatch } from "react";
import { ToClient, PartyInfo } from "../types";
import Button from "./Button";

type Props = {
  parties: Array<PartyInfo>;
  d: Dispatch<ToClient>;
};

export default ({ d, parties }: Props): JSX.Element => {
  return (
    <div className="PartyDisbanded">
      <h1>Party disbanded</h1>
      <p>The party has been disbanded</p>
      <Button
        value="OK"
        onClick={() => d({ T: "PartyChoosing", Parties: parties })}
      />
    </div>
  );
};
